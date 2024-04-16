/**
 * @Author: lenovo
 * @Description:
 * @File:  manager
 * @Version: 1.0.0
 * @Date: 2023/06/08 5:31
 */

package logic

var TicketMap = make(chan bool, 1) //订单
var SeatMap = make(chan bool, 1)   //座位
