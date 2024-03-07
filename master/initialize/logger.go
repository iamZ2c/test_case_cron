package initialize

import (
	"cron_tab_c/master/global"
	"go.uber.org/zap"
)

func Logger() {
	var (
		logger  *zap.Logger
		err     error
		slogger *zap.SugaredLogger
	)

	if logger, err = zap.NewDevelopment(); err != nil {
		return
	}
	defer logger.Sync()
	slogger = logger.Sugar()

	slogger.Info("初始化logger实例")
	global.SLogger = slogger
}
