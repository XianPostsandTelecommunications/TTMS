/**
 * @Author: lenovo
 * @Description:
 * @File:  plan
 * @Version: 1.0.0
 * @Date: 2023/06/04 5:17
 */

package reply

import "time"

type CreatePlanRly struct {
	PlanID  uint      `json:"planID"`
	StartAt time.Time `json:"startAt"`
	EndAt   time.Time `json:"endAt"`
}

type GetPlanByMovieIDAndPeriod struct {
	PlansInfo []*Plans
	Total     int64
}

type Plans struct {
	MovieName   string    `json:"movieName"`
	MovieAvatar string    `json:"movieAvatar"`
	CinemaName  string    `json:"cinemaName"`
	StartAt     time.Time `json:"startAt"`
	EndAt       time.Time `json:"endAt"`
}

type PlanList struct {
	Data  []PlanInfo
	Total int64
}

type PlanInfo struct {
	ID          uint      `json:"id"`
	SeatID      uint      `json:"seat_id"`
	MovieName   string    `json:"movieName"`
	MovieAvatar string    `json:"movieAvatar"`
	CinemaID    uint      `json:"cinema_id"`
	CinemaName  string    `json:"cinemaName"`
	Duration    int64     `json:"duration"`
	StartAt     time.Time `json:"startAt"`
	Price       float64   `json:"price"`
}
