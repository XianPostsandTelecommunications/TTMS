/**
 * @Author: lenovo
 * @Description:
 * @File:  manager_movie
 * @Version: 1.0.0
 * @Date: 2023/05/31 9:28
 */

package automigrate

import "gorm.io/gorm"

type ManagerMovie struct {
	gorm.Model
	MovieID uint

	UserID uint

	Content string `gorm:"type:varchar(1000)"`
}
