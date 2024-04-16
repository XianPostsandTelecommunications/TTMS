/**
 * @Author: lenovo
 * @Description:
 * @File:  plan
 * @Version: 1.0.0
 * @Date: 2023/06/05 12:08
 */

package query

import (
	"gorm.io/gorm"
	"mognolia/internal/dao"
	"mognolia/internal/model/automigrate"
	"mognolia/internal/model/reply"
)

type plan struct{}

func NewPlan() *plan {
	return &plan{}
}

func (plan) GetPlanList(page int64) (plansInfo []automigrate.Plan, err error) {
	if result := dao.Group.DB.Model(&automigrate.Plan{}).Scopes(Paginate(page, 0)).Preload("Cinema").Preload("Movie").Find(&plansInfo); result.RowsAffected == 0 {
		return plansInfo, gorm.ErrRecordNotFound
	}
	return plansInfo, nil
}

func (plan) GetCinemaByPlanID(planID int64) (*automigrate.Cinema, error) {
	var plan automigrate.Plan
	if result := dao.Group.DB.Model(&automigrate.Plan{}).Preload("Cinema").Where(automigrate.Plan{
		Model: gorm.Model{ID: uint(planID)},
	}).Preload("Cinema").Find(&plan); result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &plan.Cinema, nil
}

func (plan) GetTicketsByPlan(planID int64) ([]reply.TicketInfo, error) {
	var res []reply.TicketInfo
	if result := dao.Group.DB.Table("tickets").Select("seats.row, seats.col, seats.cinema_id, tickets.ticket_status, tickets.price, seats.status,tickets.seat_id").
		Joins("JOIN seats ON tickets.seat_id = seats.id").
		Where("tickets.plan_id = ? AND tickets.deleted_at IS NULL AND seats.deleted_at IS NULL", planID).
		Scan(&res); result.Error != nil {
		return nil, result.Error
		// 处理查询错误
	}

	return res, nil
}
