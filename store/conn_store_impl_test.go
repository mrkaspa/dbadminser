package store

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "os"
    "github.com/jinzhu/gorm"
)

func TestMain(m *testing.M) {
    doConn(func(db *gorm.DB) error {
        db.Delete(Conn{})
        return nil
    })
    retCode := m.Run()
    os.Exit(retCode)
}

func TestConnStoreImpl_StoreConn(t *testing.T) {
    conn := Conn{}
    connStore := NewConnStore()
    err := connStore.StoreConn(conn)
    assert.Nil(t, err)
}
