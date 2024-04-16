/**
 * @Author: lenovo
 * @Description:
 * @File:  tag
 * @Version: 1.0.0
 * @Date: 2023/05/31 8:39
 */

package automigrate

import "gorm.io/gorm"

type TagType []string
type Tag struct {
	gorm.Model
	MovieID uint
	Movie   Movie   `gorm:"foreignKey:MovieID;references:ID"`
	Tags    TagType `gorm:"type:varchar(255);not null"`
}
