/**
 * @Author: lenovo
 * @Description:
 * @File:  enter
 * @Version: 1.0.0
 * @Date: 2023/05/29 9:57
 */

package dao

import (
	"gorm.io/gorm"
	"mognolia/internal/dao/redis/query"
)

type group struct {
	DB    *gorm.DB
	Redis *query.Queries
}

var Group = new(group)
