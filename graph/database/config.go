package database

import (
	"book/graph/model"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func BuildDBConfig() model.DBConfig {
	dbConfig := model.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		DbName:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	return dbConfig
}

func DbURL() string {
	dbConfig := BuildDBConfig()
	return fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DbName)
}

func openDB() {
	var err error
	DB, err = sql.Open("postgres", DbURL())
	if err != nil {
		log.Println("db connect error")
		log.Println(err)
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
}

func GetDB() *sql.DB {
	if DB == nil {
		openDB()
	}

	return DB
}
