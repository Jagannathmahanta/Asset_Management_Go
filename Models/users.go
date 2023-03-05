package Models

//"github.com/jinzhu/gorm"
// "fmt"

type Users struct {
	ID           uint   `json:"id"`
	Username     string `json:"username"`
	EmployeeId   uint   `json:"employeeId"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	MobileNumber int64  `json:"MobileNumber"`
	Role         uint   `json:"role"`
	Location     uint   `json:"location"`
	UserStatus   *bool  `json:"userStatus"`
}

type Userlist struct {
	Id           int    `json:"id"`
	Username     string `json:"username"`
	EmployeeId   int    `json:"employeeId"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	MobileNumber int    `json:"mobileNumber"`
	Role         string `json:"role"`
	Location     string `json:"location"`
	UserStatus   *bool  `json:"userStatus"`
}

type UserRole struct {
	Id   int    `json:"id"`
	Role string `json:"role"`
}

type SignUpInput struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
	Photo           string `json:"photo" binding:"required"`
}

type SignInInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type ForgotPasswordInput struct {
	Email string `json:"email" binding:"required"`
}

// ? ResetPasswordInput struct
type ResetPasswordInput struct {
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}
type EmailData struct {
	URL       string
	FirstName string
	Subject   string
}

type ResetPassword struct {
	Password        string `json:"password"`
}
