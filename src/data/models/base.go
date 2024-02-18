package models

type Image struct {
	BaseModel
	Name     string `gorm:"type:name;size:64;not null"`
	FileName string `gorm:"type:string;size:128;not null"`
}
