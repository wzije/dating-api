package main

import (
	"fmt"
	"github.com/dating-api/src"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	app := fiber.New()

	//load dotenv
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	//load routers
	src.Routers(app)

	//running on localhost port 3000
	log.Fatal(app.Listen(fmt.Sprintf(":3000")))
}
