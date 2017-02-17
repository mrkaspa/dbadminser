package logic

import (
	"database/sql"
	"log"

	"bitbucket.org/hackmin/dbadminser/store"
	_ "github.com/go-sql-driver/mysql" // needed
)

// SQLConnManager
type sqlConnManager struct {
}

// NewSQLConManager for handling SQL
func NewSQLConManager() ConnManager {
	return sqlConnManager{}
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
