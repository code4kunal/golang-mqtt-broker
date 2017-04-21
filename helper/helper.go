package helper

import (
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"math/rand"
	"strconv"
	"time"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang-task/models"
	"os"
	"encoding/json"

)

// Publish message on /acuators/room-1 topic
func PubValveValue (str []byte, client mqtt.Client){

	token := client.Publish("/actuators/room-1", 0, true, str)
	token.Wait()
	//str := strings.Replace(str, \,  " ", -1)
	fmt.Printf("Publishing message on topic: /actuators/room-1 : %v \n", string(str))

}


// Unsubscribe from a topic
func Unsubscribe (broker *models.BrokerSub,  client mqtt.Client) error{

	if token := client.Unsubscribe(broker.Topic); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return token.Error()
	}
	return nil
}



// Publish message on /readings/temperature
func Publish(broker *models.BrokerPub, client mqtt.Client) error {

	if broker.Action != "pub" {
		err := errors.New("Invalid setting for action, must be pub")
		return err
	}

	if broker.Topic == "" {
		err := errors.New("Invalid setting for topic, must not be empty")
		return err
	}

	token := client.Publish(broker.Topic, 0, true, broker.Payload)
        token.Wait()

	return nil

}


// Callback on subscriber after receiving message for a given topic
func onMessageRecv(client mqtt.Client, message mqtt.Message){

	var temp models.TempReading
	if err := json.Unmarshal(message.Payload(), &temp); err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("Received message on topic: %v  Message: %+v \n ", message.Topic(), string(message.Payload()))

	valvePercent := RegulateTemp(temp.Value)
	reqNum := fmt.Sprintf("%.2f", valvePercent)
	out, err := json.Marshal(reqNum)
	if err != nil {
		log.Println(err)
		return
	}

	PubValveValue(out,client )

}

// Specifies valve openness in percentages
func RegulateTemp (value float64) float64{

	if value> 22 {
		return (22 * 100)/ value
	}else if value < 22{
              return 100.0
	}

	return 0.0
}


// Subscribes to a given topic
func Subscribe(broker *models.BrokerSub, client mqtt.Client) error {

	if broker.Topic == "" {
		err := errors.New("Invalid setting for topic, must not be empty")
		log.Errorln(err)
		os.Exit(1)
	}

	if token := client.Subscribe(broker.Topic, 0, onMessageRecv); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			return token.Error()
	}
	return nil
}


// Generates periodic messages and call publisher
func GenPeriodicTemp(d time.Duration, c mqtt.Client) error {

	rand.Seed(time.Now().UTC().UnixNano())
	ticker := time.NewTicker(d)

	go func() {
		for x := range ticker.C{

			reqRandTemp := GetRandomfloat()
			randomInt := GetRandomInt(0, 10)
			randomIntString := fmt.Sprintf("%d", randomInt)
			sensorID := "sensorID" + "-" + randomIntString

			temp := &models.TempReading{
				SensorID: sensorID,
				Type:     "Temperature",
				Value:    reqRandTemp,
			}

			out, err := json.Marshal(temp)
			if err != nil {
				log.Println(err)
				return
			}

			broker := &models.BrokerPub{
				Topic:   "/readings/temperature",
				Action:  "pub",
				Payload: string(out),
				Num:     1,
			}

			err = Publish(broker, c)
			if err != nil {
				err = errors.New("Unable to publish reading")
				return
			}
			fmt.Printf("%+v  published  at %v \n", broker, x)

		}

	}()

	time.Sleep(time.Second * 20)
	ticker.Stop()

	return nil
}

// Generates random integer
func GetRandomInt(min int, max int) int {

	return min + rand.Intn(max-min)
}


// Generates random float
func GetRandomfloat() float64 {

	num := (rand.Float64() * 30) + 7
	reqNum := fmt.Sprintf("%.2f", num)

	i, err := strconv.ParseFloat(reqNum, 64)
	if err != nil {
		log.Println(err)
		return 0.0
	}
	return i
}

