package Stream

//exchange structs used to call CreateStream interface
//contain information to relevant exchanges websocket protocols

import (

)

type Binance struct {
	Base_url string
	KafkaProducer producer
	Topic string
}

type Coinbase struct {
	Base_url string
	KafkaProducer producer
	Topic string
	
}

type Bitmex struct {
	Base_url string
	KafkaProducer producer
	Topic string
}

type BinanceFutures struct {
	Base_url string
	KafkaProducer producer
	Topic string
}

type Kraken struct {
	Base_url string
	KafkaProducer producer
	Topic string
}

type Deribit struct {
	Base_url string
	KafkaProducer producer
	Topic string
}

type Ftx struct {
	Base_url string
	KafkaProducer producer
	Topic string
}

type Bitfinex struct {
	Base_url string
	KafkaProducer producer
	Topic string
}

func GetBinanceExchange() Binance {
	return Binance{
		Base_url: "wss://stream.binance.com:9443/ws/",
		KafkaProducer: producer {
			AsyncProducer: getAsyncProducer([]string{"localhost:9092"}),
			SyncProducer: getSyncProducer([]string{"localhost:9092"}),
		},
		Topic: "binance-spot",
	}
}

func GetCoinbaseExchange() Coinbase {
	return Coinbase{
		Base_url: "wss://ws-feed.pro.coinbase.com",
		KafkaProducer: producer {
			AsyncProducer: getAsyncProducer([]string{"localhost:9092"}),
			SyncProducer: getSyncProducer([]string{"localhost:9092"}),
		},
		Topic: "coinbase-spot",
	}
}

func GetBitmexExchange() Bitmex {
	return Bitmex{
		Base_url: "wss://www.bitmex.com/realtime",
		KafkaProducer: producer {
			AsyncProducer: getAsyncProducer([]string{"localhost:9092"}),
			SyncProducer: getSyncProducer([]string{"localhost:9092"}),
		},
		Topic: "bitmex-perp",
	}
}

func GetBinanceFuturesExchange() BinanceFutures {
	return BinanceFutures{
		Base_url: "wss://fstream.binance.com",
		KafkaProducer: producer {
			AsyncProducer: getAsyncProducer([]string{"localhost:9092"}),
			SyncProducer: getSyncProducer([]string{"localhost:9092"}),
		},
		Topic: "binance-fut",
	}
}

func GetKrakenExchange() Kraken {
	return Kraken{
		Base_url: "wss://ws.kraken.com/",
		KafkaProducer: producer {
			AsyncProducer: getAsyncProducer([]string{"localhost:9092"}),
			SyncProducer: getSyncProducer([]string{"localhost:9092"}),
		},
		Topic: "kraken-spot",
	}
}

func GetDeribitExchange() Deribit {
	return Deribit{
		Base_url: "wss://www.deribit.com/ws/api/v2",
		KafkaProducer: producer {
			AsyncProducer: getAsyncProducer([]string{"localhost:9092"}),
			SyncProducer: getSyncProducer([]string{"localhost:9092"}),
		},
		Topic: "deribit-perp",
	}
}

func GetFTXExchange() Ftx {
	return Ftx{
		Base_url: "wss://ftx.com/ws/",
		KafkaProducer: producer {
			AsyncProducer: getAsyncProducer([]string{"localhost:9092"}),
			SyncProducer: getSyncProducer([]string{"localhost:9092"}),
		},
		Topic: "ftx-perp",
	}
}

func GetBitfinexExchange() Bitfinex {
	return Bitfinex{
		Base_url: "wss://api-pub.bitfinex.com/ws/2",
		KafkaProducer: producer {
			AsyncProducer: getAsyncProducer([]string{"localhost:9092"}),
			SyncProducer: getSyncProducer([]string{"localhost:9092"}),
		},
		Topic: "bitfinex-spot",
	}
}