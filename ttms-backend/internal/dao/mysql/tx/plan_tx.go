/**
 * @Author: lenovo
 * @Description:
 * @File:  plan_tx
 * @Version: 1.0.0
 * @Date: 2023/06/04 5:09
 */

package tx

import (
	"errors"
	"gorm.io/gorm"
	"mognolia/internal/dao"
	"mognolia/internal/model/automigrate"
	"mognolia/internal/model/reply"
	"mognolia/internal/model/request"
	"time"
)

var ErrPlanConflict = errors.New("演出计划出现冲突")

type planTX struct{}

func NewPlanWithTX() *planTX {
	return &planTX{}
}

func (planTX) CreatePlanWithTX(req request.CreatePlan) (*reply.CreatePlanRly, error) {

	res := &reply.CreatePlanRly{}
	tx := dao.Group.DB.Begin()
	//先判断该电影是否存在
	var movieInfo automigrate.Movie
	if result := tx.Model(&automigrate.Movie{}).Where("id =?", req.MovieID).Find(&movieInfo); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, gorm.ErrRecordNotFound
	}
	if result := tx.Model(&automigrate.Cinema{}).Where("id=?", req.CinemaID).Find(&automigrate.Cinema{}); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, gorm.ErrRecordNotFound
	}
	endAt := time.Unix(req.StartTime, 0).Add(time.Duration(movieInfo.Duration) * time.Minute).Truncate(time.Second)
	var ids []uint

	//这个sql语句还是有问题
	if result := tx.Raw(`
SELECT
	pl.id 
FROM
	plans pl
WHERE
	pl.cinema_id = ?
	AND (
		pl.start_at <= ? AND pl.end_at >= ? 
		OR pl.start_at >= ? 
		AND pl.end_at <= ? 
		OR pl.start_at BETWEEN ? 
		AND ? 
		OR pl.end_at BETWEEN ? 
	AND ? 
	)
	AND pl.deleted_at IS NULL
`, req.CinemaID, time.Unix(req.StartTime, 0), endAt, time.Unix(req.StartTime, 0), endAt, time.Unix(req.StartTime, 0), endAt, time.Unix(req.StartTime, 0), endAt).Scan(&ids); result.RowsAffected != 0 {
		tx.Rollback()
		return nil, ErrPlanConflict
	}

	//创建plan表
	plan := &automigrate.Plan{
		MovieID:  req.MovieID,
		CinemaID: req.CinemaID,
		StartAt:  time.Unix(req.StartTime, 0).Truncate(time.Second),
		EndAt:    endAt,
		Price:    req.Price,
	}
	if result := tx.Model(&automigrate.Plan{}).Create(&plan); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, result.Error
	}

	//TODO:获取seat表,以及ticket
	var seats []automigrate.Seat
	if result := tx.Model(&automigrate.Seat{}).Where("cinema_id = ?", req.CinemaID).Find(&seats); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, gorm.ErrRecordNotFound
	}

	//创建票

	ticketInfos := make([]automigrate.Ticket, 0, len(seats))
	for i := range seats {
		ticketInfos = append(ticketInfos, automigrate.Ticket{
			Price:        req.Price,
			PlanID:       plan.ID,
			SeatID:       seats[i].ID,
			TicketStatus: automigrate.ForSaleStatus,
		})
	}
	if result := tx.Model(&automigrate.Ticket{}).Create(&ticketInfos); result.RowsAffected == 0 {
		return nil, result.Error
	}
	res.PlanID = plan.ID
	res.EndAt = endAt
	res.StartAt = plan.StartAt

	tx.Commit()
	return res, nil
}

func (planTX) DeletePlan(planID uint) error {
	tx := dao.Group.DB.Begin()
	//查询

	//在演出计划中,是不能被删除的
	//	在演出时间内的planID有没有锁上
	if result := tx.Raw(`
SELECT
	1
FROM
	plan p,
	tickets t 
WHERE
	p.id = ? 
	AND t.plan_id = p.id 
	AND (
	s.STATUS = '"saled"' 
	OR s.STATUS = '"lock"')
	AND p.end_at > NOW() 
`, planID); result.RowsAffected != 0 {
		tx.Rollback()
		return ErrPlanConflict
	}
	if result := tx.Model(&automigrate.Plan{}).Where("id = ?", planID).Delete(&automigrate.Plan{}); result.RowsAffected == 0 {
		tx.Rollback()
		return result.Error
	}
	if result := tx.Model(&automigrate.Ticket{}).Where("plan_id = ?", planID).Delete(&automigrate.Ticket{}); result.RowsAffected == 0 {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}

func (planTX) GetPlansByMovieIDAndPeriod(req request.GetPlansByMovieIDAndPeriod) ([]automigrate.Plan, error) {
	var plans []automigrate.Plan
	tx := dao.Group.DB.Begin()
	if result := tx.Model(&automigrate.Plan{}).Where("movie_id = ?", req.MovieID).
		Where("end_at between ? and ?", time.Unix(req.StartAt, 0), time.Unix(req.EndAt, 0)).
		Where("start_at > now()").Preload("Movie").Preload("Cinema").Find(&plans); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, gorm.ErrRecordNotFound
	}
	return plans, nil
}
