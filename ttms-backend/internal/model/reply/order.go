/**
 * @Author: lenovo
 * @Description:
 * @File:  order
 * @Version: 1.0.0
 * @Date: 2023/06/07 12:22
 */

package reply

import "github.com/google/uuid"

type RlyOrderID struct {
	OrderID uuid.UUID `json:"orderID"`
}

type CreateOrderRly struct {
	IsLock  bool
	OrderID uuid.UUID `json:"orderID"`
}
