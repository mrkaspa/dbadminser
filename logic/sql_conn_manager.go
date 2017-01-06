package logic

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // needed
	"github.com/mrkaspa/dbadminser/store"
)

// SQLConnManager
type sqlConnManager struct {
}

func NewSqlConManager() ConnManager {
	return sqlConnManager{}
}

func (s sqlConnManager) StoreConn(conn store.Conn) error {
	return nil
}

func (s sqlConnManager) ListConns() []store.Conn {
	return nil
}

func (s sqlConnManager) DeleteConn(conn store.Conn) error {
	return nil
}

func (s sqlConnManager) DoConn(conn store.Conn, do doFunc) error {
	db, err := sql.Open(string(conn.Type),
		conn.URI())
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()
	do()
	return nil
}
