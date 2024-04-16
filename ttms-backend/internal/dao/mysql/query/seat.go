/**
 * @Author: lenovo
 * @Description:
 * @File:  seat
 * @Version: 1.0.0
 * @Date: 2023/06/04 16:16
 */

package query

import (
	"gorm.io/gorm"
	"mognolia/internal/dao"
	"mognolia/internal/model/automigrate"
	"mognolia/internal/model/request"
)

type seat struct{}

func NewSeat() *seat {
	return &seat{}
}

func (seat) GetAllSeatsInfoByCinemaID(cinemaID uint) ([]automigrate.Seat, error) {
	var seats []automigrate.Seat
	result := dao.Group.DB.Model(&automigrate.Seat{}).Where("cinema_id=?", cinemaID).Find(&seats)
	return seats, result.Error
}

func (seat) UpdateSeatStatus(req request.UpdateSeatStatus) error {
	if result := dao.Group.DB.Model(&automigrate.Seat{}).Where(automigrate.Seat{
		Model:    gorm.Model{ID: req.SeatID},
		CinemaID: req.CinemaID,
	}).Update("status", req.Status); result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
