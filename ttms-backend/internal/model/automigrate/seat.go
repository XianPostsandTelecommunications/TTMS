/**
 * @Author: lenovo
 * @Description:
 * @File:  seat
 * @Version: 1.0.0
 * @Date: 2023/06/04 6:04
 */

package automigrate

import "gorm.io/gorm"

type SeatStatus string

const (
	SeatLockStatus SeatStatus = "broken"
	SeatNormal     SeatStatus = "normal"
)

type Seat struct {
	gorm.Model
	CinemaID uint
	Cinema   Cinema     `gorm:"foreignKey:CinemaID;references:ID"`
	Row      int        `gorm:"type:int;not null"`
	Col      int        `gorm:"type:int;not null"`
	Status   SeatStatus `gorm:"type:varchar(15)"`
}
