package main

import (
	"log"
	"producer/controller"
	"producer/services"
	"strings"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func init() {
	// Set configuration file path
	viper.SetConfigFile("config.yaml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

}

func main() {
	producer, err := sarama.NewSyncProducer(viper.GetStringSlice("kafka.servers"), nil)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	event := services.NewEventProducer(producer)
	service := services.NewAccountService(event)
	controller := controller.NewAccountController(service)

	app := fiber.New()
	app.Post("/open", controller.OpneAccount)
	app.Post("/deposit", controller.DepositFund)
	app.Post("/withdraw", controller.WithDrawFund)
	app.Post("/close", controller.CloseAccount)

	app.Listen(":8000")
}
