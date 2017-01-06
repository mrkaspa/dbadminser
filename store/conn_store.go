package store

type ConnStore interface {
    StoreConn(Conn) error
}
