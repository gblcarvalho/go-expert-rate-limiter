package configs

import (
	"github.com/spf13/viper"
)

type configuration struct {
	IPMaxRequests    int64 `mapstructure:"RATE_LIMITE_IP_MAX_REQUESTS"`
	TokenMaxRequests int64 `mapstructure:"RATE_LIMITE_TOKEN_MAX_REQUESTS"`
	TimeWindow       int64 `mapstructure:"RATE_LIMITE_TIME_WINDOW"`

	RedisAddr string `mapstructure:"REDIS_ADDR"`
}

func LoadConfig(path string, fileName string) (*configuration, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(fileName)
	viper.AutomaticEnv()
	viper.ReadInConfig()

	var cfg *configuration
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
