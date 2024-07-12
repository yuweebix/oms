package kafka_test

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var (
	topic = "test"
	ch    chan string
)

func getBrokers() (brokers []string, err error) {
	err = godotenv.Load("../../../.env")
	if err != nil {
		return nil, err
	}

	environment := os.Getenv("ENVIRONMENT")
	var brokersString string

	if environment == "docker" {
		brokersString = os.Getenv("TEST_BROKERS")
	} else {
		brokersString = os.Getenv("LOCAL_TEST_BROKERS")
	}

	if brokersString == "" {
		return nil, err
	}

	return strings.Split(brokersString, ","), nil
}
