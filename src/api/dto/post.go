package dto

type CategoryCreate struct {
	Name string `json:"name" validate:"required,min=3,max=64"`
}

type CategoryUpdate struct {
	Name string `json:"name" validate:"required,min=3,max=64"`
}

type CategoryOutput struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
