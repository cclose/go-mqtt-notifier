package main

import (
	"github.com/cclose/go-mqtt-notifier/notifier"
	"log"

	MQTT "github.com/eclipse/paho.mqtt.golang"

	"github.com/cclose/go-mqtt-notifier/listener"
)

const (
	mqttBroker   = "tcp://192.168.6.166:1883"
	mqttUsername = "mqttuser"
	mqttPassword = "mqttHomeAuto1806"
	mqttTopic    = "frigate/events"
)

func main() {

	log.Printf("Starting up MQTT Notifier\n")
	log.Printf("Starting up Notifiers\n")
	ns, err := notifier.NewNotificationService()
	if err != nil {
		log.Fatal(err)
	}
	ls := listener.NewListener(ns)

	log.Printf("Starting MQTT Client\n")
	opts := MQTT.NewClientOptions().
		AddBroker(mqttBroker).
		SetUsername(mqttUsername).
		SetPassword(mqttPassword)
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect to MQTT broker: %v", token.Error())
	}
	log.Printf(" - Connected to %s\n", mqttBroker)

	ns.SendNotification(notifier.FrigateEvent, "You've been setup to receive MQTT Events for "+mqttTopic)

	if token := client.Subscribe(mqttTopic, 0, ls.HandleMQTTMessage); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to subscribe to topic: %v", token.Error())
	}
	log.Printf(" - Subscribed to %s\n", mqttTopic)

	// Keep the program running
	select {}
}
