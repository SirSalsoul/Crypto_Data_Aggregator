package Stream

//Stream data object to store a market event for evey exchange
//Thus every market event produced by kafka is normalized accors exchanges

type marketEvent struct {
	Side	string
	Price	float64
	Time	float64
	Size	float64
}

type Dictionary map[string]interface{}
//opens websocket connect with exchange and produces kafka messages

type Service interface {
	CreateStream(pairs []string)
}