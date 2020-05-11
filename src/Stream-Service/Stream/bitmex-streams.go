
package Stream


import (
	"fmt"
	"encoding/json"
	"strings"
	"github.com/gorilla/websocket"
	"time"
)


func (bm Bitmex) CreateStream(pairs []string) {
	format_pairs := strings.Join(pairs, ",")
	
	//currency pairs must be uppercase, at least for xbtusd it seems
	url := fmt.Sprintf("%s?subscribe=trade:%s", bm.Base_url, format_pairs)
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		fmt.Printf("dialer error %v", err)
	}
	

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

			dmsg_map := dmsg.(map[string]interface{})

			//fmt.Printf("Bitmex ::: %v\n\n\n", dmsg_map)

			if val, ok := dmsg_map["success"]; ok && val == true{
				fmt.Printf("SUBSCRIBED TO BITMEX-PERP \n")
			}
			if val, ok := dmsg_map["action"]; ok && val == "insert"{ 
				bulk_data := dmsg_map["data"].([]interface{})
	
				for i := 0; i < len(bulk_data); i++{

					data := bulk_data[i].(map[string]interface{})

					t, err3 := time.Parse(time.RFC3339, data["timestamp"].(string))
					if err3 != nil {
						fmt.Printf("Something wrong with time conversion %v", err3)
					}
			
					//fmt.Printf("Timestamp ::: %v\n\n\n", t.UnixNano())
					//TRADE MESSAGE CAN HAVE MULTIPLE MARKET ORDERS COUPLED TOGETHER
			
					order := marketEvent {
						Side: data["side"].(string),
						Price: data["price"].(float64),
						Time: float64(t.UnixNano()/1000000),
						Size: data["size"].(float64),				
					}

					bm.KafkaProducer.sendSyncMarketEventMessage(bm.Topic, order)
				}
			}
		}

	}()
}

