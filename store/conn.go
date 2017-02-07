package store

import (
	"fmt"
	"time"
)

const MYSQL = "mysql"

// Conn to the DB
type Conn struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Type      string `json:"type"`
	Name      string `json:"name"`
	User      string `json:"user"`
	Pass      string `json:"pass"`
	Host      string `json:"host"`
	DBName    string `json:"dbname"`
	Port      int    `json:"port"`
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
