package auth_route

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/herizal95/project-starter/controllers/auth_controller"
)

func AuthRoutes(app *gin.Engine) {

	router := app

	routes := router.Group(os.Getenv("BaseUrl") + "/auth")

	// AUTH ROUTES
	routes.POST("/register", auth_controller.Register)
	routes.POST("/login", auth_controller.Login)
}
