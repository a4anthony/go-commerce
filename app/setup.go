package app

import (
	"fmt"
	"os"

	"github.com/a4anthony/go-commerce/config"
	"github.com/a4anthony/go-commerce/database"
	"github.com/a4anthony/go-commerce/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type info struct {
	Name string
}

func SetupAndRunApp() error {
	// load env
	err := config.LoadENV()
	if err != nil {
		return err
	} else {
		fmt.Println("ENV loaded successfully")
	}

	// start database
	database.ConnectDb()

	// create app
	app := fiber.New()

	// attach middleware
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))

	// setup routes
	router.SetupRoutes(app)

	// attach swagger
	config.AddSwaggerRoutes(app)

	// d := info{"Joy"}
	// d.sendMail()

	// get the port and start
	port := os.Getenv("PORT")
	app.Listen(":" + port)

	return nil
}

// func (i info) sendMail() {
// 	result, err := mailer.GetTemplate("template.html", i)
// 	if err != nil {
// 		panic(err)
// 	}
// 	mailer.SendMail("alex@example.com", "bob@example.com", "Welcome", result, "")

// }
