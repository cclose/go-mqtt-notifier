package listener

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/cclose/go-mqtt-notifier/notifier"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type Event struct {
	Before EventDetails `json:"before"`
	After  EventDetails `json:"after"`
	Type   string       `json:"type"`
}

type EventDetails struct {
	ID            string   `json:"id"`
	Camera        string   `json:"camera"`
	Label         string   `json:"label"`
	TopScore      float64  `json:"top_score"`
	FalsePositive bool     `json:"false_positive"`
	CurrentZones  []string `json:"current_zones"`
	EnteredZones  []string `json:"entered_zones"`
	StartTime     float64  `json:"start_time"`
	EndTime       *float64 `json:"end_time"`
}

type Listener struct {
	noteService *notifier.NotificationService
}

func NewListener(ns *notifier.NotificationService) *Listener {
	return &Listener{
		noteService: ns,
	}
}

//var MessageHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {

func (l *Listener) HandleMQTTMessage(client MQTT.Client, msg MQTT.Message) {
	//fmt.Printf("Message received: %s\n", msg.Payload())

	var event Event
	if err := json.Unmarshal(msg.Payload(), &event); err != nil {
		log.Printf("Could not unmarshal JSON payload: %v", err)
		return
	}

	// Construct message body with event details
	var message string
	switch event.Type {
	case "new":
		message = fmt.Sprintf("New %s detected by %s @ %s\n", event.After.Label, event.After.Camera, event.After.StartTime)
	case "update":
		message = fmt.Sprintf(" -- %s still detecting by %s @ %s\n", event.After.Label, event.After.Camera, event.After.StartTime)
	case "end":
		message = fmt.Sprintf(" !-- %s has left %s @ %s\n", event.After.Label, event.After.Camera, event.After.StartTime)
	default:
		message = fmt.Sprintf("!!!! Unknown Event type: %s\n", event.Type)
	}

	l.noteService.SendNotification("frigate_event", message)
}
