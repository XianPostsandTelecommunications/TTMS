package global

import (
	"mognolia/internal/model/config"
	"mognolia/internal/pkg/goroutine/work"
	"mognolia/internal/pkg/logger"
	"mognolia/internal/pkg/token"

	ut "github.com/go-playground/universal-translator"
)

var (
	RootDir  string
	Settings config.AllConfig
	Logger   *logger.Log   // 日志
	Trans    ut.Translator //翻译器
	Worker   *work.Worker
	Maker    token.Maker
)
