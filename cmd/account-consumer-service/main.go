package main

import (
	"account-consumer-service/cmd/account-consumer-service/listner"
	"account-consumer-service/internal/config"
	"account-consumer-service/internal/pkg/utils"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo"
)

func main() {
	ctx := context.Background()

	config := config.NewConfig()

	go func() {
		setupHttpServer()
	}()

	utils.Logger.Info("start application")

	interrupt := make(chan os.Signal, 1) //se a appp está no ar
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	killsignal := <-interrupt
	switch killsignal {
	case os.Interrupt:
		utils.Logger.Info("got sigint signal... interrupt")
	case syscall.SIGTERM:
		utils.Logger.Info("got sigterm signal... interrupt")
	}

	listner.Start(ctx, config.Kafka)

}

func setupHttpServer() *echo.Echo {
	server := echo.New() //cria um servidor hhtp

	//a := service.NewAccountService()
	//handler.NewAccountHandler(server, a)

	server.Start("0.0.0.0:8080")

	return server
}
