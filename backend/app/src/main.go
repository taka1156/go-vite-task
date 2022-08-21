package main

import (
	"app/infra/server"
	"app/injecter"
)

func main() {
	injecter.Injection()
	server.App()
}
