package app

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/herizal95/project-starter/config"
	"github.com/herizal95/project-starter/database"
	"github.com/herizal95/project-starter/routes"
	"github.com/joho/godotenv"
)

func Run() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	} else {
		log.Println("Success loading .env file")
	}

	// INIT CONFIG
	config.InitConfig()

	// INIT DATABASE
	database.ConnectDatabase()

	// GIN
	app := gin.Default()

	// INIT ROUTES
	routes.InitRoutes(app)

	// START SERVER
	app.Run(os.Getenv("APP_PORT"))
}
