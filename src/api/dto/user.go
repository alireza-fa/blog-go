package dto

type CreateUser struct {
	UserName string `json:"userName" validate:"required,max=64,min=5"`
	Email    string `json:"email" validate:"email,required,min=10,max=64"`
	FullName string `json:"fullName" validate:"min=5,max=64"`
	Password string `json:"password" validate:"required,min=8,max=64"`
}
