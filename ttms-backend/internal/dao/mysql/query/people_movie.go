/**
 * @Author: lenovo
 * @Description:
 * @File:  people_movie
 * @Version: 1.0.0
 * @Date: 2023/06/01 22:49
 */

package query

import (
	"errors"
	"mognolia/internal/dao"
	"mognolia/internal/model/automigrate"
)

type qm struct{}

func NewQM() *qm {
	return &qm{}
}

var ErrUserHasFavor = errors.New("用户已经关注过了这部电影")

func (qm) CreateFavorToMovie(userID, movieID uint) (uint, error) {
	u := &automigrate.PeopleMovie{
		UserID:  userID,
		MovieID: movieID,
	}
	if result := dao.Group.DB.Model(&automigrate.PeopleMovie{}).Create(u); result.RowsAffected == 0 {
		return 0, result.Error
	}
	return u.ID, nil
}

func (qm) DelUserFavor(userID, movieID uint) error {
	result := dao.Group.DB.Model(&automigrate.PeopleMovie{}).Where("user_id=? and movie_id =?", userID, movieID).Delete(&automigrate.PeopleMovie{})
	return result.Error
}

func (qm) CheckUserAndMovieIsRepeat(userID, movieID uint) error {
	if result := dao.Group.DB.Model(&automigrate.PeopleMovie{}).Where("user_id = ? and movie_id =?", userID, movieID).Find(&automigrate.PeopleMovie{}); result.RowsAffected != 0 {
		return ErrUserHasFavor
	}
	return nil
}
