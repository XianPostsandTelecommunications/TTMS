/**
 * @Author: lenovo
 * @Description:
 * @File:  common
 * @Version: 1.0.0
 * @Date: 2023/05/30 9:31
 */

package query

import (
	"gorm.io/gorm"
)

func Paginate(page, pageSize int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}
