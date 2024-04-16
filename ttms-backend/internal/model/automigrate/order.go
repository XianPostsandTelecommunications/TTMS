/**
 * @Author: lenovo
 * @Description:
 * @File:  order
 * @Version: 1.0.0
 * @Date: 2023/06/07 11:27
 */

package automigrate

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderStatus string

var (
	OrderNotPaid OrderStatus = "未支付"
	OrderHasPaid OrderStatus = "已支付"
)

type Order struct {
	gorm.Model
	OrderID    uuid.UUID `gorm:"type:varchar(100);not null"`
	PlanID     uint      `gorm:"type:int;not null"`
	UserID     uint      `gorm:"type:int;not null"`
	MovieID    uint
	Movie      Movie       `gorm:"foreignKey:MovieID;references:ID"`
	UserName   string      `gorm:"type:varchar(20);not null"`
	CinemaName string      `gorm:"type:varchar(30);not null"`
	Status     OrderStatus `gorm:"type:varchar(20)"`
	Seats      string      `gorm:"type:varchar(1000);not null"`
	Price      float32     `gorm:"type:float"`
	SeatsID    string      `gorm:"type:varchar(200);not null"`
}
