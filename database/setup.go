package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetUpDB() (conn *gorm.DB, err error) {
	conn, err = gorm.Open(sqlite.Open("./database/database.sqlite"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	db, err := conn.DB()
	defer func() {
		err = db.Close()
		if err != nil {
			panic(err)
		}
	}()
	if err != nil {
		db.Close()
	}

	return conn, err
}
