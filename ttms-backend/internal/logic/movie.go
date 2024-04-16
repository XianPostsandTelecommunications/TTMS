/**
 * @Author: lenovo
 * @Description:
 * @File:  movie
 * @Version: 1.0.0
 * @Date: 2023/05/31 8:51
 */

package logic

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"mognolia/internal/dao"
	"mognolia/internal/dao/mysql/query"
	tx2 "mognolia/internal/dao/mysql/tx"
	"mognolia/internal/global"
	"mognolia/internal/middleware"
	"mognolia/internal/model/automigrate"
	"mognolia/internal/model/reply"
	"mognolia/internal/model/request"
	"mognolia/internal/myerr"
	"mognolia/internal/pkg/app/errcode"
	"mognolia/internal/pkg/utils"
)

type movie struct{}

const (
	OrderByPeriod    = "period"
	OrderByReadCount = "readCount"
	OrderByScore     = "score"
)

const (
	GetByExpectedKey = "expected"
)

func (m *movie) CreateMovie(ctx *gin.Context, param request.CreateMovieReq) (*reply.CreateMovieRly, errcode.Err) {
	tx := tx2.NewMovie()
	content, err := middleware.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	rly, err1 := tx.CreateMovieWithTX(content.ID, param)
	if err1 != nil {
		if errors.Is(err, tx2.ErrMovieAlreadyExists) {
			return nil, myerr.ErrMovieAlreadyExists
		}
		return nil, errcode.ErrServer.WithDetails(err1.Error())
	}
	return &reply.CreateMovieRly{MovieID: rly.MovieID}, nil
}

func (m *movie) DeleteMovie(ctx *gin.Context, movieID uint) errcode.Err {
	//TODO:先去判断一下该电影时候在演出计划中,演出计划中的电影不能删除

	content, err := middleware.GetContext(ctx)
	if err != nil {
		return err
	}
	uid := content.ID
	tx := tx2.NewMovie()
	if err := tx.DeleteMovieByID(uid, movieID); err != nil {
		return errcode.ErrServer.WithDetails(err.Error())
	}
	return nil
}

func (m *movie) GetMovieByID(movieID uint) (*reply.GetMovieDetails, errcode.Err) {
	tx := tx2.NewMovie()
	movieDetail, err := tx.GetMovieDetails(movieID)
	if err != nil {
		return nil, myerr.NoRecords
	}
	_, err = dao.Group.Redis.AddReadCountToMovie(context.Background(), movieID)
	if err != nil {
		return nil, errcode.ErrServer
	}
	return movieDetail, nil
}

// 我对不起这个接口 呜呜呜~~~~~~~~~~
func (m *movie) GetMovieByTagAreaPeriod(param request.GetMovieTagsAreaPeriod) (*reply.GetMovieRly, errcode.Err) {
	var result reply.GetMovieRly
	switch param.OrderType {
	case OrderByPeriod:
		q := query.NewMovie()
		movies, err := q.GetAreaTagsPeriodMovieOrderByPeriod(param, global.Settings.Rule.DefaultPagePerNum, (param.Page-1)*global.Settings.Rule.DefaultPagePerNum)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, myerr.NoRecords
			}
			return nil, errcode.ErrServer
		}
		for _, movie := range movies {
			result.MovieInfo = append(result.MovieInfo, reply.MovieInfo{
				MovieID:   movie.ID,
				Name:      movie.Name,
				Avatar:    movie.Avatar,
				Actors:    movie.Actors,
				Director:  movie.Director,
				Score:     movie.Score,
				Duration:  movie.Duration,
				ShowTime:  movie.ShowTime,
				BoxOffice: movie.BoxOffice,
			})
		}
		result.Total = int64(len(movies))
		return &result, nil
	case OrderByReadCount:
		q := query.NewMovie()
		movies, err := q.GetAreaTagsPeriodMovieOrderByReadCount(param, global.Settings.Rule.DefaultPagePerNum, (param.Page-1)*global.Settings.Rule.DefaultPagePerNum)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, myerr.NoRecords
			}
			return nil, errcode.ErrServer
		}
		for _, movie := range movies {
			result.MovieInfo = append(result.MovieInfo, reply.MovieInfo{
				MovieID:   movie.ID,
				Name:      movie.Name,
				Avatar:    movie.Avatar,
				Actors:    movie.Actors,
				Director:  movie.Director,
				Score:     movie.Score,
				Duration:  movie.Duration,
				ShowTime:  movie.ShowTime,
				BoxOffice: movie.BoxOffice,
			})
		}
		result.Total = int64(len(movies))
		return &result, nil
	case OrderByScore:
		q := query.NewMovie()
		movies, err := q.GetAreaTagsPeriodMovieOrderByScore(param, global.Settings.Rule.DefaultPagePerNum, (param.Page-1)*global.Settings.Rule.DefaultPagePerNum)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, myerr.NoRecords
			}
			return nil, errcode.ErrServer
		}
		for _, movie := range movies {
			result.MovieInfo = append(result.MovieInfo, reply.MovieInfo{
				MovieID:   movie.ID,
				Name:      movie.Name,
				Avatar:    movie.Avatar,
				Actors:    movie.Actors,
				Director:  movie.Director,
				Score:     movie.Score,
				Duration:  movie.Duration,
				ShowTime:  movie.ShowTime,
				BoxOffice: movie.BoxOffice,
			})
		}
		result.Total = int64(len(movies))
		return &result, nil
	}
	return &result, nil
}

