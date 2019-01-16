package configuration

import (
	"os"

	"github.com/joho/godotenv"
)

// Configuration Env Struct
type Configuration struct {
	Port            string
	APIPrefix       string
	APIVersion      string
	KafkaTopic      string
	KafkaClientID   string
	KafkaBrokerUrls string
}

// Conf Var
var Conf Configuration

// LoadEnv ...
func init() {

	if Conf == (Configuration{}) {
		godotenv.Load()

		// Default configuration
		Conf.Port = os.Getenv("PORT")
		Conf.APIPrefix = os.Getenv("API_PREFIX")
		Conf.APIVersion = os.Getenv("API_VERSION")

		// Kafka configuration
		Conf.KafkaTopic = os.Getenv("KAFKA_TOPIC")
		Conf.KafkaClientID = os.Getenv("KAFKA_CLIENT_ID")
		Conf.KafkaBrokerUrls = os.Getenv("KAFKA_BROKER_URLS")
	}
}
