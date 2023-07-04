package config

import (
	_ "github.com/MatiSierraDev/bookstore-add-db-iteracion-0/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var DATABASE_URI = "root:rootroot@tcp(localhost:3306)/bookstore_test?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() (*gorm.DB, error) {
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// var err error

	// Database, err = gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
	// 	SkipDefaultTransaction: true,
	// 	PrepareStmt:            true,
	// })

	// if err != nil {
	// 	panic(err)
	// }

	// return Database, err
	return gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
}
