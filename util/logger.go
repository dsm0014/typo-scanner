package util

import (
	"io"
	"log"
)

func GetLogger(suppress bool) *log.Logger {
	if suppress {
		return log.New(io.Discard, "", 0)
	}
	return log.Default()
}
