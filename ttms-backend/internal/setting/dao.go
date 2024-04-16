/**
 * @Author: lenovo
 * @Description:
 * @File:  dao
 * @Version: 1.0.0
 * @Date: 2023/05/29 10:33
 */

package setting

import (
	"mognolia/internal/dao"
	"mognolia/internal/dao/mysql"
	"mognolia/internal/dao/redis"
	"mognolia/internal/global"
)

type mdao struct{}

func (d *mdao) Init() {
	mysql.InitMySql()
	dao.Group.Redis = redis.Init(global.Settings.Redis.Addr, global.Settings.Redis.Password, global.Settings.Redis.PoolSize, 0)
}
