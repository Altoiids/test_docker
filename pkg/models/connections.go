package models

import (
	"mvc/pkg/types"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

func dsn() string {
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var databaseInfo types.DBInfo

	if err := yaml.Unmarshal(file, &databaseInfo); err != nil {
		log.Fatal(err)
	}

	username := databaseInfo.DB_USERNAME
	password := databaseInfo.DB_PASSWORD
	hostname := databaseInfo.DB_HOST
	dbName := databaseInfo.DB_NAME

	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func Connection() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn())
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(200)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, err
}