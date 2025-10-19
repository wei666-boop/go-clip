package pkg

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"time"
)

var Logger *zerolog.Logger

func CreateLog(path string) *zerolog.Logger {
	f, err := os.OpenFile("./../log/"+path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil
	}

	logger := zerolog.New(f).With().Timestamp().Logger()

	Logger = &logger
	return Logger
}
func WriteInfoLog(path string, info string) {
	CreateLog(path).Info().Msg(info)
}

func WriteErrorLog(path string, error string) {
	CreateLog(path).Error().Msg(error)
}

func TerminalInfo(data string) {
	timeNow := time.Now()
	fmt.Fprintf(os.Stdout, "[%s]%s", timeNow, data)
}

func TerminalError(error string) {
	timeNow := time.Now()
	fmt.Fprintf(os.Stdout, "[%s]错误:%s", timeNow, error)
}
