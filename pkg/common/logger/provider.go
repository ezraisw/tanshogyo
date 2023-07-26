package logger

import (
	"io"
	"runtime"
	"strconv"
	"strings"

	"github.com/rs/zerolog"
)

func ProvideLogger(configGetter LoggerConfigGetter, w io.Writer) (*zerolog.Logger, error) {
	level, err := zerolog.ParseLevel(strings.ToLower(configGetter.GetLoggerConfig().Level))
	if err != nil {
		return nil, err
	}
	logger := zerolog.New(w).
		Hook(zerolog.HookFunc(addCallerInfoHook)).
		Level(level)
	return &logger, nil
}

func addCallerInfoHook(e *zerolog.Event, level zerolog.Level, message string) {
	pc, file, line, ok := runtime.Caller(4)
	if !ok {
		return
	}
	fileName := file[strings.LastIndex(file, "/")+1:] + ":" + strconv.Itoa(line)
	funcName := runtime.FuncForPC(pc).Name()
	e.Fields(map[string]any{
		"fileName": fileName,
		"funcName": funcName,
	})
}
