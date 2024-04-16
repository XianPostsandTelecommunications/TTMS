/**
 * @Author: lenovo
 * @Description:
 * @File:  ticket
 * @Version: 1.0.0
 * @Date: 2023/06/06 20:11
 */

package request

type LockTickets struct {
	PlanID   uint   `json:"planID"`
	CinemaID uint   `json:"cinemaID"`
	SeatIDs  []uint `json:"seatIDs"`
}

type PayTicketReq struct {
	OrderID  string  `json:"orderID"`
	SeatIDs  []int64 `json:"seatIDs"`
	PlanID   uint    `json:"planID"`
	CinemaID uint    `json:"cinemaID"`
}

type PayTickerRowReq struct {
	UserID   uint   `json:"userID"`
	OrderID  string `json:"orderID"`
	SeatIDs  uint   `json:"seatIDs"`
	PlanID   uint   `json:"planID"`
	CinemaID uint   `json:"cinemaID"`
}

type SearchTicketReq struct {
	OrderID string `json:"orderID"`
}

type BackTicketParam struct {
	OrderID string `json:"orderID"`
	PlanID  uint   `json:"planID"`
	SeatIDs []uint `json:"SeatIDs"`
}

type LockOneTicketReq struct {
	PlanID   uint `json:"planID"`
	CinemaID uint `json:"cinemaID"`
	SeatIDs  uint `json:"seatIDs"`
}
