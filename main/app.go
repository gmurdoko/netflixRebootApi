package main

import (
	"netflixReboot/config"
	"netflixReboot/main/masters"
)

func main() {
	db := config.EnvConn()
	router := config.CreateRouter()
	masters.Init(router, db)
	config.RunServer(router)
}
