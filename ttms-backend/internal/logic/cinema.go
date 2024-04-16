/**
 * @Author: lenovo
 * @Description:
 * @File:  cinema
 * @Version: 1.0.0
 * @Date: 2023/06/04 1:33
 */

package logic

import (
	"errors"
	"gorm.io/gorm"
	"mognolia/internal/dao/mysql/query"
	tx2 "mognolia/internal/dao/mysql/tx"
	"mognolia/internal/model/reply"
	"mognolia/internal/model/request"
	"mognolia/internal/myerr"
	"mognolia/internal/pkg/app/errcode"
)

type cinema struct{}

func (cinema) CreateCinema(req request.CreateCinema) errcode.Err {
	tx := tx2.NewCinemaTX()
	if err := tx.CreateCinema(req); err != nil {
		return errcode.ErrServer.WithDetails(err.Error())
	}
	return nil
}

func (cinema) DeleteCinema(cinemaID int) errcode.Err {
	tx := tx2.NewCinemaTX()
	if err := tx.DeleteCinema(cinemaID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return myerr.NoRecords
		}
		return errcode.ErrServer
	}
	return nil
}

func (cinema) GetAllCinemaByPage(page int) ([]reply.CinemaInfo, errcode.Err) {
	var result []reply.CinemaInfo
	q := query.NewCinema()
	cinemas, err := q.GetAllCinemaByPage(page)
	if err != nil {
		return nil, errcode.ErrServer
	}
	for i := range cinemas {
		result = append(result, reply.CinemaInfo{
			ID:     cinemas[i].ID,
			Name:   cinemas[i].Name,
			Avatar: cinemas[i].Avatar,
			Rows:   cinemas[i].Rows,
			Cols:   cinemas[i].Cols,
		})
	}
	return result, nil
}

func (c *cinema) UpdateCinemaAvatarOrName(req request.UpdateCinemaInfo) errcode.Err {
	q := query.NewCinema()
	if err := q.UpdateAvatarOrName(req); err != nil {
		return errcode.ErrServer
	}
	return nil
}

func (c cinema) GetCinemaDetails(cinemaID int) (*reply.CinemaInfo, errcode.Err) {
	q := query.NewCinema()
	cinemaInfo, err := q.GetCinamemaDetails(cinemaID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerr.NoRecords
		}
		return nil, errcode.ErrServer
	}
	return &reply.CinemaInfo{
		Name:   cinemaInfo.Name,
		Avatar: cinemaInfo.Avatar,
		Rows:   cinemaInfo.Rows,
		Cols:   cinemaInfo.Cols,
	}, nil
}
