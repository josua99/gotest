package api

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Server struct {
	api *fiber.App
	db *gorm.DB
}


func NewServer(db *gorm.DB) *Server {
	s := &Server{
		db: db,
	}
	s.setupAPI()

	return s
}

func (s *Server) setupAPI() {
	s.api = fiber.New()
	s.setupRoutes()
}

func (s *Server) Start() error {
	return s.api.Listen(":8080")
}