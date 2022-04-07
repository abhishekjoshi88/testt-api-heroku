package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const Time = "15:04:05"

// DbConfig hold the values required to connect with the database
type DbConfig struct {
	DbType     string `mapstructure:"DB_TYPE"`
	DbUsername string `mapstructure:"DB_USERNAME"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbName     string `mapstructure:"DB_NAME"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbURL      string `mapstructure:"DatabaseURL"`
}

// AppConfiguration represent configuration for this application
type AppConfiguration struct {
	Environment string `mapstructure:"APPLICATION_ENVIRONMENT"`
	Port        string `mapstructure:"APPLICATION_PORT"`
}

// GetDbConfig gets the database configuration vairbales from the config.yml file, then overrides them with the environemtn values for the same (if any). Is also responsible
// for creating the database connection url
func GetDbConfig() (config DbConfig, err error) {
	// Set configuration file which will be used to get/set config values
	viper.SetConfigFile(`config.yml`)
	// Ask viper to overwrite any configuration values with their corresponding enviornment counterparts
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// Load the database configuration in the database struct
	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}
	config.DbURL, err = GenerateDatabaseURL(config.DbType, config.DbUsername, config.DbPassword, config.DbHost, config.DbPort, config.DbName)
	return config, err
}

// GenerateDatabaseURL will generate the url which will be used by our connector
func GenerateDatabaseURL(databaseType string, username string, password string, host string, port string, databaseName string) (string, error) {
	if databaseType == "mysql" {
		return fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, databaseName), nil
	}
	return "", fmt.Errorf("Invalid Database Type: %s", databaseType)
}

// GetAppConfig will load the application config
func GetAppConfig() (appConfig AppConfiguration, err error) {
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// Load the app configuration in the AppConfiguration struct
	err = viper.Unmarshal(&appConfig)
	if err != nil {
		return
	}

	return
}
