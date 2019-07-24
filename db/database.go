package db

import (
	"github.com/jinzhu/gorm"

	"github.com/grahamgilbert/mdmdirector/utils"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Open() error {
	var err error
	DB, err = gorm.Open("postgres", utils.DBConnectionString())
	if err != nil {
		return err
	}

	return nil

}

func Close() error {
	return DB.Close()
}
