package app

type User struct {
	Base
	Username string `json:"username" valid:"required~username is required"`
	Email    string `gorm:"unique; not null" json:"email" valid:"required~email is required,email~invalid email address"`
	Password string `json:"password" valid:"required,minstringlength(6)~password must be at least 6 characters"`
}

type LoginRequest struct {
	Email    string `json:"email" valid:"required~email is required,email~invalid email address"`
	Password string `json:"password" valid:"required~password is required"`
}

type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
