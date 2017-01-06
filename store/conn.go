package store

import (
	"fmt"
	"time"
)

const MYSQL = "mysql"

// Conn to the DB
type Conn struct {
	ID        uint `gorm:"primary_key"`
	Type      string
	Name      string
	User      string
	Pass      string
	Host      string
	DBName    string
	Port      int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// URI Connection
func (c Conn) URI() string {
	switch c.Type {
	case MYSQL:
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.User, c.Pass, c.Host, c.Port, c.DBName)
	default:
		return ""
	}
}
