/**
 * @Author: lenovo
 * @Description:
 * @File:  cinema
 * @Version: 1.0.0
 * @Date: 2023/06/04 1:23
 */

package request

type CreateCinema struct {
	Name   string `json:"name" binding:"required"`
	Rows   int    `json:"rows" binding:"required"`
	Cols   int    `json:"cols" binding:"required"`
	Avatar string `json:"avatar"`
}

type DeleteCinema struct {
	CinemaID int `json:"cinema_id" binding:"required"`
}

type GetCinemaByPage struct {
	Page int `json:"page" binding:"required"`
}

type UpdateCinemaInfo struct {
	CinemaID uint   `json:"cinemaID" binding:"required"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
}

type GetCinemaDetails struct {
	CinemaID int `json:"cinemaID" binding:"required"`
}
