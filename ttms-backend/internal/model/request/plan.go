/**
 * @Author: lenovo
 * @Description:
 * @File:  plan
 * @Version: 1.0.0
 * @Date: 2023/06/04 4:37
 */

package request

import (
	"errors"
	"mognolia/internal/pkg/app/errcode"
)

var ErrTime = errors.New("结束时间不能大于开始时间")

type CreatePlan struct {
	MovieID   uint    `json:"movieID" binding:"required"`
	CinemaID  uint    `json:"cinemaID" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
	StartTime int64   `json:"startTime" binding:"required"`
}

type DelPlan struct {
	PlanID uint `json:"planID" binding:"required"`
}

type GetPlanByPage struct {
	Page int `json:"page" binding:"required"`
}

type GetPlansByMovieIDAndPeriod struct {
	MovieID uint  `json:"MovieID" binding:"required"`
	StartAt int64 `json:"startAt" binding:"required"`
	EndAt   int64 `json:"endAt" binding:"required"`
}

func (p *GetPlansByMovieIDAndPeriod) Check() errcode.Err {
	if p.EndAt < p.StartAt {
		return errcode.ErrParamsNotValid.WithDetails()
	}
	return nil
}
