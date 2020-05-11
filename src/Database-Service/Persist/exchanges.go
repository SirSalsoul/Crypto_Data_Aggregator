package Persist

import (
	rts "github.com/RedisTimeSeries/redistimeseries-go"
	//"database/sql"
)

type Exchange struct {
	name string
	topic string
	redis_server *rts.Client
	postgres_client DBC
}

func NewExchange(name string, topic string, redis_server string, postgres_client DBC) Exchange {
	return Exchange{
		name: name,
		topic: topic,
		redis_server: rts.NewClient(redis_server, "nohelp", nil),
		postgres_client: postgres_client,
	}
} 