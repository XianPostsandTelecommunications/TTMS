/**
 * @Author: lenovo
 * @Description:
 * @File:  FuncMgr
 * @Version: 1.0.0
 * @Date: 2023/06/06 21:22
 */

package FuncMgr

import "gorm.io/gorm"

type ManagerFunc interface {
	DeleteOutTimeTicket(db *gorm.DB, ticketID, planID, seatID uint) error
}
