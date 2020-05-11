package main

import (
	"Database-Service/Persist"
	"os"
	"os/signal"
	"syscall"
	//"github.com/golang-collections/go-datastructures/queue"
	//"fmt"
)

//symbol must be lower case
//var base_url = "wss://stream.binance.com:9443/ws/ethbtc@depth"

// struct api service with universal shit?

//store exchange info in json flat file
//function to initialize all exchange objects neatly

func main() {


	kill_streams := make(chan os.Signal, 1)
	signal.Notify(kill_streams, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(kill_streams, os.Interrupt)

	postgres_client := Persist.Initialize_db_connection()

	binance_spot := Persist.NewExchange("binance_spot", "binance-spot", "localhost:6379", postgres_client)
	coinbase_spot := Persist.NewExchange("coinbase_spot", "coinbase-spot", "localhost:6379", postgres_client)
	//bitmex_perp := Persist.NewExchange("bitmex_perp", "bitmex-perp", "localhost:6379")
	//bitfinex_spot := Persist.NewExchange("bitfinex_spot", "bitfinex-spot", "localhost:6379")
	
	binance_spot.ConsumeTopic(kill_streams)
	coinbase_spot.ConsumeTopic(kill_streams)
	//bitmex_perp.ConsumeTopic(kill_streams)
	//bitfinex_spot.ConsumeTopic(kill_streams)


	<-kill_streams
}