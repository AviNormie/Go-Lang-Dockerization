package Application

import (
	"database/sql"
	"fmt"
	// "os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Create new connection
func makeConnection() *gorm.DB {

	dsn := "username:password@tcp(localhost:3306)/dbname"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print("Error in connecting to database")
	}
	return db
}

// Return connection
func returnConnetcion(db *gorm.DB) *sql.DB {
	connection, err := db.DB()
	if err != nil {
		fmt.Print("Error in connecting to database")
	}
	return connection
}

// connect to database
func connectToDatabase(share ShareResources) {
	switch share.(type) {
	case *Bootstrap:
		application := share.(*Bootstrap)
		application.DB = makeConnection()
		application.Connection = returnConnetcion(application.DB)
	case *Request:
		req := share.(*Request)
		req.DB = makeConnection()
		req.Connection = returnConnetcion(req.DB)
	}
}

// Close Database connection
func CloseConnection(share ShareResources) {
	switch share.(type) {
	case *Bootstrap:
		application := share.(*Bootstrap)
		application.Connection.Close()
	case *Request:
		req := share.(*Request)
		req.Connection.Close()
	}
}
