package logic

import "bitbucket.org/hackmin/dbadminser/store"

type doFunc func() error

// ConnManager handles the conn
type ConnManager interface {
	StoreConn(store.Conn) error
	ListConns() []store.Conn
	DeleteConn(store.Conn) error
	DoConn(store.Conn, doFunc) error
}
