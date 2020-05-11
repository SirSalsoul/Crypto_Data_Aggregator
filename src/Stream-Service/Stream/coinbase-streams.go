package Stream

import (
	"fmt"
	"time"
	"encoding/json"
	"github.com/gorilla/websocket"
)

func send_coinbase_subscribe_message(pairs []string, stream_type string) []byte{
	

	subscribeMessage := map[string]interface{}{
		"type": "subscribe",
		"channels": []Dictionary{
			map[string]interface{}{
			"name": stream_type,
			"product_ids": pairs, 
			},
		},
	}

	bytesRep, err := json.Marshal(subscribeMessage)
	if err != nil {
		fmt.Printf("json marshal error - %v\n", err)
	}

	return bytesRep
}

func (cb Coinbase) CreateStream(pairs []string) {

	c, _, err := websocket.DefaultDialer.Dial(cb.Base_url, nil)
	if err != nil {
		fmt.Printf("dialer error %v", err)
	}


	errr := c.WriteMessage(websocket.TextMessage, send_coinbase_subscribe_message(pairs, "ticker"))
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

			dmsg_map := dmsg.(map[string]interface{})

			//fmt.Printf("coinbase trade : %v\n\n", dmsg_map)

			if val, ok := dmsg_map["type"]; ok{
				switch val {
				case "subscriptions":
					fmt.Printf("SUBSCRIBED TO COINBASE-SPOT \n")
				case "ticker":
					t, err3 := time.Parse(time.RFC3339, dmsg_map["time"].(string))
					if err3 != nil {
						fmt.Printf("Something wrong with time conversion %v", err3)
					}

					//fmt.Printf("Market %v at %v executed at %v \n\n", dmsg_map["side"], dmsg_map["price"], float64(t.UnixNano()/1000000))					
					
					
					order := marketEvent {
						Side: dmsg_map["side"].(string),
						Price: stringToFloat(dmsg_map["price"].(string)),
						Time: float64(t.UnixNano()/1000000),
						Size: stringToFloat(dmsg_map["last_size"].(string)),
					}
					
					
					cb.KafkaProducer.sendSyncMarketEventMessage(cb.Topic, order)
				}
			}
		}
	}()
}