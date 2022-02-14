/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	appKafka "github.com/c-4u/check-pad/app/kafka"
	"github.com/c-4u/check-pad/app/rest"
	"github.com/c-4u/check-pad/infra/client/kafka"
	"github.com/c-4u/check-pad/infra/client/kafka/topic"
	"github.com/c-4u/check-pad/infra/db"
	"github.com/c-4u/check-pad/utils"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// restCmd represents the rest command
func restCmd() *cobra.Command {
	var servers string
	var groupId string
	var restPort int
	var dsn string
	var dsnType string

	restCmd := &cobra.Command{
		Use:   "rest",
		Short: "Run rest Service",

		Run: func(cmd *cobra.Command, args []string) {
			pg, err := db.NewPostgreSQL(dsnType, dsn)
			if err != nil {
				log.Fatal(err)
			}

			if utils.GetEnv("DB_DEBUG", "false") == "true" {
				pg.Debug(true)
			}

			if utils.GetEnv("DB_MIGRATE", "false") == "true" {
				pg.Migrate()
			}
			defer pg.Db.Close()

			kc, err := kafka.NewKafkaConsumer(servers, groupId, topic.CONSUMER_TOPICS)
			if err != nil {
				log.Fatal("cannot start kafka consumer", err)
			}

			deliveryChan := make(chan ckafka.Event)
			kp, err := kafka.NewKafkaProducer(servers, deliveryChan)
			if err != nil {
				log.Fatal("cannot start kafka producer", err)
			}

			go kp.DeliveryReport()
			go appKafka.StartKafkaServer(pg, kc, kp)
			rest.StartRestServer(pg, kp, restPort)
		},
	}

	dDsn := os.Getenv("DSN")
	sDsnType := os.Getenv("DSN_TYPE")
	dServers := utils.GetEnv("KAFKA_BOOTSTRAP_SERVERS", "kafka:9094")
	dGroupId := utils.GetEnv("KAFKA_CONSUMER_GROUP_ID", "check-pad")

	restCmd.Flags().StringVarP(&dsn, "dsn", "d", dDsn, "dsn")
	restCmd.Flags().StringVarP(&dsnType, "dsnType", "t", sDsnType, "dsn type")
	restCmd.Flags().StringVarP(&servers, "servers", "s", dServers, "kafka servers")
	restCmd.Flags().StringVarP(&groupId, "groupId", "i", dGroupId, "kafka group id")
	restCmd.Flags().IntVarP(&restPort, "port", "p", 8080, "rest server port")

	return restCmd
}

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	if os.Getenv("ENV") == "dev" {
		err := godotenv.Load(basepath + "/../.env")
		if err != nil {
			log.Printf("Error loading .env files")
		}
	}

	rootCmd.AddCommand(restCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// restCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// restCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
