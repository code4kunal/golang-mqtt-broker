package models


type TempReading struct {
	SensorID string
	Type     string
	Value    float64
}

type BrokerSub struct {
	Topic string
}

type BrokerPub struct {
	Topic   string
	Action  string
	Payload string
	Num     int
}
