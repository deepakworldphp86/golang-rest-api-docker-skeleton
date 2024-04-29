package config

import (
	"fmt"
	"os"

	"github.com/deepakworldphp86/golang-api/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()

	if errEnv != nil {
		panic("failed to load env")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&entity.Book{}, &entity.User{}, &entity.Receipt{}, &entity.Categories{})

	return db
}

//Postgrees sql

func SetupDatabaseConnectionPgsql() *gorm.DB {

	errEnv := godotenv.Load()

	if errEnv != nil {
		panic("failed to load env")
	}

	db_hostname := os.Getenv("POSTGRES_HOST")
	db_name := os.Getenv("POSTGRES_DB")
	db_user := os.Getenv("POSTGRES_USER")
	db_pass := os.Getenv("POSTGRES_PASSWORD")
	db_port := os.Getenv("POSTGRES_PORT")

	dbURl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", db_user, db_pass, db_hostname, db_port, db_name)
	database, err := gorm.Open(postgres.Open(dbURl), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	database.AutoMigrate(&entity.LoginUser{}, &entity.User{}, &entity.Categories{}, &entity.Customers{}, &entity.Products{}, &entity.Customers{}, &entity.Products{}, &entity.Orders{}, &entity.OrderDetails{})

	return database
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()

	if err != nil {
		panic("failed to close database connection")
	}
	dbSQL.Close()
}
