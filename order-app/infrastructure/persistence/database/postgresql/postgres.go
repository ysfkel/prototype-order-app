package postgresql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	_db *gorm.DB
)

// GetDB object
func GetDB() (*gorm.DB, error) {
	if _db != nil {
		return _db, nil
	}
	return OpenDB()
}

//OpenDB opens a connection to postgresql
func OpenDB() (*gorm.DB, error) {

	db, err := gorm.Open("postgres", "host=postgres port=5432 sslmode=disable user=yusuf dbname=storedb password=yusuf")

	_db = db

	return db, err
}
