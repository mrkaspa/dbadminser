package store

import (
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	retCode := m.Run()
	os.Exit(retCode)
}

func clean() {
	doConn(func(db *gorm.DB) error {
		db.Delete(Conn{})
		return nil
	})
}

func TestConnStoreImpl_StoreConn(t *testing.T) {
	clean()
	conn := Conn{}
	connStore := NewConnStore()
	err := connStore.StoreConn(conn)
	assert.Nil(t, err)
}

func TestConnStoreImpl_ListConn(t *testing.T) {
	clean()
	conn := Conn{}
	connStore := NewConnStore()
	connStore.StoreConn(conn)
	res := connStore.ListConn()
	assert.Len(t, res, 1)
}

func TestConnStoreImpl_DeleteConn(t *testing.T) {
	clean()
	conn := Conn{}
	connStore := NewConnStore()
	connStore.StoreConn(conn)
	connStore.DeleteConn(int(conn.ID))
	res := connStore.ListConn()
	assert.Len(t, res, 0)
}
