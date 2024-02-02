package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/rs/zerolog/log"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Hello, Gopher!!")
	})

	app.Get("/health", monitor.New())

	shutdownCh := make(chan struct{})
	signalCh := make(chan os.Signal, 1)
	signal.Notify(
		signalCh,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	go func() {
		<-signalCh
		log.Info().Msg("Gracefully shutting down...")
		app.Shutdown()
		shutdownCh <- struct{}{}
	}()

	if err := app.Listen(":8080"); err != nil {
		log.Fatal().Msg(err.Error())
	}

	<-shutdownCh
	log.Info().Msg("Running cleanup task...")
}
