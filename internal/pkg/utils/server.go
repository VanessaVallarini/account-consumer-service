package utils

import (
	"account-consumer-service/internal/models"
	"context"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type server struct {
	Server *echo.Echo
}

func NewServer() *server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAuthorization, echo.HeaderAccept},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.PUT},
	}))

	return &server{
		Server: e,
	}
}

func (s *server) Start(c *models.Config) {
	Logger.Info("starting server in port " + c.ServerHost)
	err := s.Server.Start(c.ServerHost)

	if err != nil {
		Logger.Fatal(context.Background(), err, "unable to start server")
	}
}
