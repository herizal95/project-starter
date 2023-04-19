package user_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/herizal95/project-starter/database"
	"github.com/herizal95/project-starter/helper"
	"github.com/herizal95/project-starter/models/entity"
)

func FindUser(ctx *gin.Context) {

	user := new([]entity.User)

	if err := database.DB.Table("users").Find(&user).Error; err != nil {
		helper.Response(ctx.Writer, 404, "User is Not Found!", nil)
	}

	helper.Response(ctx.Writer, 200, "Successfully fetch all user data!", user)
}
