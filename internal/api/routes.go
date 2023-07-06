package api

// https://docs.gofiber.io/guide/routing
// https://gorm.io/docs
import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func (s *Server)setupRoutes() {
	api := s.api.Group("/api")

	setGlobalMiddleware(api)

	s.setRoutesV1(api)
}

func setGlobalMiddleware(app fiber.Router) {
	app.Use(recover.New())
}

func (s *Server) setRoutesV1(app fiber.Router) {
	v1 := app.Group("/v1")

	v1.Get("/hello", s.helloHandler)
}

func (s *Server) helloHandler(c *fiber.Ctx) error {
	return c.SendString("hello2")
}