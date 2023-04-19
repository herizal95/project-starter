package user_route

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/herizal95/project-starter/controllers/user_controller"
	"github.com/herizal95/project-starter/middleware"
)

func UserRoutes(app *gin.Engine) {

	router := app

	routes := router.Group(os.Getenv("BaseUrl"))

	// USER ROUTE
	userRoute := routes.Group("/users")
	userRoute.Use(middleware.DeserializerMiddleware()) // Protect API use Middleware
	userRoute.GET("/", user_controller.FindUser)
}
