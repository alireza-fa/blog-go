package models

type User struct {
	BaseModel
	UserName  string `gorm:"type:string;size:64;not null;unique"`
	Email     string `gorm:"type:string;size:64;not null;unique"`
	FullName  string `gorm:"type:string;size:64;null"`
	Password  string `gorm:"type:string;size:120;not null"`
	UserRoles *[]UserRole
}

type Role struct {
	BaseModel
	Name      string `gorm:"type:string;size:10;not null;unique"`
	UserRoles *[]UserRole
}

type UserRole struct {
	BaseModel
	User   User `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	Role   Role `gorm:"foreignKey:RoleId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	UserId int
	RoleId int
}
