package main

import (
	"github.com/bryantakari/microservice-payment/order-service/internal/app"
)

func main() {
	app := app.New()
	app.Listen(":8080")
}
