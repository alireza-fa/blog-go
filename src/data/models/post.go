package models

type Post struct {
	BaseModel
	User           User `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	UserId         int
	Title          string `gorm:"type:string;size:64;not null"`
	Description    string `gorm:"type:string;size:224;not null"`
	Body           string `gorm:"type:string;size:2048;not null"`
	Image          Image  `gorm:"foreignKey:ImageId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	ImageId        int
	IsPublish      bool `gorm:"type:boolean;default:true;not null"`
	IsActive       bool `gorm:"type:boolean;default:false;not null"`
	TimeToRead     int  `gorm:"type:integer;not null"`
	PostCategories *[]PostCategory
	PostComments   *[]PostComment
}

type Category struct {
	BaseModel
	Name           string `gorm:"type:string;size:64;not null;unique"`
	PostCategories *[]PostCategory
}

type PostCategory struct {
	BaseModel
	Post       Post `gorm:"foreignKey:PostId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	PostId     int
	Category   Category `gorm:"foreignKey:CategoryId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	CategoryId int
}

type PostComment struct {
	BaseModel
	Post    Post `gorm:"foreignKey:PostId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	PostId  int
	User    User `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	UserId  int
	Message string `gorm:"type:string;size:224;not null"`
}
