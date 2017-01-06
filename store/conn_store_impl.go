package store

import (
	"os/user"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // db
	"github.com/pkg/errors"
)

type connStoreImpl struct{}

type doFunc func(*gorm.DB) error

// NewConnStore returns a new store
func NewConnStore() ConnStore {
	return connStoreImpl{}
}

func (c connStoreImpl) StoreConn(conn Conn) error {
	return doConn(func(db *gorm.DB) error {
		db.Create(&conn)
		if db.NewRecord(conn) {
			return errors.New("Could not store the conn")
		}
		return nil
	})
}

func doConn(do doFunc) error {
	usr, _ := user.Current()
	db, err := gorm.Open("sqlite3", usr.HomeDir+"/dbadmin.db")
	if err != nil {
		return err
	}
	return do(db)
}
