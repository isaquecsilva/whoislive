package routes

import (
	"io"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/isaquecsilva/whoislive/src/controller"
)

func StartRoutesAndApp(addr, staticPath string, controller *controller.Controller) error {
	app := fiber.New()

	app.Static("/static/", staticPath, fiber.Static{
		Compress:      true,
		Browse:        false,
		CacheDuration: time.Hour,
	})

	app.Get("/", controller.LiveChannels)

	enableMiddlewares(app)
	enableMetrics(app)

	return app.Listen(addr)
}

func enableMiddlewares(app *fiber.App) {
	// Enabling CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost",
		AllowHeaders: "Origin",
		AllowMethods: "GET, OPTIONS",
	}))

	// Enabling compression
	app.Use(compress.New())

	// Enabling logger
	logs, err := os.Create("logs.log")
	if err != nil {
		panic(err)
	}

	writer := io.MultiWriter(os.Stderr, logs)
	app.Use(logger.New(logger.Config{
		Output:        writer,
		Format:        "[ ${time} | ${ip} ]: ${method} ${path} - ${status}\n",
		DisableColors: false,
	}))
}

func enableMetrics(app *fiber.App) {
	app.Get("/metrics", monitor.New())
}
