/**
 * @Author: lenovo
 * @Description:
 * @File:  people_movie
 * @Version: 1.0.0
 * @Date: 2023/06/01 22:32
 */

package logic

import (
	"errors"
	"github.com/gin-gonic/gin"
	"mognolia/internal/dao/mysql/query"
	"mognolia/internal/middleware"
	"mognolia/internal/model/reply"
	"mognolia/internal/myerr"
	"mognolia/internal/pkg/app/errcode"
)

type pm struct{}

func (pm) CreateFavorToMovie(ctx *gin.Context, movieID uint) (*reply.CreateFavorRly, errcode.Err) {
	q := query.NewQM()
	result := &reply.CreateFavorRly{}
	context, err := middleware.GetContext(ctx)
	if err != nil {
		return nil, errcode.ErrUnauthorizedAuthNotExist
	}
	id, err1 := q.CreateFavorToMovie(context.ID, movieID)
	if err1 != nil {
		return nil, errcode.ErrServer.WithDetails(err1.Error())
	}
	result.CreateID = id
	return result, nil
}

func (pm) DelUserFavor(ctx *gin.Context, movieID uint) errcode.Err {
	q := query.NewQM()
	context, err := middleware.GetContext(ctx)
	if err != nil {
		return errcode.ErrUnauthorizedAuthNotExist
	}
	if err := q.DelUserFavor(context.ID, movieID); err != nil {
		return errcode.ErrServer.WithDetails(err.Error())
	}
	return nil
}

func (pm) CheckIsRepeat(userID, movieID uint) errcode.Err {
	q := query.NewQM()
	if err := q.CheckUserAndMovieIsRepeat(userID, movieID); err != nil {
		if errors.Is(err, query.ErrUserHasFavor) {
			return myerr.ErrHasFavoriteThisMovie
		}
		return errcode.ErrServer
	}
	return nil
}
