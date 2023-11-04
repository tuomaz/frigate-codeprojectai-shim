package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {

	var mqttURI string
	value, ok := os.LookupEnv("MQTTURI")
	if ok {
		mqttURI = value
	} else {
		log.Fatalf("no MQTT URI found")
	}

	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://" + mqttURI)
	opts.SetClientID("frigate-codeprojectai-shim")

	client := MQTT.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	defer client.Disconnect(250)

	topic := "frigate/events" // Använd det ämne du vill prenumerera på
	if token := client.Subscribe(topic, 0, messageHandler); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	fmt.Printf("Prenumererar på ämnet: %s\n", topic)

	// Vänta på avslutssignal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	fmt.Println("Avslutar prenumerationen och stänger anslutningen...")
}

func messageHandler(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Meddelande mottaget på ämnet %s: %s\n", msg.Topic(), msg.Payload())
}
