package main

import (
	"gzfs/Go~Edita/config"
	"gzfs/Go~Edita/internal/database"
	"gzfs/Go~Edita/internal/users"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	envConfig config.Env
}

func main() {
	envConfig, err := config.Load()
	if err != nil {
		panic(err)
	}

	httpServ := Server{
		envConfig: envConfig,
	}

	cleanupFunc, err := httpServ.Start()
	if err != nil {
		panic(err)
	}

	defer cleanupFunc()

}

func (httpServ *Server) Start() (func(), error) {
	fiberApp, cleanupFunc, err := httpServ.buildServer()
	if err != nil {
		return func() {}, err
	}

	func() {
		fiberApp.Listen("0.0.0.0" + httpServ.envConfig.PORT)
	}()

	return func() {
		cleanupFunc()
		fiberApp.ShutdownWithTimeout(5 * time.Second)
	}, nil
}

func (httpServ *Server) buildServer() (*fiber.App, func(), error) {
	fiberApp := fiber.New()
	fiberApp.Use(logger.New())
	fiberApp.Use(cors.New())

	tursoStorage, err := database.New(httpServ.envConfig)

	if err != nil {
		return nil, func() {}, err
	}

	fiberApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	users.Init(fiberApp, users.NewUserStorage(tursoStorage))

	return fiberApp, func() {
		tursoStorage.TursoDB.Close()
	}, nil
}
