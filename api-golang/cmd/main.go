package main

import (
	"toorme-api-golang/config"
	"toorme-api-golang/internal/app/server"
)

func main() {
	config.LoadEnv()

	server := server.NewServer()
	server.Start()

}
