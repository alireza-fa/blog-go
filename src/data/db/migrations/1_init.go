package migrations

import (
	"fmt"
	"github.com/alireza-fa/blog-go/src/data/db"
	"github.com/alireza-fa/blog-go/src/data/models"
	"gorm.io/gorm"
)

func Up1() {
	database := db.GetDb()

	createTables(database)
}

func createTables(database *gorm.DB) {
	tables := []interface{}{}

	// User
	tables = addNewTable(database, models.User{}, tables)
	tables = addNewTable(database, models.Role{}, tables)
	tables = addNewTable(database, models.UserRole{}, tables)

	for i := 0; i < len(tables); i++ {
		err := database.Migrator().CreateTable(tables[i])
		if err != nil {
			panic(fmt.Sprintf("Error while creating table: %T, error: %s", tables[i], err))
		}
	}
	fmt.Println("Table creating...")
}

func addNewTable(database *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}
