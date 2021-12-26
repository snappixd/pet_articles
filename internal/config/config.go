package config

import (
	"time"
)

const (
	defaultMongoURI      = "uri"
	defaultMongoUser     = "user"
	defaultMongoPassword = "password"
	defaultMongoDb       = "dbName"

	defaultHTTPHost      = "127.0.0.1"
	defaultHTTPPort      = "80"
	defaultHTTPRWTimeout = 10 * time.Second
)

type (
	Config struct {
		Mongo MongoConfig
		HTTP  HTTPConfig
	}

	MongoConfig struct {
		URI      string
		User     string
		Password string
		DbName   string
	}

	HTTPConfig struct {
		Host         string
		Port         string
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}
)

func Init() *Config {
	var cfg Config

	populateDefaults(&cfg)

	return &cfg
}

func populateDefaults(cfg *Config) {
	cfg.Mongo.URI = defaultMongoURI
	cfg.Mongo.User = defaultMongoUser
	cfg.Mongo.Password = defaultMongoPassword
	cfg.Mongo.DbName = defaultMongoDb

	cfg.HTTP.Port = defaultHTTPPort
	cfg.HTTP.Host = defaultHTTPHost
}

// func setFromEnv(cfg *Config) {
// 	cfg.Mongo.URI = os.Getenv("MONGO_URI")
// 	cfg.Mongo.User = os.Getenv("MONGO_USER")
// 	cfg.Mongo.Password = os.Getenv("MONGO_PASSWORD")
// 	cfg.Mongo.DbName = os.Getenv("MONGO_DB")
//
// 	cfg.HTTP.Host = os.Getenv("HTTP_HOST")
// 	cfg.HTTP.Port = os.Getenv("HTTP_PORT")
// }
