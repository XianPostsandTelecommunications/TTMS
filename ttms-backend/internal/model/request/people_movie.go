/**
 * @Author: lenovo
 * @Description:
 * @File:  people_movie
 * @Version: 1.0.0
 * @Date: 2023/06/01 22:30
 */

package request

type UserMovie2Action struct {
	MovieID uint `json:"movieID" binding:"required"`
	Choice  int  `json:"choice" binding:"required,oneof=1 2"` //1表示关注了这电影,2表示取消对这电影的关注
}
