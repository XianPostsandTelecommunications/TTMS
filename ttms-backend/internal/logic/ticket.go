/**
 * @Author: lenovo
 * @Description:
 * @File:  ticket
 * @Version: 1.0.0
 * @Date: 2023/06/06 8:20
 */

package logic

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"mognolia/internal/dao"
	"mognolia/internal/dao/mysql/query"
	tx2 "mognolia/internal/dao/mysql/tx"
	"mognolia/internal/global"
	"mognolia/internal/model/automigrate"
	"mognolia/internal/model/reply"
	"mognolia/internal/model/request"
	"mognolia/internal/myerr"
	"mognolia/internal/pkg/app/errcode"
	"strconv"
	"sync"
	"time"
)

type ticket struct {
	Mutex *sync.RWMutex
}

func NewTicket() *ticket {
	return &ticket{Mutex: &sync.RWMutex{}}
}

func (t *ticket) GetTicketsByPlan(planID int64) (*reply.ShowTicket, errcode.Err) {
	q := query.NewPlan()
	cinema, err := q.GetCinemaByPlanID(planID)
	if err != nil {
		zap.S().Infof("getCinemaByPlanID failed: %v", err)
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}
	tickets, err := q.GetTicketsByPlan(planID)
	if err != nil {
		zap.S().Infof("getTicketsByPlan failed: %v", err)
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}
	rlyTicketInfos := make([][]reply.TicketInfo, cinema.Rows)
	for i := range rlyTicketInfos {
		rlyTicketInfos[i] = make([]reply.TicketInfo, cinema.Cols)
	}
	for i := range tickets {
		t := reply.TicketInfo{
			SeatID:       tickets[i].SeatID,
			Row:          tickets[i].Row,
			Col:          tickets[i].Col,
			Price:        tickets[i].Price,
			CinemaID:     tickets[i].CinemaID,
			Status:       tickets[i].Status,
			TicketStatus: tickets[i].TicketStatus,
		}
		rlyTicketInfos[t.Row-1][t.Col-1] = t
	}
	return &reply.ShowTicket{Tickets: rlyTicketInfos}, nil
}

func (t *ticket) LockTickets(userID uint, req request.LockTickets) (*reply.CreateOrderRly, errcode.Err) {

	var seats []string
	for _, seat := range req.SeatIDs {
		seats = append(seats, strconv.Itoa(int(seat)))
	}
	if err := dao.Group.Redis.LockTicket(context.Background(), seats, req.PlanID, userID); err != nil {
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}
	timeout := global.Settings.Rule.LockTicketTime
	go func() {
		select {
		case <-SeatMap:
			key := fmt.Sprintf("%d:%d", req.PlanID, req.SeatIDs)
			if err := dao.Group.Redis.Expire2Zero(context.Background(), key); err != nil {
				zap.S().Infof("err:%v", err)
			}
			return
		case <-time.After(timeout):
			if result := dao.Group.DB.Model(&automigrate.Ticket{}).Where("plan_id = ? ", req.PlanID).
				Where("seat_id in ?", req.SeatIDs).Updates(automigrate.Ticket{
				UserID:       userID,
				TicketStatus: automigrate.ForSaleStatus,
			}); result.RowsAffected == 0 {
				return
			}
			if err := dao.Group.Redis.DelSeats(context.Background(), seats, req.PlanID); err != nil {
				return
			}
		}
	}()
	//缓存中已经有了,数据库同步
	tx := tx2.TicketTX{}
	rly, err := tx.LockTickets(req, userID)
	if err != nil {
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}
	if err := dao.Group.Redis.SetOrder(context.Background(), rly.OrderID, userID); err != nil {
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}

	go func() {
		select {
		case <-time.After(timeout):
			//要把之前修改的订单状态给改回去
			if result := dao.Group.DB.Model(&automigrate.Order{}).Where("order_id = ?", rly.OrderID).Updates(automigrate.Order{Status: automigrate.OrderNotPaid}); result.RowsAffected == 0 {
				zap.S().Infof("updated order failed")
				return
			}
		case <-TicketMap:
			return
		}
	}()
	return &reply.CreateOrderRly{
		IsLock:  true,
		OrderID: rly.OrderID,
	}, nil

}

