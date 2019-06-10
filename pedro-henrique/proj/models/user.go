package models

type User struct {


	Name     string `json:"name" validate:"required,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=20"`
	


}

type Username struct {

	Name     string `json:"name" validate:"required,max=50"`
	Email    string `json:"email" validate:"required,email,unique-field"`

}

type UserAuth struct {

	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=20"`

}