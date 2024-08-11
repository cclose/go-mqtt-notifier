package main

import (
	"github.com/cclose/go-mqtt-notifier/notifier"
	"log"
)

func main() {

	ss, err := notifier.NewSESService()
	if err != nil {
		log.Fatal(err)
	}

	ss.SendSESEmail("pulsar2612@hotmail.com", "Testing SES MEssages", "Hello, Me!")
}

khabavorsk @ 5am
Khabavorsk -> Moscow => 7 hour flight
2 hour layover
MSK -> Paris => 4.5 hour
Layover 6 hours
Paris -> Chicago => 9 hours
