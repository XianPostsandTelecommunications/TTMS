/**
 * @Author: lenovo
 * @Description:
 * @File:  seat_tx
 * @Version: 1.0.0
 * @Date: 2023/06/04 16:21
 */

package tx

import (
	"gorm.io/gorm"
	"mognolia/internal/dao"
	"mognolia/internal/model/automigrate"
)

type seatWithTX struct{}

func NewSeatWithTX() *seatWithTX {
	return &seatWithTX{}
}

func (seatWithTX) GetAllSeatsInfoByCinemaID(cinemaID uint) ([]automigrate.Seat, error) {
	tx := dao.Group.DB.Begin()
	var movieInfo automigrate.Cinema
	if result := tx.Model(&automigrate.Cinema{}).Find(&movieInfo, cinemaID); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, gorm.ErrRecordNotFound
	}
	var SeatsInfoRows []automigrate.Seat
	if result := tx.Model(&automigrate.Seat{}).Where("cinema_id = ?", cinemaID).Find(&SeatsInfoRows); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, gorm.ErrRecordNotFound
	}
	tx.Commit()
	return SeatsInfoRows, nil
}
