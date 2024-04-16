/**
 * @Author: lenovo
 * @Description:
 * @File:  carry
 * @Version: 1.0.0
 * @Date: 2023/06/06 21:29
 */

package FuncMgr

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"mognolia/internal/model/automigrate"
)

type carrier struct{}

func (c *carrier) DeleteOutTimeTicket(db *gorm.DB, planID, seatID, userID uint) error {
	if result := db.Model(&automigrate.Ticket{}).Where("plan_id =? and seat_id =? and user_id =?", planID, seatID, userID).
		Update("status", automigrate.ForSaleStatus); result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *carrier) DeleteOutTimeOrder(db *gorm.DB, orderID uuid.UUID) error {
	if result := db.Model(&automigrate.Order{}).Where("order_id = ?", orderID); result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}

func NewFuncMgr() *carrier {
	return &carrier{}
}
