/**
 * @Author: lenovo
 * @Description:
 * @File:  seat
 * @Version: 1.0.0
 * @Date: 2023/06/04 6:29
 */

package logic

import (
	"go.uber.org/zap"
	"mognolia/internal/dao/mysql/query"
	tx2 "mognolia/internal/dao/mysql/tx"
	"mognolia/internal/model/reply"
	"mognolia/internal/model/request"
	"mognolia/internal/pkg/app/errcode"
)

type seat struct{}

func (seat) ShowSeats(planID, cinemaID uint) (*reply.ShowCinemaSeats, errcode.Err) {
	var result reply.ShowCinemaSeats
	tx := tx2.NewSeatWithTX()
	ShowCinemaSeatsRows, err := tx.GetAllSeatsInfoByCinemaID(cinemaID)
	if err != nil {
		zap.S().Info("Error getting ", zap.Any("err", err))
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}
	q1 := query.NewCinema()
	cinemaInfo, err := q1.GetCinamemaDetails(int(cinemaID))
	if err != nil {
		zap.S().Info("Error getting ", zap.Any("err", err))
		return nil, errcode.ErrServer.WithDetails(err.Error())
	}
	showSeats := make([][]*reply.SeatsInfo, cinemaInfo.Rows)
	for i := 0; i < len(showSeats); i++ {
		showSeats[i] = make([]*reply.SeatsInfo, cinemaInfo.Cols)
	}
	for i := range ShowCinemaSeatsRows {
		seatInfo := &reply.SeatsInfo{
			SeatID: ShowCinemaSeatsRows[i].ID,
			Row:    ShowCinemaSeatsRows[i].Row,
			Col:    ShowCinemaSeatsRows[i].Col,
			Status: ShowCinemaSeatsRows[i].Status,
		}
		showSeats[ShowCinemaSeatsRows[i].Row-1][ShowCinemaSeatsRows[i].Col-1] = seatInfo
	}
	result.CinemaSeats = showSeats
	return &result, nil
}

func (seat) UpdateSeatStatus(req request.UpdateSeatStatus) errcode.Err {
	q := query.NewSeat()
	if err := q.UpdateSeatStatus(req); err != nil {
		return errcode.ErrServer.WithDetails(err.Error())
	}
	return nil
}
