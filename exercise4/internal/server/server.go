package server

import (
	"github.com/gofiber/fiber/v2"

	"exercise4/internal/database"
)

type FiberServer struct {
	*fiber.App
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "exercise4",
			AppName:      "exercise4",
		}),
	}
	database.ConnectToDB()

	return server
}
