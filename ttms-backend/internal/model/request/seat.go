/**
 * @Author: lenovo
 * @Description:
 * @File:  seat
 * @Version: 1.0.0
 * @Date: 2023/06/04 6:27
 */

package request

import "mognolia/internal/model/automigrate"

type ShowSeatsParam struct {
	CinemaID uint `json:"cinemaID" binding:"required"`
	PlanID   uint `json:"planID" binding:"required"`
}

type UpdateSeatStatus struct {
	CinemaID uint                   `json:"cinemaID" binding:"required"`
	SeatID   uint                   `json:"seatID" binding:"required"`
	Status   automigrate.SeatStatus `json:"status" binding:"required,oneof=normal broken"`
}
