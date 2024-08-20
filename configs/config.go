package configs

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type ServerConf struct {
	Port int
}

type DataBaseConf struct {
	Host     string
	Port     string
	User     string
	Password string
	DBname   string
}

type Config struct {
	Server   ServerConf
	Database DataBaseConf
}

func GetDBConfig() DataBaseConf {
	return DataBaseConf{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBname:   os.Getenv("DB_NAME"),
	}
}

func GetConfig() Config {
	serverPort, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("There is no server port in env")
	}

	serv := ServerConf{
		Port: serverPort,
	}

	return Config{Server: serv, Database: GetDBConfig()}
}

func GetDBUrl() string {
	dbConfig := GetDBConfig()
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBname)
}
