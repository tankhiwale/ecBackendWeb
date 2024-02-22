package database

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib" // just initialize it
	"github.com/tankhiwale/ecBackendWeb/config"
	"github.com/tankhiwale/ecBackendWeb/internal/logger"
)

func InitializeDatabase(config *config.DatabaseConfig, logger logger.ILogger) (*pgxpool.Pool, error) {

	logger.Info("Initializing database")

	get := func(configName string) string {
		var value string
		switch configName {
		case "host":
			value = config.Host
		case "port":
			value = config.Port
		case "username":
			value = config.Username
		case "password":
			value = config.Password
		case "databaseName":
			value = config.DatabaseName
		case "sslMode":
			value = config.SslMode
		default:
			value = ""
		}
		if value == "" {
			logger.Error("database config read error")
			log.Fatalf("database config cannot be blank. ensure you have provided all the required values.")
		}
		return value
	}
	dataSourceName := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(get("username"), get("password")),
		Host:   fmt.Sprintf("%s:%s", get("host"), get("port")),
		Path:   get("databaseName"),
	}
	q := dataSourceName.Query()

	q.Add("sslmode", get("sslMode"))
	dataSourceName.RawQuery = q.Encode()

	pool, err := pgxpool.New(context.Background(), dataSourceName.String())
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return pool, nil

}
