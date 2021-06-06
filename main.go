package main

import (
	"fmb/addr"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/template/html"
)

// variable "port" will be set by "-ldgflags" on makefile
var (
	port string
)

func main() {
	// template engine
	eng := html.New("./addr", ".miho")

	// make new application
	app := fiber.New(fiber.Config{
		Views: eng,
	})

	// designate statics
	app.Static("/", "./misc")

	// designate favicon
	app.Use(favicon.New(
		favicon.Config{
			File: "./misc/favicon.ico",
		},
	))

	// set route
	addr.Routing(app)

	// start server
	app.Listen(port)
}
