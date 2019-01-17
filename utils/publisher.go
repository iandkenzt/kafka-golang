package utils

import (
	"context"
	"errors"
	"fmt"

	c "github.com/iandkenzt/kafka-golang/configuration"
	kafka "github.com/segmentio/kafka-go"
)

// PublishPanicRecover ...
func PublishPanicRecover() (status bool, err error) {

	if r := recover(); r != nil {
		recoverMsg := fmt.Sprintf("Error publish %v", r)
		errReturn := errors.New(recoverMsg)
		return false, errReturn
	}

	return true, nil
}

// Publish ...
func Publish(key []byte, value []byte) (status bool, err error) {

	defer PublishPanicRecover()

	parent := context.Background()
	defer parent.Done()

	pMessage := kafka.Message{
		Key:   key,
		Value: value,
	}

	w := kafka.NewWriter(c.ProducerConfig)
	w.WriteMessages(parent, pMessage)
	w.Close()

	return true, nil

}
