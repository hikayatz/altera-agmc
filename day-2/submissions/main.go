package main

import (
	"log"

	"github.com/hikayat13/alterra-agcm/day-2/submission/config"
	"github.com/hikayat13/alterra-agcm/day-2/submission/routes"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitDB()
}
func main() {

	e := routes.New()
	config.SetupLogger(e)
	e.Logger.Fatal(e.Start(":8000"))
}
