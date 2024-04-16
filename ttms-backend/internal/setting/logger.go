/**
 * @Author: lenovo
 * @Description:
 * @File:  logger
 * @Version: 1.0.0
 * @Date: 2023/05/29 8:55
 */

package setting

import (
	"go.uber.org/zap"
	"mognolia/internal/global"
	"mognolia/internal/pkg/logger"
)

type log struct{}

func (log) Init() {
	global.Logger = logger.NewLogger(&logger.InitStruct{
		LogSavePath:   global.Settings.Log.LogSavePath,
		LogFileExt:    global.Settings.Log.LogFileExt,
		MaxSize:       global.Settings.Log.MaxSize,
		MaxBackups:    global.Settings.Log.MaxBackups,
		MaxAge:        global.Settings.Log.MaxAge,
		Compress:      global.Settings.Log.Compress,
		LowLevelFile:  global.Settings.Log.LowLevelFile,
		HighLevelFile: global.Settings.Log.HighLevelFile,
	}, global.Settings.Log.Level)
	zap.ReplaceGlobals(global.Logger.Logger)
}
