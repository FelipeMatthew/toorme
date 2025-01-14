package main

import (
	"toorme-api-golang/config"
	"toorme-api-golang/internal/server"
)

func main() {
	config.LoadEnv()

	config.ConnectDb()

	server := server.NewServer()
	server.Start()

}
