package storage

// import (
// 	"database/sql"

// 	_ "github.com/go-sql-driver/mysql"
// )

// type Storage struct {
// 	db *sql.DB
// }

// func SetupStorage() (*Storage, error) {
// 	db, err := sql.Open("mysql", "local:12345678@/testdb")
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &Storage{db: db}, nil
// }
