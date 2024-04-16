/**
 * @Author: lenovo
 * @Description:
 * @File:  movie_tx
 * @Version: 1.0.0
 * @Date: 2023/05/31 9:26
 */

package tx

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"mognolia/internal/dao"
	"mognolia/internal/model/automigrate"
	"mognolia/internal/model/reply"
	"mognolia/internal/model/request"
	"time"
)

type MovieTX struct{}

func NewMovie() *MovieTX {
	return &MovieTX{}
}

var ErrMovieAlreadyExists = errors.New("电影已经存在了")

func (MovieTX) CreateMovieWithTX(managerID uint, param request.CreateMovieReq) (*reply.CreateMovieRly, error) {
	tx := dao.Group.DB.Begin()
	if result := tx.Model(&automigrate.Movie{}).Where("name =?", param.Name).Find(&automigrate.Movie{}); result.RowsAffected != 0 {
		return nil, ErrMovieAlreadyExists
	}
	m := &automigrate.Movie{
		Name:     param.Name,
		Area:     param.Area,
		Avatar:   param.Avatar,
		Actors:   param.Actors,
		Content:  param.Content,
		Duration: int64(param.Duration),
		ShowTime: time.Unix(param.ShowTime, 0).Format("2006-01-02"),
		Director: param.Director,
	}
	if result := tx.Model(&automigrate.Movie{}).Create(m); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, result.Error
	}
	if result := tx.Model(&automigrate.Tag{}).Create(&automigrate.Tag{
		MovieID: m.ID,
		Tags:    param.Tags,
	}); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, result.Error
	}
	var u automigrate.User
	if result := dao.Group.DB.Model(&automigrate.User{}).Where("id = ?", managerID).Find(&u); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, gorm.ErrRecordNotFound
	}
	tx.Model(&automigrate.ManagerMovie{}).Create(&automigrate.ManagerMovie{
		MovieID: m.ID,
		UserID:  managerID,
		Content: fmt.Sprintf("管理员:%s创建了电影%s", u.UserName, param.Name),
	})
	tx.Commit()
	return &reply.CreateMovieRly{MovieID: m.ID}, nil
}

func (m *MovieTX) DeleteMovieByID(mID, movieID uint) error {
	var movie automigrate.Movie
	tx := dao.Group.DB.Begin()
	if result := dao.Group.DB.Model(&automigrate.Movie{}).Where("id =?", movieID).Find(&movie); result.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	if result := tx.Model(&automigrate.Tag{}).Where("movie_id = ?", movieID).Delete(&automigrate.Tag{}); result.RowsAffected == 0 {
		return result.Error
	}
	if result := tx.Model(&automigrate.Movie{}).Delete(&automigrate.Movie{}, movieID); result.RowsAffected == 0 {
		tx.Rollback()
		return result.Error
	}

	var u automigrate.User
	if result := dao.Group.DB.Model(&automigrate.User{}).Where("id = ?", mID).Find(&u); result.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}
	if result := tx.Model(&automigrate.ManagerMovie{}).Create(&automigrate.ManagerMovie{
		MovieID: movieID,
		UserID:  mID,
		Content: fmt.Sprintf("管理员:%s删除了电影%s", u.UserName, movie.Name),
	}); result.RowsAffected == 0 {
		return result.Error
	}
	tx.Commit()
	return nil
}

func (m *MovieTX) GetMovieDetails(movieID uint) (*reply.GetMovieDetails, error) {
	tx := dao.Group.DB.Begin()
	var movie automigrate.Movie
	if result := tx.Model(&automigrate.Movie{}).Where("id =?", movieID).Find(&movie); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, gorm.ErrRecordNotFound
	}
	//从tag中获取tag字段
	t := automigrate.Tag{}
	if result := dao.Group.DB.Model(&automigrate.Tag{}).Where("movie_id =?", movieID).Find(&t); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, gorm.ErrRecordNotFound
	}
	//todo: 查询是否是该用户关注过或者评论过

	return &reply.GetMovieDetails{
		MovieInfoRow: reply.MovieInfo{
			MovieID:   movie.ID,
			Name:      movie.Name,
			Content:   movie.Content,
			Actors:    movie.Actors,
			Avatar:    movie.Avatar,
			Director:  movie.Director,
			Score:     movie.Score,
			Duration:  movie.Duration,
			ShowTime:  movie.ShowTime,
			BoxOffice: movie.BoxOffice,
		},
		Tag: t.Tags,
	}, nil
}
