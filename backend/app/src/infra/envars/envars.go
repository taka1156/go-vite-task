package envars

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var ServerEnv = &ServerSetting{
	PORT: os.Getenv("PORT"),
}

var DbEnv = &DbSetting{
	DB_PORT:     os.Getenv("DB_PORT"),
	DB_NAME:     os.Getenv("DB_NAME"),
	DB_DATABASE: os.Getenv("DB_DATABASE"),
	DB_USER:     os.Getenv("DB_USER"),
	DB_PASSWORD: os.Getenv("DB_PASSWORD"),
	CHARSET:     os.Getenv("CHARSET"),
	COLLATION:   os.Getenv("COLLATION"),
	PARSETIME:   os.Getenv("PARSETIME"),
}
