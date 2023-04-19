package config

import (
	"github.com/herizal95/project-starter/config/app_config"
	"github.com/herizal95/project-starter/config/db_config"
)

func InitConfig() {
	app_config.InitAppConfig()
	db_config.InitDatabaseConfig()
}