func (t *ticket) PayTicket(userID uint, req request.PayTicketReq) errcode.Err {
	if ok := dao.Group.Redis.CheckLockedAndIsSelf(context.Background(), req.PlanID, userID); !ok {
		return myerr.NotLockTicket
	}
	for _, v := range req.SeatIDs {
		tx := tx2.NewTicketTX()
		if err := tx.PayTicket(request.PayTickerRowReq{
			UserID:   userID,
			OrderID:  req.OrderID,
			SeatIDs:  uint(v),
			PlanID:   req.PlanID,
			CinemaID: req.CinemaID,
		}); err != nil {
			return errcode.ErrServer.WithDetails(err.Error())
		}
	}
	SeatMap <- true
	TicketMap <- true
	return nil
}

func (t *ticket) LockOneTicket(req request.LockOneTicketReq, userID uint) errcode.Err {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	var seats []string
	seats = append(seats, strconv.Itoa(int(req.SeatIDs)))

	//先去判断一下是否锁了票
	if ok := dao.Group.Redis.CheckIsLock(context.Background(), strconv.Itoa(int(req.SeatIDs)), req.PlanID, userID); ok {
		return myerr.ErrLockTickets
	}
	if err := dao.Group.Redis.LockOneTicket(context.Background(), strconv.Itoa(int(req.SeatIDs)), req.PlanID, userID); err != nil {
		return errcode.ErrServer.WithDetails(err.Error())
	}
	timeout := global.Settings.Rule.LockTicketTime
	go func() {
		select {
		case <-SeatMap:
			key := fmt.Sprintf("%d:%d", req.PlanID, req.SeatIDs)
			if err := dao.Group.Redis.Expire2Zero(context.Background(), key); err != nil {
				zap.S().Infof("err:%v", err)
			}
			return
		case <-time.After(timeout):
			if result := dao.Group.DB.Model(&automigrate.Ticket{}).Where("plan_id = ? ", req.PlanID).
				Where("seat_id in ?", req.SeatIDs).Updates(automigrate.Ticket{
				UserID:       userID,
				TicketStatus: automigrate.ForSaleStatus,
			}); result.RowsAffected == 0 {
				return
			}
			if err := dao.Group.Redis.DelSeats(context.Background(), seats, req.PlanID); err != nil {
				return
			}
		}
	}()
	return nil
}

func (ticket) BackTicket(req request.BackTicketParam, userID uint) errcode.Err {
	tx := tx2.NewTicketTX()
	if err := tx.CheckTicket(req, userID); err != nil {
		return errcode.ErrServer.WithDetails(err.Error())
	}
	//更改票的状态
	if err := tx.UpdateTicketStatus(req); err != nil {
		return errcode.ErrServer
	}
	return nil
}

func (ticket) SearchTicket(uuid string) (bool, errcode.Err) {
	tx := tx2.NewTicketTX()
	if err := tx.SearchTicket(uuid); err != nil {
		return false, errcode.ErrServer.WithDetails(err.Error())
	}
	return true, nil
}

func (ticket) GetAllOrders(userID uint) (*reply.Orders, errcode.Err) {
	var orders []automigrate.Order
	if result := dao.Group.DB.Model(&automigrate.Order{}).Where("user_id = ?", userID).Find(&orders); result.RowsAffected == 0 {
		return nil, myerr.NoRecords
	}
	var res reply.Orders
	for _, order := range orders {
		t := &reply.OrderInfoRly{
			PlanID:     order.PlanID,
			OrderID:    order.OrderID.String(),
			SeatIDs:    order.SeatsID,
			Seats:      order.Seats,
			UpdateTime: order.UpdatedAt,
			Price:      order.Price,
			Status:     order.Status,
		}
		res.OrderInfos = append(res.OrderInfos, t)
	}
	return &res, nil
}
