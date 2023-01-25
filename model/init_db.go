package model

import (
	"github.com/fqzz2000/tiny-tictok/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init_DB(debug bool) {
	var err error
	if (debug) {
			DB, err = gorm.Open(mysql.Open(config.DebugDBConnectString()), &gorm.Config {
				PrepareStmt: true,
				SkipDefaultTransaction: true,
		})
		if (err != nil) {
			panic(err)
		}
	} else {
		DB, err = gorm.Open(mysql.Open(config.DBConnectString()), &gorm.Config {
				PrepareStmt: true,
				SkipDefaultTransaction: true,
		})
		if (err != nil) {
			panic(err)
		}
	}
}