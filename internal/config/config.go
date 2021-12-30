package config

import (
	"time"

	"github.com/spf13/viper"
)

const (
	defaultMongoURI      = "mongodb+srv://admin:admin@cluster0.cweoh.mongodb.net/articlesDb?retryWrites=true&w=majority"
	defaultMongoUser     = "admin"
	defaultMongoPassword = "admin"
	defaultMongoDb       = "articlesDb"

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

func Init(cfgDir string) (*Config, error) {
	populateDefaults()

	if err := parseConfigFile(cfgDir); err != nil {
		return nil, err
	}

	var cfg Config

	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func populateDefaults() {
	viper.SetDefault("http.host", defaultHTTPHost)
	viper.SetDefault("http.port", defaultHTTPPort)
	viper.SetDefault("http.readTimeout", defaultHTTPRWTimeout)
	viper.SetDefault("http.writeTimeout", defaultHTTPRWTimeout)

	viper.SetDefault("mongo.uri", defaultMongoURI)
	viper.SetDefault("mongo.user", defaultMongoUser)
	viper.SetDefault("mongo.password", defaultMongoPassword)
	viper.SetDefault("mongo.dbName", defaultMongoDb)
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("http.host", &cfg.HTTP.Host); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("http.port", &cfg.HTTP.Port); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("http.readTimeout", &cfg.HTTP.ReadTimeout); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("http.writeTimeout", &cfg.HTTP.WriteTimeout); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("mongo.uri", &cfg.Mongo.URI); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("mongo.user", &cfg.Mongo.User); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("mongo.password", &cfg.Mongo.Password); err != nil {
		return err
	}

	return viper.UnmarshalKey("mongo.dbName", &cfg.Mongo.DbName)
}

func parseConfigFile(folder string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.MergeInConfig()
}
