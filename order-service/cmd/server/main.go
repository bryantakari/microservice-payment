package main

import (
	"os"

	"github.com/bryantakari/microservice-payment/order-service/internal/app"
	"github.com/bryantakari/microservice-payment/order-service/internal/config"
	"github.com/bryantakari/microservice-payment/order-service/internal/db"
	"github.com/bryantakari/microservice-payment/order-service/internal/infra/kafka"
	"github.com/bryantakari/microservice-payment/order-service/internal/logger"
	"github.com/bryantakari/microservice-payment/order-service/internal/order"
)

func main() {
	log := logger.New()
	log.Info(
		"order-service starting",
	)
	conf := config.Load()
	kafka.BootstrapKafka(&conf.KafkaConf)

	db, err := db.NewDbConnection(conf)

	if err != nil {
		log.Error("failed to connect to database", "err", err)
		os.Exit(1)
	}
	stats := db.Stats()
	log.Info("Connection Created ", "stats", stats)

	eventPublisher := kafka.NewOrderPublisher(conf.KafkaConf)

	orderRepository := order.NewRepository(db, log)
	orderService := order.NewService(orderRepository, log, eventPublisher)

	app := app.New(orderService, log)

	if err := app.Listen(":" + conf.AppPort); err != nil {
		log.Error("server stopped", "err", err)
		os.Exit(1)
	}

}
