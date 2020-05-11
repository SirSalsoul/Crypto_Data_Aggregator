package Stream

import (
	"github.com/gorilla/websocket"
	"fmt"
	"strings"
	"encoding/json"
)

//symbol must be lower case


func (bc Binance) CreateStream(pairs []string) {
	format_pairs := strings.Join(pairs, ",")
	url := fmt.Sprintf("%s%s@trade", bc.Base_url, strings.ToLower(format_pairs))
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		fmt.Printf("dialer error %v", err)
	}
	fmt.Printf("SUBSCRIBED TO BINANCE-SPOT \n")

	go func(){

		defer c.Close()

		for {
			_, message, err1 := c.ReadMessage()
			if err1 != nil {
				fmt.Printf("read message error: %v", err1)
			}

			var dmsg interface{}
			if err2 := json.Unmarshal(message, &dmsg); err2 != nil {
				fmt.Printf("unmarshal error: %v", err2)
			}
			
			if err != nil {
				fmt.Printf("read error: %v", err)
			}

			//dmsg_map := dmsg.(map[string]interface{})
			
			if dmsg_map, ok := dmsg.(map[string]interface{}); ok {

				order := marketEvent {
					Side: marketMakerSide(dmsg_map["m"].(bool)),
					Price: stringToFloat(dmsg_map["p"].(string)),
					Time: dmsg_map["T"].(float64),
					Size: stringToFloat(dmsg_map["q"].(string)),
				}
			
			
				//fmt.Printf("Market %v at %v for %v units executed at %v \n", marketMakerSide(dmsg_map["m"].(bool)), dmsg_map["p"].(string), dmsg_map["q"].(string), dmsg_map["T"].(float64))
				bc.KafkaProducer.sendSyncMarketEventMessage(bc.Topic, order)
			} else {
				fmt.Printf("Unexpected socket data %v\n", dmsg)
			}
		}
	}()
}