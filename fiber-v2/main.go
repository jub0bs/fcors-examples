package main

import (
	"log"
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/jub0bs/fcors"
)

func main() {
	// create an app
	app := fiber.New()

	// configure CORS middleware
	cors, err := fcors.AllowAccess(
		fcors.FromOrigins("https://example.com"),
		fcors.WithMethods(
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
		),
		fcors.WithRequestHeaders("Authorization"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// apply the CORS middleware to the app
	app.Use(adaptor.HTTPMiddleware(cors))

	// register the hello handler for the /hello route
	app.Get("/hello", hello)

	// start the server on port 8080
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}

func hello(c *fiber.Ctx) error {
	data := struct {
		Hello string `json:"hello"`
	}{
		Hello: "world",
	}
	return c.JSON(data)
}
