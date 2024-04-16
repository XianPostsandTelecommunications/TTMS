/**
 * @Author: lenovo
 * @Description:
 * @File:  movie
 * @Version: 1.0.0
 * @Date: 2023/05/31 8:16
 */

package request

import (
	"fmt"
	"mognolia/internal/pkg/app/errcode"
	"time"
)

type CreateMovieReq struct {
	Name     string   `json:"name" binding:"required,gte=1,lte=20"`        //电影名称
	Duration int16    `json:"duration" binding:"required,gte=1"`           //时长(分钟)
	ShowTime int64    `json:"show_time" binding:"required"`                //上映时间戳
	Tags     []string `json:"tags" binding:"required"`                     //电影标签
	Actors   []string `json:"actors" binding:"required,gte=1,dive,max=10"` //演员
	Director string   `json:"director" binding:"required,gte=1"`           //导演
	Area     string   `json:"area" binding:"required,gte=1,lte=20"`        //地区
	Avatar   string   `json:"avatar"`
	Content  string   `json:"content" binding:"required,gte=1,max=1000"`
}

type DeleteMovieReq struct {
	MovieID uint `json:"movie_id" binding:"required"`
}

type GetMovieByID struct {
	MovieID uint `json:"movie_id" binding:"required"`
}

type GetMovieTagsAreaPeriod struct {
	Tag       string `json:"tag" `    //标签
	Area      string `json:"area" `   //地区
	Period    string `json:"period" ` // 时间
	OrderType string `json:"order_type" binding:"omitempty,oneof=period score readCount"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	Page      int64  `json:"page"`
}

func (g *GetMovieTagsAreaPeriod) Check() errcode.Err {
	var msg string
	switch {
	case g.StartTime > g.EndTime:
		msg = "起始时间不能大于结束时间"
	default:
		if g.Period != "" {
			g.Period = fmt.Sprintf("%s%", g.Period)
		} else {
			g.Period = "%"
		}

		if g.Area == "" {
			g.Area = "%"
		}
		if g.Tag == "" {
			g.Tag = "%"
		}
		if g.Page <= 0 {
			g.Page = 1
		}
		if g.EndTime == 0 {
			g.EndTime = time.Now().Unix()
		}
		return nil
	}
	return errcode.ErrServer.WithDetails(msg)
}

type GetByKey struct {
	Key  string `json:"key"`
	Page int64  `json:"page"`
}

func (g *GetByKey) Check() {
	if g.Key == "" {
		g.Key = "%"
	}
}

type GetByExpected struct {
	Page int64 `json:"page" binding:"required"`
}

type GetByVisitCount struct {
	Page int64 `json:"page" binding:"required"`
}

type UpdateMovieInfo struct {
	MovieID  int64    `json:"movieID" binding:"required"`
	Name     string   `json:"name" `
	Avatar   string   `json:"avatar"`
	Actors   []string `json:"actors"`
	Director string   `json:"director"`
	Area     string   `json:"area"`
	Content  string   `json:"content"`
}
