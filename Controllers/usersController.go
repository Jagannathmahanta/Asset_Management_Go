package Controllers

import (
	"AssetManagementSystem/Config"
	"AssetManagementSystem/Models"
	"AssetManagementSystem/utils"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser create users
func CreateUser(c *gin.Context) {
	//validate input
	status := true
	var input Models.Users
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	password := input.Password
	hash, _ := HashPassword(password)
	// Creating User
	user := Models.Users{
		ID:           input.ID,
		Username:     input.Username,
		EmployeeId:   input.EmployeeId,
		Email:        input.Email,
		Password:     hash,
		MobileNumber: input.MobileNumber,
		Role:         input.Role,
		Location:     input.Location,
		UserStatus:   &status}

	//Password encryption

	result := Config.DB.Create(&user)

	//validate if unable to store data in db bez of foreign keys or duplicate entry
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "Unable to create User... Please try again"})
		return
	}
	//success validation
	c.JSON(http.StatusCreated,
		gin.H{
			"status":  "201",
			"message": "user created successfully",
		})

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// FindUser find active users
func FindUser(c *gin.Context) {
	var users []Models.Userlist
	result := Config.DB.Model(&Models.Users{}).
		Select("users.id,users.username,users.employee_id,users.email,users.password,users.mobile_number,roles.role,location_details.location,users.user_status").
		Joins("JOIN roles on users.role=roles.id").
		Joins("JOIN location_details on users.location = location_details.id").
		Where("user_status = 1").Scan(&users)

	//validation for not able get data
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "Unable to fetch user details"})
		return
	}
	//c.JSON(http.StatusOK, users)
	c.JSON(http.StatusOK,
		gin.H{
			"status":  "200",
			"message": "user Fetched successfully",
			"Data":    users,
		})
}

// FindUserById API to get users by id
func FindUserById(c *gin.Context) { // Get model if exist
	var user Models.Users

	result := Config.DB.Where("id = ?", c.Param("id")).First(&user)
	//validation for not able get data
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Unable to fetch users by id"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "user Fetched successfully",
		"Data":    user})
}

// Update update user
func Update(c *gin.Context) {
	var user Models.Users

	if err := Config.DB.Where("id=?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}
	var input Models.Users
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//Config.DB.Find(&users)

	result := Config.DB.Model(&user).Updates(Models.Users{
		MobileNumber: input.MobileNumber,
		Role:         input.Role,
		Location:     input.Location,
	})
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "Unable to update... Please try again"})
		return
	}
	c.JSON(200, gin.H{
		"status":  204,
		"message": "User updated successfully",
		"data":    user,
	})
}

// DeleteUsers delete users
func DeleteUsers(c *gin.Context) {
	var user Models.Users

	if err := Config.DB.Where("id=?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	userStatus := false

	result := Config.DB.Model(&user).Updates(Models.Users{
		UserStatus: &userStatus,
	})

	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "Unable to delete User... Please try again"})
		return
	}

	c.JSON(200, gin.H{
		"status":  204,
		"message": "User deleted successfully",
		"data":    user,
	})
}

type MyCustomClaims struct {
	Email string `json:"email"`
	Id    string `json:"id"`
	jwt.RegisteredClaims
}

// ForgotPasswordSendUrl creates a url to reset password and send it to the user email.

func ForgotPasswordSendUrl(c *gin.Context) {

	var payload *Models.ForgotPasswordInput

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
	}

	// message := "You will receive a reset email if user with that email exist"

	var user Models.Users

	result := Config.DB.First(&user, "email = ?", strings.ToLower(payload.Email))

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": "Invalid email "})
	}

	if !*user.UserStatus {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Account not verified"})
		return
	}

	var secret = os.Getenv(os.Getenv("SECRET")) + user.Password
	claim := &MyCustomClaims{Email: user.Email, Id: strconv.Itoa(int(user.ID)), RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15))}}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	token, err := claims.SignedString([]byte(secret))
	if err != nil {
		c.JSON(400,
			gin.H{
				"status": http.StatusNotImplemented,
				"error":  "failed to create token",
			})
		return
	}
	fmt.Println("Token For forget password", token)

	emailData := Models.EmailData{
		URL:       os.Getenv("CLIENT_URL") + "/" + strconv.Itoa(int(user.ID)) + "/" + token,
		FirstName: user.Username,
		Subject:   "Your password reset token (valid for 10min)",
	}

	if err := utils.SendEmail(&user, &emailData, "resetPassword.html"); err != nil {
		c.JSON(http.StatusNotImplemented,
			gin.H{
				"status": http.StatusNotImplemented,
				"error":  "Failed to send Email",
			})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "You will receive a reset email, which is valid for 15 minutes"})

}

// ForgotPasswordNewPassword verify the given token and set new password

func ForgotPasswordNewPassword(c *gin.Context) {

	var payload *Models.ResetPasswordInput

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
	}
	if payload.Password != payload.PasswordConfirm {
		c.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": "Password and Confirm password do not match"})
		return
	}
	paramId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		fmt.Println("Invalid Id")
		c.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": "Invalid user id"})
		return
	}

	var user Models.Users

	result := Config.DB.First(&user, "id = ?", paramId)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": "Invalid user id"})
		return
	}
	var secret = os.Getenv(os.Getenv("SECRET")) + user.Password
	token, err := jwt.ParseWithClaims(c.Param("token"), &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok && !token.Valid {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": "Invalid Token"})
		return
	}
	fmt.Printf("%v %v %v", claims.Email, claims.RegisteredClaims.Issuer, claims.Id)

	if claims.RegisteredClaims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		fmt.Println("Token is expired")
		c.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": "Token is expired"})
		return
	}
	updateResult := Config.DB.Model(&user).Updates(Models.Users{
		Password: payload.Password,
	})
	if updateResult.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "Unable to update... Please try again"})
		return
	}

	c.JSON(200, gin.H{
		"status":  204,
		"message": "Password changed successfully,Kindly Re-login!! ",
	})
}

func ChangePassword(c *gin.Context) {
	var user Models.Users

	var input Models.ResetPassword

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := Config.DB.Where("id=?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found!",
		})
		return
	}

	//Config.DB.Find(&users)

	result := Config.DB.Model(&user).Where("id = ?", c.Param("id")).Update(Models.Users{
		Password: input.Password,
	})
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "Unable to change password... Please try again"})
		return
	}
	c.JSON(200, gin.H{
		"status":  204,
		"message": "Password Changed successfully",
		//"data":    user,
	})
}
