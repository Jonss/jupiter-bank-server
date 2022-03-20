package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Env             string `mapstruct:"ENV"`
	Port            string `mapstructure:"PORT"`
	APPMessage      string `mapstructure:"APP_MESSAGE"`
	DBDriver        string `mapstructure:"DATABASE_DRIVER"`
	DBURL           string `mapstructure:"DATABASE_URL"`
	DBName          string `mapstructure:"DATABASE_NAME"`
	DBMigrationPath string `mapstructure:"DATABASE_MIGRATION_PATH"`
}

func LoadConfig(path string) (Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	var config Config
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return config, nil
		}
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}
	return config, nil
}
