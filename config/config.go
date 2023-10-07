package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
	"strconv"
)

var DB *gorm.DB

type Config struct {
	ServerPort    int
	DBPort        int
	DBHost        string
	DBUser        string
	DBPass        string
	DBName        string
	Secret        string
	RefreshSecret string
}

func InitConfig() *Config {
	var res = new(Config)
	res = loadConfig()

	if res == nil {
		logrus.Fatal("Config : cannot load configuration")
		return nil
	}
	return res

}

func loadConfig() *Config {

	var res = new(Config)
	err := godotenv.Load(".env")

	if err != nil {
		logrus.Fatal("Config : cannot load config file", err.Error())
		return nil
	}

	if value, found := os.LookupEnv("SERVER"); found {
		port, err := strconv.Atoi(value)
		if err != nil {
			logrus.Fatal("Config : invalid server port", err.Error())
			return nil
		}
		res.ServerPort = port
	}

	if value, found := os.LookupEnv("DBPORT"); found {
		port, err := strconv.Atoi(value)
		if err != nil {
			logrus.Fatal("Config : invalid db port", err.Error())
			return nil
		}
		res.DBPort = port

	}
	if value, found := os.LookupEnv("DBHOST"); found {
		res.DBHost = value
	}
	if value, found := os.LookupEnv("DBUSER"); found {
		res.DBUser = value
	}
	if value, found := os.LookupEnv("DBPASS"); found {
		res.DBPass = value
	}
	if value, found := os.LookupEnv("DBNAME"); found {
		res.DBName = value
	}
	if val, found := os.LookupEnv("SECRET"); found {
		res.Secret = val
	}
	if val, found := os.LookupEnv("REFSECRET"); found {
		res.RefreshSecret = val
	}
	return res
}
