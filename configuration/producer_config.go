package configuration

import (
	"strings"
	"time"

	kafka "github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
)

// ProducerConfig ...
var ProducerConfig kafka.WriterConfig

// InitProducerConfig ...
func InitProducerConfig() {

	// prepare kafka env configurations
	kafkaBrokerUrls := strings.Split(Conf.KafkaBrokerUrls, ",")
	kafkaTopic := Conf.KafkaTopic
	kafkaClientID := "client-id-01"
	if Conf.KafkaClientID != "" {
		kafkaClientID = Conf.KafkaClientID
	}

	// Write Dialer for open kafka connections instead of raw network connections.
	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
		ClientID: kafkaClientID,
	}

	// write kafka configuration
	ProducerConfig = kafka.WriterConfig{
		Brokers:          kafkaBrokerUrls,
		Topic:            kafkaTopic,
		Balancer:         &kafka.LeastBytes{},
		Dialer:           dialer,
		WriteTimeout:     10 * time.Second,
		ReadTimeout:      10 * time.Second,
		CompressionCodec: snappy.NewCompressionCodec(),
	}
}
