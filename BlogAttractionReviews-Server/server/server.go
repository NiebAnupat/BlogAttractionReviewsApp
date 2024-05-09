package server

import (
	"fmt"
	"sync"

	"github.com/NiebAnupat/BlogAttractionReviewsApp/Server/config"
	databases "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type fiberServer struct {
	app  *fiber.App
	db   databases.Database
	conf *config.Config
}

var (
	server *fiberServer
	once   sync.Once
)

func NewFiberServer(conf *config.Config, db databases.Database) *fiberServer {
	fiberApp := fiber.New()
	// fiberApp.Use(logger.New())

	once.Do(func() {
		server = &fiberServer{
			app:  fiberApp,
			db:   db,
			conf: conf,
		}
	})

	return server
}

func (f *fiberServer) Start() {
	f.app.Use(logger.New())
	f.app.Use(recover.New())

	f.app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	f.app.Get("/v1/healthCheck", f.healthCheck)

	f.initAuthRouter()
	f.initBlogPostRouter()

	f.httpListening()
}

func (f *fiberServer) healthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Server is running",
	})
}

func (f *fiberServer) httpListening() {
	url := fmt.Sprintf(":%v", f.conf.Server.Port)
	f.app.Listen(url)
}
