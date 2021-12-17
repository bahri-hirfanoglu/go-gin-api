package entity

type Users struct {
	UserName string `json:"username" binding:"required"`
	PassWord string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Picutre  string `json:"picture" binding:"required,url"`
}
