package store

import (
	"os/user"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // db
)

type connStoreImpl struct{}

type doFunc func(*gorm.DB) error

// NewConnStore returns a new store
func NewConnStore() ConnStore {
	return connStoreImpl{}
}

func (c connStoreImpl) StoreConn(conn Conn) error {
	return doConn(func(db *gorm.DB) error {
		if db.NewRecord(conn) {
			return db.Create(&conn).Error
		}
		return db.Save(&conn).Error
	})
}

func (c connStoreImpl) DeleteConn(connID int) error {
	return doConn(func(db *gorm.DB) error {
		var conn Conn
		err := db.Where(&Conn{ID: uint(connID)}).First(&conn).Error
		if err != nil {
			return err
		}
		return db.Delete(&conn).Error
	})
}

func (c connStoreImpl) ListConn() []Conn {
	var conns []Conn
	doConn(func(db *gorm.DB) error {
		return db.Find(&conns).Error
	})
	return conns
}

func doConn(do doFunc) error {
	usr, _ := user.Current()
	db, err := gorm.Open("sqlite3", usr.HomeDir+"/dbadmin.db")
	if err != nil {
		return err
	}
	return do(db)
}
