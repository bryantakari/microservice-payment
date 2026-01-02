package main

import (
	"os"

	"github.com/bryantakari/microservice-payment/order-service/internal/app"
	"github.com/bryantakari/microservice-payment/order-service/internal/config"
	"github.com/bryantakari/microservice-payment/order-service/internal/logger"
)

func main() {
	log := logger.New()
	log.Info(
		"order-service starting",
	)
	conf := config.Load()
	app := app.New(conf, log)

	if err := app.Listen(":" + conf.AppPort); err != nil {
		log.Error("server stopped", "err", err)
		os.Exit(1)
	}

}
