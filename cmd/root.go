package cmd

import (
	"github.com/omerkaya1/notification/internal/config"
	"github.com/omerkaya1/notification/internal/errors"
	"github.com/omerkaya1/notification/internal/mq"
	"github.com/spf13/cobra"
	"log"
)

var configFile string

// RootCmd is the main entry point to the programme
var RootCmd = &cobra.Command{
	Use:     "notification",
	Short:   "simple notification service that queries RabbitMQ for messages that belong to a particular message queue",
	Example: "# Initialise from configuration file notification -c /path/to/config.json",
	Run:     startNotificationService,
}

func init() {
	RootCmd.Flags().StringVarP(&configFile, "config", "c", "", "specifies the path to a configuration file")
}

func startNotificationService(cmd *cobra.Command, args []string) {
	if configFile == "" {
		log.Fatalf("%s: %s", errors.ErrCMDPrefix, errors.ErrBadConfigFile)
	}

	conf, err := config.InitConfig(configFile)
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrCMDPrefix, err)
	}

	messageQueue, err := mq.NewMessageQueue(conf)
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrMQPrefix, err)
	}
	log.Println("Notification service initialisation")
	if err := messageQueue.EmmitMessages(); err != nil {
		log.Fatalf("%s: %s", errors.ErrMQPrefix, err)
	}
}
