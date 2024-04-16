/**
 * @Author: lenovo
 * @Description:
 * @File:  movie
 * @Version: 1.0.0
 * @Date: 2023/06/04 1:13
 */

package automigrate

import "gorm.io/gorm"

type Cinema struct {
	gorm.Model
	Avatar string `gorm:"type:varchar(1000);not null"`
	Name   string `gorm:"type:varchar(100)"`
	Rows   int    `gorm:"type:int;not null"`
	Cols   int    `gorm:"type:int;not null"`
}
