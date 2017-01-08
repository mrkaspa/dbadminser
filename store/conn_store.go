package store

// ConnStore provides store and query functionalities
type ConnStore interface {
	StoreConn(Conn) error
	DeleteConn(int) error
	ListConn() []Conn
}
