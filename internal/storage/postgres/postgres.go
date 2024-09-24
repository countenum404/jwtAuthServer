package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"log"
	"net/url"
	"time"
)

var Module = fx.Module("postgres", fx.Provide(NewPostgresStorage, LoadNewPostgresConfig, NewDataSourceUrl))

type Storage struct {
	db *sql.DB
}

func NewPostgresStorage(lc fx.Lifecycle, url DataSourceUrl) (*Storage, error) {
	var db *sql.DB
	log.Println(string(url))
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

func LoadNewPostgresConfig(lc fx.Lifecycle) (*Config, error) {
	config := &Config{}

	viper.SetConfigFile("config/postgres_config.yml")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}
	return config, nil
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
