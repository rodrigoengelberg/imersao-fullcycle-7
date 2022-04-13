/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rodrigoengelberg/imersao-fullcycle-7/codepix/application/kafka"
	"github.com/rodrigoengelberg/imersao-fullcycle-7/codepix/infrastructure/db"
	"os"

	"github.com/spf13/cobra"
)

// kafkaCmd represents the kafka command
var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "Start consuming transaction using Apache Kafka",
	Run: func(cmd *cobra.Command, args []string) {

		database := db.ConnectDB(os.Getenv("env"))
		producer := kafka.NewKafkaProducer()
		deliveryChan := make(chan ckafka.Event)

		kafka.Publish("Olá Consumer", "teste", producer, deliveryChan)
		go kafka.DeliveryReport(deliveryChan)

		kafkaProcessor := kafka.NewKafkaProcessor(database, producer, deliveryChan)
		kafkaProcessor.Consume()
	},
}

func init() {
	rootCmd.AddCommand(kafkaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kafkaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kafkaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
