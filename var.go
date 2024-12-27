package ncmdl

import (
	"github.com/caiknife/mp3lister/lib/logger"
	"github.com/sirupsen/logrus"
)

var DefaultPoolSize = PoolSize

var AppLogger *logrus.Logger

func init() {
	AppLogger = logger.NewConsoleLogger(
		logger.DisableTimestamp(true),
	)
}
