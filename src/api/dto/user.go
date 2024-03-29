package dto

type Profile struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
}

type ProfileUpdate struct {
	FullName string `json:"fullName" validate:"min=5,max=64"`
}

type CreateUser struct {
	UserName string `json:"userName" validate:"required,max=64,min=5"`
	Email    string `json:"email" validate:"email,required,min=10,max=64"`
	FullName string `json:"fullName" validate:"min=5,max=64"`
	Password string `json:"password" validate:"required,min=8,max=64,containsany=abcdefghijklmnopqrstuvwxyz,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=1234567890"`
}

type UserVerify struct {
	UserName string `json:"userName" validate:"required,max=64,min=5"`
	Code     int    `json:"code" validate:"required,min=1000,max=9999"`
}

type UserLogin struct {
	UserName string `json:"userName" validate:"required,max=64,min=5"`
	Password string `json:"password" validate:"required,min=8,max=64,containsany=abcdefghijklmnopqrstuvwxyz,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=1234567890"`
}
