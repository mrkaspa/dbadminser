package store

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func init() {
	fmt.Println("Init the DB")

	doConn(func(db *gorm.DB) error {
		db.AutoMigrate(&Conn{})
		return nil
	})
}
