package Stream

import (
	"github.com/gorilla/websocket"
	"fmt"
	"strings"
	"encoding/json"
)

func (bcf BinanceFutures) CreateStream(pairs []string) {
	format_pairs := strings.Join(pairs, "/")
	url := fmt.Sprintf("%s/ws/%s@trade", bcf.Base_url, strings.ToLower(format_pairs))
	//implement multiple trading paris scenario
	//Combined streams are accessed at /stream?streams=<streamName1>/<streamName2>/<streamName3>
	//Combined stream events are wrapped as follows: {"stream":"<streamName>","data":<rawPayload>}
	//All symbols for streams are lowercase
	/*
	if len(pairs) > 1 {
		url := fmt.Sprintf("%s/stream?streams=%s")
	}
	*/
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		fmt.Printf("dialer error %v", err)
	}
	fmt.Printf("SUBSCRIBED TO BINANCE-FUT \n")

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

			dmsg_map := dmsg.(map[string]interface{})
			//fmt.Printf("BINANCE-FUT :::: %v \n", dmsg_map)

			
			order := marketEvent {
				Side: marketMakerSide(dmsg_map["m"].(bool)),
				Price: stringToFloat(dmsg_map["p"].(string)),
				Time: dmsg_map["T"].(float64),
				Size: stringToFloat(dmsg_map["q"].(string)),
			}
			
			
			//fmt.Printf("Market %v at %v for %v units executed at %v \n", marketMakerSide(dmsg_map["m"].(bool)), dmsg_map["p"].(string), dmsg_map["q"].(string), dmsg_map["T"].(float64))
			bcf.KafkaProducer.sendSyncMarketEventMessage(bcf.Topic, order)
		}
	}()
}