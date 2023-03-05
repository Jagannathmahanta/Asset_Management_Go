package Controllers

import (
	"AssetManagementSystem/Config"
	"AssetManagementSystem/Models"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	//"golang.org/x/crypto/bcrypt"
)

type body struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// SignIn login api

func SignIn(c *gin.Context) {

	var user Models.Users

	c.BindJSON(&user)

	email := user.Email
	password := user.Password
	var role Models.UserRole

	userData := Models.Users{}
	userQueryColumn := "email=?"
	//search for email in db that matches with given and returns 1st matched data
	if getUserData := Config.DB.Select("password").Where(userQueryColumn,
		strings.ToLower(email)).First(&userData); getUserData.Error != nil {

		c.JSON(400,
			gin.H{
				"status":  "400",
				"message": "Email is incorrect",
			})
		fmt.Println(getUserData)
		return
	}
	//if incorrect password then that checks
	userQueryColumn2 := "password=?"
	if getUserData := Config.DB.Select("email").Where(userQueryColumn2,
		strings.ToLower(password)).First(&userData); getUserData.Error != nil {

		c.JSON(400,
			gin.H{
				"status":  "400",
				"message": "Password is incorrect",
			})
		fmt.Println(getUserData)
		return
	}
	fmt.Println(&userData)

	Config.DB.Select("users.id,roles.role").
		Table("users").
		Joins("JOIN roles on users.role= roles.id").
		Where("email=?", userData.Email).First(&role)

	//generation validation
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Email,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	fmt.Println("Here comes the secret", os.Getenv("SECRET"))

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(400,
			gin.H{
				"status": "400",
				"error":  "failed to create token",
			})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"Message": "logged in successfully",
		"User_id": role.Id,
		"Role":    role.Role,
		"token":   tokenString,
	})

}

func Logout(c *gin.Context) {
	//logout
}
