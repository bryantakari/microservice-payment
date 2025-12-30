package main

import (
	"github.com/bryantakari/microservice-payment/order-service/internal/app"
	"github.com/bryantakari/microservice-payment/order-service/internal/config"
)

func main() {
	conf := config.Load()
	app := app.New(conf)
	app.Listen(":" + conf.AppPort)
}
