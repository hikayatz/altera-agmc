package main

import (
	"github.com/hikayat13/alterra-agcm/day-2/submission/middlewares"
	"log"
	"os"

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

	// Auto recovery when panic occurred

	middlewares.Logger(e)
	middlewares.RateLimit(e)
	middlewares.Cors(e)
	config.SetupValidator(e)

	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
