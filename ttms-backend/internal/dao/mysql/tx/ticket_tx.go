/**
 * @Author: lenovo
 * @Description:
 * @File:  ticket_tx
 * @Version: 1.0.0
 * @Date: 2023/06/06 23:14
 */

package tx

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"mognolia/internal/dao"
	"mognolia/internal/model/automigrate"
	"mognolia/internal/model/reply"
	"mognolia/internal/model/request"
	"strconv"
)

type TicketTX struct{}

func NewTicketTX() *TicketTX {
	return &TicketTX{}
}

var (
	ErrNotOwnTicket = errors.New("不是自己的票")
)

func (t *TicketTX) LockTickets(req request.LockTickets, userID uint) (*reply.RlyOrderID, error) {
	//1.先去更新一下作为的状态
	tx := dao.Group.DB.Begin()
	if result := tx.Model(&automigrate.Ticket{}).Where("plan_id = ? ", req.PlanID).
		Where("seat_id in ?", req.SeatIDs).Updates(automigrate.Ticket{
		UserID:       userID,
		TicketStatus: automigrate.LockStatus,
	}); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, result.Error
	}

	//2.获得seats相关
	var seats string
	var seatIDs string
	for _, seatID := range req.SeatIDs {
		var seatInfo automigrate.Seat
		if result := tx.Model(&automigrate.Seat{}).Where("id = ?", seatID).Preload("Cinema").Find(&seatInfo); result.RowsAffected == 0 {
			tx.Rollback()
			return nil, result.Error
		}
		seats += strconv.Itoa(seatInfo.Row) + "---" + strconv.Itoa(seatInfo.Col) + " "
		seatIDs += strconv.Itoa(int(seatID)) + " "
	}

	//4.获取planID

	var planInfo automigrate.Plan
	if result := tx.Model(&automigrate.Plan{}).Where("id = ?", req.PlanID).Preload("Cinema").Find(&planInfo); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, result.Error
	}
	var movie automigrate.Movie
	if result := tx.Model(&automigrate.Movie{}).Where("id = ?", planInfo.MovieID).Find(&movie); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, result.Error
	}
	var userInfo automigrate.User
	if result := tx.Model(&automigrate.User{}).Where("id = ?", userID).Find(&userInfo); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, result.Error
	}
	order := &automigrate.Order{
		OrderID:    uuid.New(),
		PlanID:     planInfo.ID,
		MovieID:    movie.ID,
		UserID:     userID,
		UserName:   userInfo.UserName,
		CinemaName: planInfo.Cinema.Name,
		Status:     automigrate.OrderNotPaid,
		Seats:      seats,
		Price:      float32(planInfo.Price) * (float32(len(seatIDs) - 1)),
		SeatsID:    seatIDs,
	}
	if result := tx.Model(&automigrate.Order{}).Create(order); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, result.Error
	}
	tx.Commit()
	id := order.OrderID
	return &reply.RlyOrderID{OrderID: id}, nil
}

func (t *TicketTX) PayTicket(req request.PayTickerRowReq) error {
	//订单状态
	tx := dao.Group.DB.Begin()
	if result := tx.Model(&automigrate.Order{}).Where("order_id = ?", req.OrderID).Updates(automigrate.Order{Status: automigrate.OrderHasPaid}); result.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}
	if result := tx.Model(&automigrate.Ticket{}).Where("plan_id = ? and seat_id =?", req.PlanID, req.SeatIDs).
		Updates(automigrate.Ticket{UserID: req.UserID, TicketStatus: automigrate.SaledStatus}); result.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}
	tx.Commit()
	return nil
}

func (t *TicketTX) CheckTicket(req request.BackTicketParam, userID uint) error {
	var tickets []automigrate.Ticket
	tx := dao.Group.DB.Begin()
	if result := tx.Model(&automigrate.Ticket{}).Where("plan_id = ? and seat_id  in  ?", req.PlanID, req.SeatIDs).Find(&tickets); result.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}
	for _, ticket := range tickets {
		if ticket.UserID != userID || ticket.TicketStatus != automigrate.SaledStatus {
			return ErrNotOwnTicket
		}
	}
	return nil

}

func (t *TicketTX) UpdateTicketStatus(req request.BackTicketParam) error {
	tx := dao.Group.DB.Begin()
	if result := tx.Model(&automigrate.Ticket{}).Where("plan_id =? and seat_id in ?", req.PlanID, req.SeatIDs).
		Updates(map[string]interface{}{"ticket_status": automigrate.ForSaleStatus, "user_id": 0}); result.RowsAffected == 0 {
		tx.Rollback()
		return result.Error
	}
	if result := tx.Model(&automigrate.Order{}).Where("order_id = ?", req.OrderID).Delete(&automigrate.Order{}); result.RowsAffected == 0 {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}

func (t *TicketTX) SearchTicket(uuid string) error {
	tx := dao.Group.DB.Begin()
	if result := tx.Model(&automigrate.Order{}).Where("order_id = ?", uuid).Find(&automigrate.Order{}); result.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}
	tx.Commit()
	return nil
}
