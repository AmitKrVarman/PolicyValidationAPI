package dbclient

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"flag"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/AmitKrVarman/PolicyValidationAPI/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	ENV_DB_DRIVER         = "DB_DRIVER"
	ENV_DB_HOST           = "DB_HOST"
	ENV_DB_PORT           = "DB_PORT"
	ENV_DB_NAME           = "DB_NAME"
	ENV_DB_USER           = "DB_USER"
	ENV_DB_PASSWORD       = "DB_PASSWORD"
	ENV_DB_SSL_MODE       = "DB_SSL_MODE"
	ENV_DB_MAX_CONNECTION = "DB_MAX_CONNECTION"
	ENV_DB_LOG_MODE       = "DB_LOG_MODE"
)

//Default values
var (
	DB_DRIVER         string = "postgres"
	DB_HOST           string = "localhost"
	DB_PORT           string = "26257"
	DB_NAME           string = "POLICYDB"
	DB_USER           string = "dbuser"
	DB_PASSWORD       string = "Password123"
	DB_SSL_MODE       string = "disable" // disable | require
	DB_MAX_CONNECTION int    = 1
	DB_LOG_MODE       bool   = true
)

var ONCE sync.Once
var DBSession *gorm.DB

func init() {
	getEnvDatabaseConfig()
	GetInstance()
}

func getEnvDatabaseConfig() {
	log.Print("[CONFIG] Reading Env variables")
	dbDriver := os.Getenv(ENV_DB_DRIVER)
	dbHost := os.Getenv(ENV_DB_HOST)
	dbPort := os.Getenv(ENV_DB_PORT)
	dbName := os.Getenv(ENV_DB_NAME)
	dbUser := os.Getenv(ENV_DB_USER)
	dbPassword := os.Getenv(ENV_DB_PASSWORD)
	dbSslMode := os.Getenv(ENV_DB_SSL_MODE)
	dbMaxConnection := os.Getenv(ENV_DB_MAX_CONNECTION)
	dbLogMode := os.Getenv(ENV_DB_LOG_MODE)
	maxConnection, err1 := strconv.Atoi(dbMaxConnection)
	logMode, err2 := strconv.ParseBool(dbLogMode)

	if len(dbDriver) > 0 {
		DB_DRIVER = dbDriver
	}
	if len(dbHost) > 0 {
		DB_HOST = dbHost
	}
	if len(dbPort) > 0 {
		DB_PORT = dbPort
	}
	if len(dbName) > 0 {
		DB_NAME = dbName
	}
	if len(dbUser) > 0 {
		DB_USER = dbUser
	}
	if len(dbPassword) > 0 {
		DB_PASSWORD = dbPassword
	}
	if len(dbSslMode) > 0 {
		DB_SSL_MODE = dbSslMode
	}
	if err1 == nil {
		DB_MAX_CONNECTION = int(maxConnection)
	}
	if err2 == nil {
		DB_LOG_MODE = logMode
	}
}

//GET DB CONNECTION
func GetInstance() *gorm.DB {
	ONCE.Do(func() {
		DBSession = buildConnection()
	})
	return DBSession
}

func buildConnection() *gorm.DB {
	strConnection := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		DB_HOST, DB_PORT, DB_USER, DB_NAME, DB_SSL_MODE, DB_PASSWORD)
	log.Print("DB URL : " + strConnection)
	db, err := gorm.Open(DB_DRIVER, strConnection)
	if err != nil {
		panic(err)
	}

	log.Print("[DATABASE] connection successful")
	//Ativate Logging Mode (SQL)
	db.LogMode(DB_LOG_MODE)
	//Setup Max Connection
	db.DB().SetMaxIdleConns(DB_MAX_CONNECTION)
	db.DB().SetMaxOpenConns(DB_MAX_CONNECTION)
	return db
}

func AutoMigrate() {
	db := GetInstance()
	db.AutoMigrate(&model.Policy{})
	db.AutoMigrate(&model.Person{})
	db.AutoMigrate(&model.Address{})

}
