package envvars

import (
	"os"

	"github.com/joho/godotenv"
)

var ServerEnv ServerSetting
var DbEnv DbSetting

func EnvLoad() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error)
	}

	ServerEnv = ServerSetting{
		PORT: os.Getenv("PORT"),
	}

	DbEnv = DbSetting{
		DB_PORT:      os.Getenv("DB_PORT"),
		DB_NAME:      os.Getenv("DB_NAME"),
		DB_USER:      os.Getenv("DB_USER"),
		DB_PASSWORD:  os.Getenv("DB_PASSWORD"),
		DB_CHARSET:   os.Getenv("DB_CHARSET"),
		DB_COLLATION: os.Getenv("DB_COLLATION"),
		DB_PARSETIME: os.Getenv("DB_PARSETIME"),
	}
}
