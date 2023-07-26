package gormds

type GORMConfigGetter interface {
	GetGORMConfig() *GORMConfig
}

type GORMConfig struct {
	Logging           LoggingConfig
	DefaultConnection string
	Connections       map[string]ConnectionConfig
}

type ConnectionConfig struct {
	Driver   string
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

type LoggingConfig struct {
	SlowThreshold             string
	IgnoreRecordNotFoundError bool
	LogLevel                  string
}
