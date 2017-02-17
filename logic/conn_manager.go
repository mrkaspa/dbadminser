package logic

import "github.com/mrkaspa/dbadminser/store"

type doFunc func() error

// ConnManager handles the conn
type ConnManager interface {
	DoConn(store.Conn, doFunc) error
}
