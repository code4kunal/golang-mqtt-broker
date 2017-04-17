package main

import (
	"os"
	"time"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/wattx-task/config"
	"github.com/wattx-task/helper"
	"github.com/wattx-task/models"
	"log"
	"flag"

)

func main() {


	brokerConn := flag.String("broker", "tcp://localhost:1883", "The broker URI. ex: tcp://10.10.1.1:1883")
        flag.Parse()

        opts := config.Init(*brokerConn)

	// option for system generated logs for Debug
	//mqtt.DEBUG = log.New(os.Stdout, "", 0)
	//mqtt.ERROR = log.New(os.Stdout, "", 0)


	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)
	c := mqtt.NewClient(opts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
		os.Exit(1)
	}

	log.Println("Connection Established \n")

	brokerSub := &models.BrokerSub{
				Topic: "/readings/temperature",
			}

	err := helper.Subscribe(brokerSub, c)
	if err != nil {
		log.Println("Unable to Subscribe token")
		return
	}


	err = helper.GenPeriodicTemp(10* time.Second, c)
	if err!= nil {
		log.Println(err)
		return
	}

	err = helper.Unsubscribe(brokerSub, c)
	if err!= nil {
		log.Println(err)
		return
	}

	c.Disconnect(250)
	time.Sleep(1 * time.Second)
}