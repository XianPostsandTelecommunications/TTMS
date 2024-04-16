/**
 * @Author: lenovo
 * @Description:
 * @File:  ticket
 * @Version: 1.0.0
 * @Date: 2023/06/04 6:06
 */

package automigrate

import (
	"gorm.io/gorm"
	"time"
)

type TicketStatus string

const (
	LockStatus    TicketStatus = "lock"
	ForSaleStatus TicketStatus = "for_sale"
	SaledStatus   TicketStatus = "saled"
)

type Ticket struct {
	gorm.Model
	PlanID       uint
	Plan         Plan `gorm:"foreignKey:PlanID;references:ID"`
	UserID       uint `gorm:"index"`
	SeatID       uint
	Seat         Seat         `gorm:"foreignKey:SeatID;references:ID"`
	Price        float64      `gorm:"type:float;not null"`
	TicketStatus TicketStatus `gorm:"type:varchar(100)"`
	LockTime     *time.Time
}
