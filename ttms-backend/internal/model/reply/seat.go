/**
 * @Author: lenovo
 * @Description:
 * @File:  seat
 * @Version: 1.0.0
 * @Date: 2023/06/04 6:32
 */

package reply

import "mognolia/internal/model/automigrate"

type SeatsInfo struct {
	SeatID uint
	Row    int
	Col    int
	Status automigrate.SeatStatus
}

type ShowCinemaSeats struct {
	CinemaSeats [][]*SeatsInfo
}
