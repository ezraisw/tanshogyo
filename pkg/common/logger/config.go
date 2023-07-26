package logger

type LoggerConfigGetter interface {
	GetLoggerConfig() *LoggerConfig
}

type LoggerConfig struct {
	Level string
}