func (m *movie) GetMoviesByKey(key string, page int64) (*reply.GetMovieRly, errcode.Err) {
	q := query.NewMovie()
	movies, err := q.GetKey(key, page)
	var result reply.GetMovieRly
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerr.NoRecords
		}
		return nil, errcode.ErrServer
	}
	result.Total = int64(len(movies))
	for _, movie := range movies {
		result.MovieInfo = append(result.MovieInfo, reply.MovieInfo{
			MovieID:   movie.ID,
			Name:      movie.Name,
			Avatar:    movie.Avatar,
			Actors:    movie.Actors,
			Director:  movie.Director,
			Score:     movie.Score,
			Duration:  movie.Duration,
			ShowTime:  movie.ShowTime,
			BoxOffice: movie.BoxOffice,
		})
	}
	return &result, nil
}

func (m *movie) GetMovieByExpected(page int64) (*reply.GetMovieRly, errcode.Err) {
	var result reply.GetMovieRly
	var movies []automigrate.Movie
	//  expected+page数量 取出相应的数量
	if err := dao.Group.Redis.Get(context.Background(), utils.IDToSting(uint(page)), &movies); err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, myerr.NoRecords
		}
		return nil, errcode.ErrServer
	}
	result.Total = int64(len(movies))
	for _, movie := range movies {
		result.MovieInfo = append(result.MovieInfo, reply.MovieInfo{
			MovieID:   movie.ID,
			Name:      movie.Name,
			Avatar:    movie.Avatar,
			Actors:    movie.Actors,
			Director:  movie.Director,
			Score:     movie.Score,
			Duration:  movie.Duration,
			ShowTime:  movie.ShowTime,
			BoxOffice: movie.BoxOffice,
		})
	}
	return &result, nil
}

func (movie) GetMoviesByReadCount() (*reply.GetMovieRly, errcode.Err) {
	var result reply.GetMovieRly
	q := query.NewMovie()
	movies, err := q.GetMoviesOrderbyReadCount(0)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerr.NoRecords
		}
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}
	result.Total = int64(len(movies))
	for _, movie := range movies {
		result.MovieInfo = append(result.MovieInfo, reply.MovieInfo{
			MovieID:   movie.ID,
			Name:      movie.Name,
			Avatar:    movie.Avatar,
			Actors:    movie.Actors,
			Director:  movie.Director,
			Score:     movie.Score,
			Duration:  movie.Duration,
			ShowTime:  movie.ShowTime,
			BoxOffice: movie.BoxOffice,
		})
	}
	return &result, nil
}

func (movie) UpdateMovieInfo(param request.UpdateMovieInfo) errcode.Err {
	q := query.NewMovie()
	if err := q.UpdateMovieInfo(param); err != nil {
		return errcode.ErrServer
	}
	return nil
}

func (movie) GetAllMovie() (*reply.GetMovieRly, errcode.Err) {
	result := &reply.GetMovieRly{}
	q := query.NewMovie()
	movieInfos, err := q.GetAllMoviesInfo()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerr.NoRecords
		}
		return nil, errcode.ErrServer
	}
	result.Total = int64(len(movieInfos))
	for _, movie := range movieInfos {
		result.MovieInfo = append(result.MovieInfo, reply.MovieInfo{
			MovieID:   movie.ID,
			Name:      movie.Name,
			Avatar:    movie.Avatar,
			Actors:    movie.Actors,
			Director:  movie.Director,
			Score:     movie.Score,
			Duration:  movie.Duration,
			ShowTime:  movie.ShowTime,
			BoxOffice: movie.BoxOffice,
		})
	}
	return result, nil
}
