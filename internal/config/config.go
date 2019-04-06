package config

import (
	"github.com/globalsign/mgo"
	"github.com/spf13/viper"
)

// Constants is the type defining the config constants structure
type Constants struct {
	PORT  string
	Mongo struct {
		URL    string
		DBName string
	}
}

// Config is the type defining the config object structure
type Config struct {
	Constants
	Database *mgo.Database
}

// New generates a Config instance which can be passed around the codebase
func New() (*Config, error) {
	conf := Config{}

	constants, err := initViper()
	conf.Constants = constants
	if err != nil {
		return &conf, err
	}

	db, err := mgo.Dial(conf.Constants.Mongo.URL)
	if err != nil {
		return &conf, err
	}

	conf.Database = db.DB(conf.Constants.Mongo.DBName)
	return &conf, nil
}

func initViper() (Constants, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return Constants{}, err
	}

	viper.SetDefault("PORT", "8080")

	var constants Constants
	viper.Unmarshal(&constants)

	return constants, err
}
