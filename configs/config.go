package configs

import (
	"errors"
	"github.com/invopop/validation"
	"github.com/spf13/viper"
)

const ApiVersion = 1

var cfg *config

type config struct {
	API   APIConfig
	Redis RedisConfig
	Log   LogConfig
}

type APIConfig struct {
	Port                       string
	Environment                string
	Timezone                   string
	RateLimitRequest           int
	RateLimitTimeSecond        int
	RateLimitRequestByToken    int
	RateLimitTimeSecondByToken int
	RateLimitTimeBlock         int
}
type RedisConfig struct {
	Host     string
	Port     int64
	Password string
	DB       int
}
type LogConfig struct {
	Dir  string
	File bool
	DB   bool
}

func Load(path string) error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			panic(err)
		}
	}

	apiConfig := APIConfig{
		Port:                       viper.GetString("api.port"),
		Environment:                viper.GetString("api.environment"),
		Timezone:                   viper.GetString("api.timezone"),
		RateLimitRequest:           viper.GetInt("api.rate_limit_request"),
		RateLimitRequestByToken:    viper.GetInt("api.rate_limit_request_by_token"),
		RateLimitTimeSecond:        viper.GetInt("api.rate_limit_time_second"),
		RateLimitTimeSecondByToken: viper.GetInt("api.rate_limit_time_second_by_token"),
		RateLimitTimeBlock:         viper.GetInt("api.rate_limit_time_block_second"),
	}

	if apiErr := apiConfig.Validate(); apiErr != nil {
		return apiErr
	}

	redisConfig := RedisConfig{
		Host:     viper.GetString("redis.host"),
		Port:     viper.GetInt64("redis.port"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	}

	if redisErr := redisConfig.Validate(); redisErr != nil {
		return redisErr
	}

	logConfig := LogConfig{
		Dir:  viper.GetString("logs.dir"),
		File: viper.GetBool("logs.file"),
		DB:   viper.GetBool("logs.db"),
	}

	if logErr := logConfig.Validate(); logErr != nil {
		return logErr
	}

	cfg = new(config)
	cfg.API = apiConfig
	cfg.Redis = redisConfig
	cfg.Log = logConfig

	return nil
}

func GetAPIConfig() APIConfig {
	return cfg.API
}

func GetRedisConfig() RedisConfig {
	return cfg.Redis
}

func GetLogConfig() LogConfig {
	return cfg.Log
}

func (a APIConfig) Validate() error {
	return validation.ValidateStruct(
		&a,
		validation.Field(&a.Environment, validation.Required),
		validation.Field(&a.Port, validation.Required),
		validation.Field(&a.Timezone, validation.Required),
		validation.Field(&a.RateLimitRequest, validation.Required),
		validation.Field(&a.RateLimitTimeSecond, validation.Required),
		validation.Field(&a.RateLimitRequestByToken, validation.Required),
		validation.Field(&a.RateLimitTimeSecondByToken, validation.Required),
		validation.Field(&a.RateLimitTimeBlock, validation.Required),
	)
}

func (r RedisConfig) Validate() error {
	return validation.ValidateStruct(
		&r,
		validation.Field(&r.Host, validation.Required),
		validation.Field(&r.Port, validation.Required),
		validation.Field(&r.DB, validation.Required),
	)
}

func (l LogConfig) Validate() error {
	return validation.ValidateStruct(
		&l,
		validation.Field(&l.Dir, validation.Required),
		validation.Field(&l.File, validation.NotNil),
		validation.Field(&l.DB, validation.NotNil),
	)
}
