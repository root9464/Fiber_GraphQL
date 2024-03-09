package main

import (
	"log"
	"root/database"
	"root/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func AllRoutes(app *fiber.App) {
	app.Post("/hello", routes.Hello)
	app.Post("/create", routes.Create)
}

func main() {
    _, err := database.ConnectToDB()
    if err != nil {
		log.Fatal("не удалось подключиться к базе данных")
	}
   
    app := fiber.New()
	AllRoutes(app)
	app.Use(cors.New())
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
   

    log.Fatal(app.Listen(":9090"))
}