package envars

type ServerSetting struct {
	PORT        string
	DB_PORT     string
	DB_NAME     string
	DB_DATABASE string
	DB_USER     string
	DB_PASSWORD string
}

type DbSetting struct {
	DB_PORT     string
	DB_NAME     string
	DB_DATABASE string
	DB_USER     string
	DB_PASSWORD string
	CHARSET     string
	COLLATION   string
	PARSETIME   string
}
