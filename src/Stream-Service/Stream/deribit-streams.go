package Stream

import (
	"fmt"
	//"time"
	"encoding/json"
	"github.com/gorilla/websocket"
)

func send_deribit_subscribe_message(instruments []string) []byte{
	

	subscribeMessage := map[string]interface{}{
		"jsonrpc": "2.0",
		"method": "public/subscribe",
		"id": 42,
		"params": Dictionary{
			"channels": instruments,
		},
	}

	bytesRep, err := json.Marshal(subscribeMessage)
	if err != nil {
		fmt.Printf("json marshal error - %v\n", err)
	}

	return bytesRep
}

func (db Deribit) CreateStream(pairs []string){
	
	c, _, err := websocket.DefaultDialer.Dial(db.Base_url, nil)
	if err != nil {
		fmt.Printf("dialer error %v", err)
	}

	errr := c.WriteMessage(websocket.TextMessage, send_deribit_subscribe_message(pairs))
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

			if val, ok := dmsg_map["id"]; ok && val.(float64) == 42{
				fmt.Printf("SUBSCRIBED TO DERIBIT-PERP \n")
			}
			if val, ok := dmsg_map["method"]; ok && val == "subscription"{
				
				bulk_data := dmsg_map["params"].(map[string]interface{})["data"].([]interface{})

				for i := 0; i < len(bulk_data); i++{
					//fmt.Printf("Market %v at %v executed at %v \n", dmsg_map["side"], dmsg_map["price"], float64(t.UnixNano()))					
					
					data := bulk_data[i].(map[string]interface{})


					order := marketEvent {
						Side: data["direction"].(string),
						Price: data["price"].(float64),
						Time: data["timestamp"].(float64),
						Size: (data["amount"].(float64))/(data["price"].(float64)),
					}
					
					
					db.KafkaProducer.sendSyncMarketEventMessage(db.Topic, order)
				}
			}
		}
	}()
}