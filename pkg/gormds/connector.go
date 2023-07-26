package gormds

import (
	"fmt"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ConnectorOptions struct {
	Logger *zerolog.Logger
}

type Connector struct {
	o ConnectorOptions

	config *GORMConfig
}

func NewConnector(configGetter GORMConfigGetter) *Connector {
	return &Connector{
		config: configGetter.GetGORMConfig(),
	}
}

func (c Connector) DB() (*gorm.DB, error) {
	return c.makeDb(c.config.DefaultConnection)
}

func (c Connector) DBWith(name string) (*gorm.DB, error) {
	return c.makeDb(name)
}

func (c Connector) makeDb(name string) (*gorm.DB, error) {
	connCfg, ok := c.config.Connections[name]
	if !ok {
		return nil, ErrInvalidConnection
	}

	dialector, err := dialector(connCfg)
	if err != nil {
		return nil, err
	}

	loggerConfig, err := c.loggerConfig()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger:  logger.New(c.o.Logger, loggerConfig),
		NowFunc: func() time.Time { return time.Now().UTC() },
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (c Connector) loggerConfig() (logger.Config, error) {
	slowThreshold, err := time.ParseDuration(c.config.Logging.SlowThreshold)
	if err != nil {
		return logger.Config{}, err
	}

	return logger.Config{
		IgnoreRecordNotFoundError: c.config.Logging.IgnoreRecordNotFoundError,
		SlowThreshold:             slowThreshold,
		LogLevel:                  logLevel(c.config.Logging.SlowThreshold),
	}, nil
}

func dialector(connCfg ConnectionConfig) (gorm.Dialector, error) {
	switch connCfg.Driver {
	case "mysql":
		return mysql.Open(mysqlDsn(
			connCfg.Username,
			connCfg.Password,
			connCfg.Database,
			connCfg.Host,
			connCfg.Port,
			map[string]string{
				"charset":   "utf8mb4",
				"parseTime": "true",
			}),
		), nil
	default:
		return nil, ErrInvalidDriver
	}
}

func logLevel(l string) logger.LogLevel {
	switch strings.ToUpper(l) {
	case "ERROR":
		return logger.Error
	case "WARN":
		return logger.Warn
	case "INFO":
		return logger.Info
	case "SILENT":
		fallthrough
	default:
		return logger.Silent
	}
}

func mysqlDsn(username, password, dbname, host string, port int, options map[string]string) string {
	queries := make([]string, 0, len(options))
	for key, value := range options {
		queries = append(queries, fmt.Sprintf("%s=%s", key, value))
	}
	if len(queries) > 0 {
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", username, password, host, port, dbname, strings.Join(queries, "&"))
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, dbname)
}
