//go:build !sqlite3

package gormadapter

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var opens = map[string]func(string) gorm.Dialector{
	"mysql":     mysql.Open,
	"postgres":  postgres.Open,
	"sqlserver": sqlserver.Open,
}
