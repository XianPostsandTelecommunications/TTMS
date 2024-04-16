/**
 * @Author: lenovo
 * @Description:
 * @File:  cinema_tx
 * @Version: 1.0.0
 * @Date: 2023/06/04 2:28
 */

package tx

import (
	"errors"
	"gorm.io/gorm"
	"mognolia/internal/dao"
	"mognolia/internal/global"
	"mognolia/internal/model/automigrate"
	"mognolia/internal/model/request"
)

type CinemaWithTX struct{}

func NewCinemaTX() *CinemaWithTX {
	return &CinemaWithTX{}
}

var ErrCinemaAlraedyExist = errors.New("该影厅已经存在了")

func (CinemaWithTX) CreateCinema(req request.CreateCinema) error {
	tx := dao.Group.DB.Begin()
	if result := tx.Model(&automigrate.Cinema{}).Where("name = ?", req.Name).Find(&automigrate.Cinema{}); result.RowsAffected != 0 {
		return ErrCinemaAlraedyExist
	}
	if req.Avatar == "" {
		req.Avatar = global.Settings.Rule.DefaultAccountAvatar
	}
	cinema := &automigrate.Cinema{
		Name:   req.Name,
		Rows:   req.Rows,
		Cols:   req.Cols,
		Avatar: req.Avatar,
	}
	if result := tx.Model(&automigrate.Cinema{}).Create(&cinema); result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	//创建seat表
	seats := make([]automigrate.Seat, 0, req.Cols*req.Rows)
	for i := 1; i <= req.Rows; i++ {
		for j := 1; j <= req.Cols; j++ {
			seats = append(seats, automigrate.Seat{
				CinemaID: cinema.ID,
				Row:      i,
				Col:      j,
				Status:   automigrate.SeatNormal,
			})
		}
	}
	if result := tx.Model(&automigrate.Seat{}).Create(&seats); result.RowsAffected == 0 {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}

func (CinemaWithTX) DeleteCinema(cinemaID int) error {
	var cinemaInfo automigrate.Cinema
	if result := dao.Group.DB.Model(&automigrate.Cinema{}).Where("id = ?", cinemaID).Find(&cinemaInfo); result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	//TODO: 这里要去查询一下,在plan表中的东西是不能删的
	tx := dao.Group.DB.Begin()

	if result := tx.Model(&automigrate.Cinema{}).Delete(&automigrate.Cinema{}, cinemaID); result.RowsAffected == 0 {
		tx.Rollback()
		return result.Error
	}
	if result := tx.Model(&automigrate.Seat{}).Where("cinema_id = ?", cinemaID).Delete(&automigrate.Seat{}); result.RowsAffected == 0 {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}
