package Persist

import (
	"fmt"
	"github.com/Shopify/sarama"
	"os"
	//rts "github.com/RedisTimeSeries/redistimeseries-go"
	"encoding/json"
)

type marketEvent struct {
	Side	string
	Price	float64
	Time	float64
	Size	float64
}

func (exchange Exchange) ConsumeTopic(signal chan os.Signal) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	brokers := []string{"localhost:9092"}

	master, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		panic(err)
	}

	consumer, err2 := master.ConsumePartition(exchange.topic, 0, sarama.OffsetOldest)
	if err2 != nil {
		panic(err2)
	}

	fmt.Printf("Write kafka payload to database for %v\n", exchange.name)

	
	go func(consumer sarama.PartitionConsumer) {

		defer consumer.Close()
		defer master.Close()

		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case <-signal:
				fmt.Printf("Interrupt detected")
				return
			case msg := <-consumer.Messages():
				//fmt.Printf("CONSUMING MESSAGES\n")
				//fmt.Printf("%v ::: %v ::: %v\n", exchange.topic, string(msg.Key), string(msg.Value))

				//2 redis keys? one for price and one for quanitity.
				//add exchange label to both?insert_market_event(kafka_payload)
				var kafka_payload marketEvent
				
				err2 := json.Unmarshal(msg.Value, &kafka_payload)	
				if err2 != nil {
					fmt.Printf("unmarshal error: %v", err2)
				}

				exchange.postgres_client.insert_market_event(kafka_payload.Time, kafka_payload.Price, kafka_payload.Size, kafka_payload.Side, exchange.name)

				/* Redis shit
				client := exchange.redis_server
				price_keyname := exchange.name+"_price"
				order_size_keyname := exchange.name+"_order_size"

				price_options := rts.CreateOptions{
					Uncompressed:		true,
					RetentionMSecs:		0,
					Labels:				map[string]string{"type":"price"},
				}

				order_size_options := rts.CreateOptions{
					Uncompressed:		true,
					RetentionMSecs:		0,
					Labels:				map[string]string{"type":"size"},
				}

				
				client.AddWithOptions(price_keyname, int64(kafka_payload.Time), kafka_payload.Price, price_options)
				client.AddWithOptions(order_size_keyname, int64(kafka_payload.Time), kafka_payload.Size, order_size_options)

				*/
			}
		}
	}(consumer)
	
}

//func serializeDataForRedis()