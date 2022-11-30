package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func InitPostgresql() *sql.DB {
	username := viper.GetString("POSTGRES_USERNAME")
	password := viper.GetString("POSTGRES_PASSWORD")
	host := viper.GetString("POSTGRES_HOST")
	port := viper.GetString("POSTGRES_PORT")

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/miniwa?sslmode=disable",
		username, password, host, port,
	)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err.Error())
	}

	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(2 * time.Hour)
	db.SetMaxIdleConns(10)
	return db
}
