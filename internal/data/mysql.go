package data

import (
	gMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGorm(source string) (db *gorm.DB, err error) {
	db, err = gorm.Open(gMysql.Open(source), &gorm.Config{})
	return db, err
}
