package model

import (
	"github.com/fqzz2000/tiny-tictok/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init_DB() {
	var err error
	DB, err = gorm.Open(mysql.Open(config.DBConnectString()), &gorm.Config{
			PrepareStmt: true,
			SkipDefaultTransaction: true,
	})
	if (err != nil) {
		panic(err)
	}
}