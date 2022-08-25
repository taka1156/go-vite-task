package envvars

type ServerSetting struct {
	PORT string
}

type DbSetting struct {
	DB_PORT      string
	DB_NAME      string
	DB_USER      string
	DB_PASSWORD  string
	DB_CHARSET   string
	DB_COLLATION string
	DB_PARSETIME string
}
