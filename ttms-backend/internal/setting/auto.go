/**
 * @Author: lenovo
 * @Description:
 * @File:  auto
 * @Version: 1.0.0
 * @Date: 2023/06/03 20:25
 */

package setting

import "mognolia/internal/logic"

type auto struct{}

func (auto) Init() {
	logic.Group.Auto.Work()
}
