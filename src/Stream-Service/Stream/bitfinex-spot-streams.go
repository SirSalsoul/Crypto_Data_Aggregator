package Stream

import (
	"fmt"
	//"time"
	"math"
	"encoding/json"
	"github.com/gorilla/websocket"
)

func send_bitfinex_subscribe_message(pairs []string, stream_type string) []byte{
	

	subscribeMessage := map[string]interface{}{
		"event": "subscribe",
		"channel": stream_type,
		"symbol": pairs[0], 
	}

	bytesRep, err := json.Marshal(subscribeMessage)
	if err != nil {
		fmt.Printf("json marshal error - %v\n", err)
	}

	return bytesRep
}

func (bf Bitfinex) CreateStream(pairs []string) {

	c, _, err := websocket.DefaultDialer.Dial(bf.Base_url, nil)
	if err != nil {
		fmt.Printf("dialer error %v", err)
	}


	errr := c.WriteMessage(websocket.TextMessage, send_bitfinex_subscribe_message(pairs, "trades"))
	if errr != nil {
		fmt.Printf("subscribe message error: %v", errr)
	}

	go func(){
		
		defer c.Close()

		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				fmt.Printf("read error: %v", err)
			}
			var dmsg interface{}
			err2 := json.Unmarshal(message, &dmsg)
			if err2 != nil {
				fmt.Printf("unmarshal error: %v", err2)
			}

			if dmsg_map, ok := dmsg.(map[string]interface{}); ok {
				if dmsg_map["event"] == "subscribed" {
					fmt.Printf("SUBSCRIBED TO BITFINEX\n")
				}
			}

			if dmsg_map, ok := dmsg.([]interface{}); ok {
				//fmt.Printf("CHANNEL ID :: %v\n", dmsg_map[0])
				//fmt.Printf("DATA :: %v\n", dmsg_map[1])
				if trade_type, ok := dmsg_map[1].(string); ok && trade_type == "te"{

					data := dmsg_map[2].([]interface{})

					//fmt.Printf("BITFINEX MATCH ::: %v\n", data)
					//data[1] = millisecond timestamp
					//data[2] = price
					//data[3] = amount, buy/sell == positive/negative

					
					order := marketEvent {
						Side: marketMaker(data[2].(float64)),
						Price: data[3].(float64),
						Time: float64(data[1].(float64)),
						Size: math.Abs(data[2].(float64)),
					}
				
				
					bf.KafkaProducer.sendSyncMarketEventMessage(bf.Topic, order)
				}
			}
		}
	}()
}

func marketMaker(isBuyer float64) string {
	if (isBuyer <= 0){
		return "sell"
	} else {
		return "buy"
	}
}