/**
 * @Author: lenovo
 * @Description:
 * @File:  plan
 * @Version: 1.0.0
 * @Date: 2023/06/04 4:55
 */

package logic

import (
	"errors"
	"gorm.io/gorm"
	"mognolia/internal/dao/mysql/query"
	tx2 "mognolia/internal/dao/mysql/tx"
	"mognolia/internal/model/reply"
	"mognolia/internal/model/request"
	"mognolia/internal/myerr"
	"mognolia/internal/pkg/app/errcode"
)

type plan struct{}

func (plan) CreatePlan(req request.CreatePlan) (*reply.CreatePlanRly, errcode.Err) {
	tx := tx2.NewPlanWithTX()
	rly, err := tx.CreatePlanWithTX(req)
	if err != nil {
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}
	return rly, nil
}

func (plan) DelPlan(planID uint) errcode.Err {
	tx := tx2.NewPlanWithTX()
	if err := tx.DeletePlan(planID); err != nil {
		return errcode.ErrServer.WithDetails(err.Error())
	}
	ids := make([]uint, 0)
	ids = append(ids, planID)
	return nil
}

func (plan) GetPlans(page int) (*reply.PlanList, errcode.Err) {
	result := &reply.PlanList{}
	q := query.NewPlan()
	plans, err := q.GetPlanList(int64(page))
	result.Total = int64(len(plans))
	for _, plan := range plans {
		result.Data = append(result.Data, reply.PlanInfo{
			SeatID:      plan.CinemaID,
			ID:          plan.ID,
			MovieName:   plan.Movie.Name,
			MovieAvatar: plan.Movie.Avatar,
			CinemaID:    plan.CinemaID,
			CinemaName:  plan.Cinema.Name,
			Duration:    plan.Movie.Duration,
			StartAt:     plan.StartAt,
			Price:       plan.Price,
		})
	}
	if err != nil {
		return nil, myerr.NoRecords
	}
	return result, nil
}

func (plan) GetPlansByMovieIDAndPeriod(req request.GetPlansByMovieIDAndPeriod) (*reply.GetPlanByMovieIDAndPeriod, errcode.Err) {
	res := &reply.GetPlanByMovieIDAndPeriod{}
	tx := tx2.NewPlanWithTX()
	result, err := tx.GetPlansByMovieIDAndPeriod(req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerr.NoRecords
		}
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}
	res.Total = int64(len(result))
	for _, plan := range result {
		res.PlansInfo = append(res.PlansInfo, &reply.Plans{
			MovieName:   plan.Movie.Name,
			MovieAvatar: plan.Movie.Avatar,
			CinemaName:  plan.Cinema.Name,
			StartAt:     plan.StartAt,
			EndAt:       plan.EndAt,
		})
	}
	return res, nil
}
