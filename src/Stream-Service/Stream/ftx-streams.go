package Stream


import (
	"fmt"
	"encoding/json"
	//"strings"
	"github.com/gorilla/websocket"
	"time"
)

//subscribe message only supports subscription to a single exchange
func send_ftx_subscribe_message(pairs []string, stream_type string) []byte{

	subscribeMessage := map[string]interface{}{
		"op": "subscribe",
		"channel": stream_type,
		"market": pairs[0],
	}

	bytesRep, err := json.Marshal(subscribeMessage)
	if err != nil {
		fmt.Printf("json marshal error - %v\n", err)
	}

	return bytesRep
}

func (ft Ftx) CreateStream(pairs []string) {

	c, _, err := websocket.DefaultDialer.Dial(ft.Base_url, nil)
	if err != nil {
		fmt.Printf("dialer error %v", err)
	}


	errr := c.WriteMessage(websocket.TextMessage, send_ftx_subscribe_message(pairs, "trades"))
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

			//fmt.Printf("FTX MESSAGE :: %v\n", dmsg_map)

			
			if val, ok := dmsg_map["type"]; ok{
				switch val {
				case "subscribed":
					fmt.Printf("SUBSCRIBED TO FTX-PERP \n")
				case "update":	
					//function to aggregate trades? in util
					data := dmsg_map["data"].([]interface{})

					for i := 0; i < len(data); i++{
						single_event := data[i].(map[string]interface{})

						t, err3 := time.Parse(time.RFC3339, single_event["time"].(string))
						if err3 != nil {
							fmt.Printf("Something wrong with time conversion %v", err3)
						}
						
					
						order := marketEvent {
							Side: single_event["side"].(string),
							Price: single_event["price"].(float64),
							Time: float64(t.UnixNano()),
							Size: single_event["size"].(float64),
						}
					
						ft.KafkaProducer.sendSyncMarketEventMessage(ft.Topic, order)
					}
				}
			}
		}
	}()
}

