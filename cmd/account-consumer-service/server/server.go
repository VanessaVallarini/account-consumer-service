package server

import (
	"account-consumer-service/cmd/account-consumer-service/middleware"
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/utils"
	"context"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Server *echo.Echo
}

func NewServer() *Server {
	m := middleware.NewMetrics()
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	p := prometheus.NewPrometheus("echo", nil, m.MetricList())
	p.Use(e)
	e.Use(m.AddCustomMetricsMiddleware)

	return &Server{
		Server: e,
	}
}

func (s *Server) Start(c *models.Config) {
	utils.Logger.Info("starting server in port " + c.ServerHost)
	err := s.Server.Start(c.ServerHost)

	if err != nil {
		utils.Logger.Fatal(context.Background(), err, "unable to start server")
		panic(s.Server.Start(c.ServerHost))
	}
}
