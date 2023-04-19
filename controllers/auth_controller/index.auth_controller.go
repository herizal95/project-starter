package auth_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/herizal95/project-starter/data/request"
	"github.com/herizal95/project-starter/data/response"
	"github.com/herizal95/project-starter/database"
	"github.com/herizal95/project-starter/helper"
	"github.com/herizal95/project-starter/models/entity"
	"github.com/herizal95/project-starter/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// AUTHENTICATION LOGIN
func Login(c *gin.Context) {
	var userInput request.LoginRequest
	if err := c.ShouldBindJSON(&userInput); err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var user entity.User
	if err := database.DB.Where("email = ?", userInput.Email).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.Response(c.Writer, http.StatusNotFound, "Invalid username or password", nil)
			return
		default:
			helper.Response(c.Writer, http.StatusInternalServerError, err.Error(), nil)
			return
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		helper.Response(c.Writer, http.StatusUnauthorized, "Invalid username or password", nil)
		return
	}

	token, err := utils.CreateToken(&user)
	if err != nil {
		helper.Response(c.Writer, http.StatusInternalServerError, "eehh", nil)
		return
	}

	respToken := response.LoginResponse{
		TokenType: "Bearer",
		Token:     token,
	}

	helper.Response(c.Writer, http.StatusOK, "Successfully Login", respToken)
}

// AUTHENTICATION REGISTER
func Register(c *gin.Context) {
	var userInput request.RegisterRequest
	if err := c.ShouldBindJSON(&userInput); err != nil {
		helper.Response(c.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var user entity.User
	if err := database.DB.Where("email = ?", userInput.Email).First(&user).Error; err != gorm.ErrRecordNotFound {
		helper.Response(c.Writer, http.StatusBadRequest, "Email is already registered", nil)
		return
	}

	hashedPassword, _ := utils.GeneratePassword(userInput.Password)

	user = entity.User{
		Name:     userInput.Name,
		Email:    userInput.Email,
		Password: string(hashedPassword),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		helper.Response(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.Response(c.Writer, http.StatusOK, "Successfully registered", nil)
}
