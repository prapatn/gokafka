package main

import (
	"consumer/repositories"
	"consumer/services"
	"context"
	"events"
	"fmt"
	"log"
	"strings"

	"github.com/IBM/sarama"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

func initDB() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)

	dial := mysql.Open(dsn)
	db, err := gorm.Open(dial, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic(err)
	}

	return db
}

func main() {

	consumerGroup, err := sarama.NewConsumerGroup(viper.GetStringSlice("kafka.servers"), viper.GetString("kafka.group"), nil)
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
	defer consumerGroup.Close()

	db := initDB()
	repo := repositories.NewAccountRepository(db)
	eventHandler := services.NewAccountEventHandler(repo)
	consumerHandler := services.NewConsumerHandler(eventHandler)

	topics := events.Topics
	fmt.Println("Consumer Start..")
	for {
		consumerGroup.Consume(context.Background(), topics, consumerHandler)
	}

}

// func main() {
// 	servers := []string{"localhost:9092"}

// 	consumer, err := sarama.NewConsumer(servers, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer consumer.Close()

// 	partitionConsumer, err := consumer.ConsumePartition("king", 0, sarama.OffsetNewest)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer partitionConsumer.Close()
// 	fmt.Println("Consumer Start..")
// 	for {
// 		select {
// 		case err := <-partitionConsumer.Errors():
// 			fmt.Println(err)
// 		case msg := <-partitionConsumer.Messages():
// 			fmt.Println(msg.Topic)
// 			fmt.Println(string(msg.Value))

// 		}
// 	}

// }
