package migrations

import (
	"fmt"
	"github.com/alireza-fa/blog-go/src/constants"
	"github.com/alireza-fa/blog-go/src/data/db"
	"github.com/alireza-fa/blog-go/src/data/models"
	"golang.org/x/crypto/bcrypt"
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

	// default data
	createDefaultUserInfo(database)

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

func createDefaultUserInfo(database *gorm.DB) {
	adminRole := models.Role{Name: constants.AdminRole}
	createRoleIfNotExist(database, &adminRole)

	defaultRole := models.Role{Name: constants.DefaultRole}
	createRoleIfNotExist(database, &defaultRole)

	var user models.User
	user.UserName = "blog-go"
	user.Email = "blog-go@hotmain.com"
	user.FullName = "blog go"

	passwordByte, err := bcrypt.GenerateFromPassword([]byte("BlogGo11228"), bcrypt.DefaultCost)
	if err != nil {
		panic("error while creating default user info" + err.Error())
	}
	user.Password = string(passwordByte)

	createUserIfNotExist(database, &user, []int{adminRole.Id, defaultRole.Id})
}

func createRoleIfNotExist(database *gorm.DB, role *models.Role) {
	exists := 0
	database.
		Model(models.Role{}).
		Select("1").
		Where("name = ?", role.Name).
		First(&exists)
	if exists == 0 {
		database.Create(role)
	}
}

func createUserIfNotExist(database *gorm.DB, user *models.User, roleIds []int) {
	exists := 0
	database.
		Model(models.User{}).
		Select("1").
		Where("user_name = ?", user.UserName).
		First(&exists)
	if exists == 0 {
		database := database.Begin()
		database.Create(user)
		for i := 0; i < len(roleIds); i++ {
			userRole := models.UserRole{UserId: user.Id, RoleId: roleIds[i]}
			database.Create(&userRole)
		}
		database.Commit()
	}
}
