package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"strings"
)

// AppConfig is for the configuration values.
var (
	AppConfig *viper.Viper
	DB        *gorm.DB
)

func init() {

	AppConfig = viper.New()

	AppConfig.SetConfigFile("../../settings.yaml")
	AppConfig.ReadInConfig()

	setLog()
	//Connect()
}

func Connect() {
	fmt.Println("Welcome to golang rest api with fiber framework")

	godotenv.Load()

	dbhost := os.Getenv("MYSQL_HOST")
	dbuser := os.Getenv("MYSQL_USER")
	dbpassword := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DBNAME")

	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true", dbuser, dbpassword, dbhost, dbname)

	var db, err = gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		panic("db connection failed.")
	}
	DB = db
	fmt.Println("DB connected successfully.")
	autoMigrate(db)
}

func autoMigrate(connection *gorm.DB) {
	err := connection.Debug().AutoMigrate()
	if err != nil {
		panic("DB table  migration is failed due to unexpected error.")
	}
}

func setLog() {

	l := strings.ToUpper(AppConfig.GetString("LOG_LEVEL"))

	switch l {
	case "DEBUG":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "INFO":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "TRACE":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case "WARN":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "ERROR":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "FATAL":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "PANIC":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	fmt.Println("{ Log Level is set to: " + l + "}")

}
