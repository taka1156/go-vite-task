package main

import (
	"app/di"
	"app/infra/server"
)

func main() {
	di.Do()
	server.App()
}
