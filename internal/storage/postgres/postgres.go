package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
	"log"
	"net/url"
	"time"
)

var Module = fx.Module("postgres", fx.Provide(NewPostgresStorage, NewPostgresConfig, NewDataSourceUrl))

type Storage struct {
	db *sql.DB
}

func NewPostgresStorage(lc fx.Lifecycle, url DataSourceUrl) (*Storage, error) {
	var db *sql.DB
	for {
		db, err := sql.Open("postgres", string(url))
		if err != nil {
			return nil, err
		}
		log.Println("waiting for db")
		err = db.Ping()
		if err == nil {
			break
		}
		time.Sleep(time.Second)
		continue
	}
	return &Storage{db: db}, nil
}

type Config struct {
	Proto    string
	Host     string
	Path     string
	User     string
	Password string
	SSLMode  string
}

func NewPostgresConfig(lc fx.Lifecycle) *Config {
	return &Config{Proto: "postgres", Host: "db:5432", Path: "postgres", User: "postgres", Password: "1234", SSLMode: "disable"}
}

type DataSourceUrl string

func NewDataSourceUrl(lc fx.Lifecycle, config *Config) DataSourceUrl {
	const SSLMODE = "sslmode"
	var v = make(url.Values)
	v.Add(SSLMODE, config.SSLMode)

	var u = url.URL{
		Scheme:   config.Proto,
		Host:     config.Host,
		Path:     config.Path,
		User:     url.UserPassword(config.User, config.Password),
		RawQuery: v.Encode(),
	}
	return DataSourceUrl(u.String())
}
