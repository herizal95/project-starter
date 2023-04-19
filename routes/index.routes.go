package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/herizal95/project-starter/routes/auth_route"
	"github.com/herizal95/project-starter/routes/user_route"
)

func InitRoutes(app *gin.Engine) {

	// USER ROUTES
	user_route.UserRoutes(app)

	// AUTH ROUTES
	auth_route.AuthRoutes(app)
}
