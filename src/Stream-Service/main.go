package main

import (
	"Stream-Service/Stream"
)

//symbol must be lower case
//var base_url = "wss://stream.binance.com:9443/ws/ethbtc@depth"


func main() {
	
	binance := Stream.GetBinanceExchange()
	coinbase := Stream.GetCoinbaseExchange()
	//bitmex := Stream.GetBitmexExchange()
	//binance_futures := Stream.GetBinanceFuturesExchange()
	//kraken := Stream.GetKrakenExchange()
	//deribit := Stream.GetDeribitExchange()
	//ftx := Stream.GetFTXExchange()
	//bitfinex := Stream.GetBitfinexExchange()

	binance.CreateStream([]string{"btcusdt"})
	coinbase.CreateStream([]string{"BTC-USD"})
	//bitmex.CreateStream([]string{"XBTUSD"})
	//binance_futures.CreateStream([]string{"btcusdt"})
	//kraken.CreateStream([]string{"XBT/USD"})
	//deribit.CreateStream([]string{"trades.BTC-PERPETUAL.raw"})
	//ftx.CreateStream([]string{"BTC-PERP"})
	//bitfinex.CreateStream([]string{"tBTCUSD"})


	select {}
} 
