package database

import (
	"gorm.io/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = "P@55w0rd"
const DB_NAME = "azimio"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

var Db *gorm.DB

var (
	// DBCon is the connection handle
	// for the database
	DBCon *gorm.DB
)
