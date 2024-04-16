/**
 * @Author: lenovo
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2023/05/29 10:38
 */

package automigrate

import "gorm.io/gorm"

type Gend string

var (
	Male    Gend = "male"
	Female  Gend = "female"
	UnKnown Gend = "unknown"
)

type Roler string

var (
	Admin  Roler = "administer"
	Vistor Roler = "vistor"
)

type User struct {
	gorm.Model
	UserName  string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"type:varchar(20);index:idx_name,unique"`
	Password  string `gorm:"type:varchar(255);not null"`
	Signature string `gorm:"type:varchar(255);not null"`
	Avatar    string `gorm:"type:varchar(255);default:http://lycmall.lyc666.xyz/lycmall2/2023-04-17-22:58:18.48.jpg"`
	Gender    Gend   `gorm:"type:varchar(10);not null"`
	Role      Roler  `gorm:"type:varchar(20);not null"`
}
