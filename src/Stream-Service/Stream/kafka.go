package Stream

import (
	"github.com/Shopify/sarama"
	"time"
	"fmt"
	"encoding/json"
)

type producer struct {
	AsyncProducer sarama.AsyncProducer
	SyncProducer sarama.SyncProducer
}

func getAsyncProducer(brokerList []string) sarama.AsyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Compression = sarama.CompressionSnappy
	config.Producer.Flush.Frequency = 500 * time.Millisecond

	producer, err := sarama.NewAsyncProducer(brokerList, config)
	if err != nil {
		fmt.Printf("Error creating Async Producer %v", err)
	}

	return producer
}

func getSyncProducer(brokerList []string) sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		fmt.Printf("Error creating sync Producer % v", err)
	}

	return producer
}

func (producer producer) sendSyncMarketEventMessage(producerTopic string, payload marketEvent) {
	
	bytesRep, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("error marshaling paylod %v", err)
	}

	message := &sarama.ProducerMessage{
		Topic:  producerTopic,
		Value: sarama.ByteEncoder(bytesRep),
	}

	//partition, offset
	_, _, err2 := producer.SyncProducer.SendMessage(message)
	if err2 != nil {
		fmt.Printf("Error sending Sync message %v" , err2)
	}
}

func (producer producer) sendAsyncMarketEventMessage(producerTopic string, payload marketEvent) {
	
	bytesRep, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("error marshasling paylod %v", err)
	}

	message := &sarama.ProducerMessage{
		Topic:  producerTopic,
		Value: sarama.ByteEncoder(bytesRep),
	}

	producer.AsyncProducer.Input() <- message
}