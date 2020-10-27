package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

// Database ....
type Database struct {
	DB *gorm.DB
}

const (
	host     = "localhost"
	port     = 5432
	user     = "user_ro"
	password = "4r2w3e1q"
	dbname   = "freader"
)

// OpenDatabase ...
func (d *Database) OpenDatabase() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		log.Printf("Соединиться не удалось - %s", err)
	}
	if err2 := db.DB().Ping(); err2 != nil {
		log.Printf("База не отвечает - %v", err2)
	}
	d.DB = db
	return d.DB
}
