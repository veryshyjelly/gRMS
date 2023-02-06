package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var ChatDb *gorm.DB

func Connect(level logger.LogLevel) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  level,       // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,
		},
	)

	var err error
	ChatDb, err = gorm.Open(sqlite.Open("database/chatdata.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalln("failed to connected to database: ", err)
	}
}