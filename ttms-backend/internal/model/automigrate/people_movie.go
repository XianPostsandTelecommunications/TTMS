/**
 * @Author: lenovo
 * @Description:
 * @File:  people_movie
 * @Version: 1.0.0
 * @Date: 2023/06/01 22:58
 */

package automigrate

import "gorm.io/gorm"

type PeopleMovie struct {
	gorm.Model
	MovieID uint
	Movie   Movie `gorm:"foreignKey:MovieID;references:ID"`

	UserID uint
	User   User `gorm:"foreignKey:UserID;references:ID"`
}
