/**
 * @Author: lenovo
 * @Description:
 * @File:  tag
 * @Version: 1.0.0
 * @Date: 2023/05/31 18:56
 */

package request

type AddTagParam struct {
	MovieID uint     `json:"movieID"`
	Tag     []string `json:"tag"`
}

type GetTagsParam struct {
	MovieID uint `json:"movieID"`
}
