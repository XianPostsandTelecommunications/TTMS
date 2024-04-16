/**
 * @Author: lenovo
 * @Description:
 * @File:  cinema
 * @Version: 1.0.0
 * @Date: 2023/06/04 1:45
 */

package query

import (
	"gorm.io/gorm"
	"mognolia/internal/dao"
	"mognolia/internal/model/automigrate"
	"mognolia/internal/model/request"
)

type cinema struct{}

func NewCinema() *cinema {
	return &cinema{}
}

func (cinema) CreateCinema(req request.CreateCinema) error {
	if result := dao.Group.DB.Model(&automigrate.Cinema{}).Create(&automigrate.Cinema{
		Name: req.Name,
		Rows: req.Rows,
		Cols: req.Cols,
	}); result.Error != nil {
		return result.Error
	}
	return nil
}
func (cinema) GetAllCinemaByPage(page int) ([]automigrate.Cinema, error) {
	var cinemas []automigrate.Cinema
	if result := dao.Group.DB.Model(&automigrate.Cinema{}).Scopes(Paginate(int64(page), 0)).Find(&cinemas); result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return cinemas, nil
}

func (cinema) UpdateAvatarOrName(req request.UpdateCinemaInfo) error {
	if result := dao.Group.DB.Model(&automigrate.Cinema{}).Where("id =?", req.CinemaID).Updates(&automigrate.Cinema{
		Avatar: req.Avatar,
		Name:   req.Name,
	}); result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}

func (cinema) GetCinamemaDetails(cinemaID int) (*automigrate.Cinema, error) {
	var cinema automigrate.Cinema
	if result := dao.Group.DB.Model(&automigrate.Cinema{}).Where("id =?", cinemaID).Find(&cinema); result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &cinema, nil
}
