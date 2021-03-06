package cmd

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rodrigoengelberg/imersao-fullcycle-7/codepix/application/kafka"
	"github.com/rodrigoengelberg/imersao-fullcycle-7/codepix/infrastructure/db"
	"os"

	"github.com/spf13/cobra"
)

var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "Start consuming transaction using Apache Kafka",
	Run: func(cmd *cobra.Command, args []string) {
		database := db.ConnectDB(os.Getenv("env"))
		producer := kafka.NewKafkaProducer()
		deliveryChan := make(chan ckafka.Event)

		go kafka.DeliveryReport(deliveryChan)

		kafkaProcessor := kafka.NewKafkaProcessor(database, producer, deliveryChan)
		kafkaProcessor.Consume()
	},
}

func init() {
	rootCmd.AddCommand(kafkaCmd)
}
