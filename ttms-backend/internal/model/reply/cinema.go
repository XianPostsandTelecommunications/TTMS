/**
 * @Author: lenovo
 * @Description:
 * @File:  cinema
 * @Version: 1.0.0
 * @Date: 2023/06/04 2:58
 */

package reply

type CinemaInfo struct {
	ID     uint   `json:"ID"`
	Name   string `json:"name" `
	Avatar string `json:"avatar"`
	Rows   int    `json:"rows"`
	Cols   int    `json:"cols"`
}
