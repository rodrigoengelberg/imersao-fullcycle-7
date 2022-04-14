package cmd

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rodrigoengelberg/imersao-fullcycle-7/codepix/application/grpc"
	"github.com/rodrigoengelberg/imersao-fullcycle-7/codepix/application/kafka"
	"github.com/rodrigoengelberg/imersao-fullcycle-7/codepix/infrastructure/db"
	"github.com/spf13/cobra"
	"os"
)

var (
	gRPCPortNumber int
)

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Run gRPC and a Kafka Consumer",
	Run: func(cmd *cobra.Command, args []string) {
		database := db.ConnectDB(os.Getenv("env"))
		go grpc.StartGrpcService(database, portNumber)

		producer := kafka.NewKafkaProducer()
		deliveryChan := make(chan ckafka.Event)
		go kafka.DeliveryReport(deliveryChan)
		kafkaProcessor := kafka.NewKafkaProcessor(database, producer, deliveryChan)
		kafkaProcessor.Consume()
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
	allCmd.Flags().IntVarP(&gRPCPortNumber, "grpc-port", "p", 50051, "gRPC port")
}
