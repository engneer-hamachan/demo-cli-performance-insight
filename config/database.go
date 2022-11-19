package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

var (
	db  *gorm.DB
	err error
)

type StoreData struct {
	Label     string
	Data      float64
	Color     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Connect() *gorm.DB {

	db, err := gorm.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&StoreData{})
	return db

}
