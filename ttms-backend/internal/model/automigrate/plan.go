/**
 * @Author: lenovo
 * @Description:
 * @File:  plan
 * @Version: 1.0.0
 * @Date: 2023/06/04 4:06
 */

package automigrate

import (
	"gorm.io/gorm"
	"time"
)

type Plan struct {
	gorm.Model
	MovieID uint
	Movie   Movie `gorm:"foreignKey:MovieID;references:ID"`

	CinemaID uint
	Cinema   Cinema `gorm:"foreignKey:CinemaID;references:ID"`

	StartAt time.Time
	EndAt   time.Time
	Price   float64
}
