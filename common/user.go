package common

type User struct {
	UserName  string
	FirstName string
	LastName  string
}

type Login struct {
	UserName string `form:"username" json:"username" binding:"required"`
	PassWord string `form:"password" json:"password" binding:"required"`
	Captcha  string `form:"captcha" json:"captcha" binding:"required"`
}
