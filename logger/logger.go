package logger

import (
	"bytes"

	"github.com/buguang01/util"
	logs2 "github.com/huanglishi/gofly-mqttv5/logger/logs"
)

// buffer holds a byte Buffer for reuse. The zero value is ready for use.
type buffer struct {
	bytes.Buffer
	tmp  [64]byte // temporary byte array for creating headers.
	next *buffer
}

var Logger *logs2.AdamLog

func LogInit(level string) {
	util.SetLocation(util.BeiJing)
	logs2.LogInit(level)
	Logger = logs2.GetLogger()
}
