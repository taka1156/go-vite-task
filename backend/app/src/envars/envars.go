package envars

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var Env = &GlobalSetting{
	PORT:          os.Getenv("PORT"),
	DEPLOY_STATUS: os.Getenv("DEPLOY_STATUS"),
}
