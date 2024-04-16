/**
 * @Author: lenovo
 * @Description:
 * @File:  order
 * @Version: 1.0.0
 * @Date: 2023/06/07 18:33
 */

package query

import (
	"gorm.io/gorm"
	"mognolia/internal/dao"
	"mognolia/internal/model/automigrate"
)

type order struct{}

func NewOrder() *order {
	return &order{}
}

func (order) GetOrderInfoByOrderID(orderID string) (*automigrate.Order, error) {
	var orderInfo automigrate.Order
	if result := dao.Group.DB.Model(&automigrate.Order{}).Where("order_id = ?", orderID).Find(&orderInfo); result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &orderInfo, nil
}
