package logger

import (
	"log/slog"
)

func LogOnError(err error, msg string) {
	if err != nil {
		slog.Error(msg + ": " + err.Error())
	}
}
