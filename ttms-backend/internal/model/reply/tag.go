/**
 * @Author: lenovo
 * @Description:
 * @File:  tag
 * @Version: 1.0.0
 * @Date: 2023/05/31 19:47
 */

package reply

type GetTagFromMovie struct {
	Tags []string `json:"tags"`
}

type GetAllTagsParam struct {
	Tags []string `json:"tags"`
}
