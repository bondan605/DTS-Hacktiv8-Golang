package database

import (
	"database/sql"
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

const (
	HOST     = "localhost"
	PORT     = 5432
	USERNAME = "postgres"
	PASSWORD = "bondan2000"
	DBNAME   = "hacktiv-books"
)

var (
	db  *sql.DB
	err error
)

func GetConnection() *sql.DB {
	if db == nil {
		lock.Lock()
		defer lock.Unlock()
		if db == nil {
			psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
				HOST, PORT, USERNAME, PASSWORD, DBNAME)

			db, err = sql.Open("postgres", psqlInfo)
			if err != nil {
				panic(err)
			}

			err = db.Ping()
			if err != nil {
				panic(err)
			}

			fmt.Println("connected to database")
		} else {
			fmt.Println("Database already connected")
		}
	} else {
		fmt.Println("Database already connected")
	}

	return db
}
