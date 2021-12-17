package entity

type Users struct {
	UserName string `json:"username" binding:"required"`
	PassWord string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Picture  string `json:"picture" binding:"required,url" validate:"is-image"`
}
