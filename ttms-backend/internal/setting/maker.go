/**
 * @Author: lenovo
 * @Description:
 * @File:  maker
 * @Version: 1.0.0
 * @Date: 2023/05/29 18:40
 */

package setting

import (
	"mognolia/internal/global"
	"mognolia/internal/pkg/token"

	"go.uber.org/zap"
)

type maker struct {
}

func (maker) Init() {
	var err error
	global.Maker, err = token.NewPasetoMaker([]byte(global.Settings.Token.Key))
	if err != nil {
		zap.S().Infof("token.NewPasetoMaker failed: %v", err)
		return
	}
}
