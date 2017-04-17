package config

import (

	"github.com/eclipse/paho.mqtt.golang"
)


func Init(conn string) *mqtt.ClientOptions{

	return mqtt.NewClientOptions().AddBroker(conn)
}
