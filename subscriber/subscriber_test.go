package main

import (
	DefaultClient "mqtt/src/common"
	"testing"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func TestSubscriber(t *testing.T) {
	t.Run("Subscription to topic", func(t *testing.T) {
		client := DefaultClient.CreateClient(DefaultClient.Broker, DefaultClient.IdSubscriber, DefaultClient.Handler)

		defer client.Disconnect(250)

		if token := client.Connect(); token.Wait() && token.Error() != nil {
			t.Error(token.Error())
		}

		if token := client.Subscribe("sensors/", 1, nil); token.Wait() && token.Error() != nil {
			t.Error(token.Error())
			return
		}

		t.Log("Subscribed successfully to Topic")


	})

	t.Run("Check Payload Integrity", func(t *testing.T) {
		
		publisher := DefaultClient.CreateClient(DefaultClient.Broker, DefaultClient.IdPublisher, DefaultClient.Handler)

		if token := publisher.Connect(); token.Wait() && token.Error() != nil {
			t.Fatal(token.Error())
		}

		defer publisher.Disconnect(250)

		subscriber := DefaultClient.CreateClient(DefaultClient.Broker, DefaultClient.IdSubscriber, DefaultClient.Handler)
		
		if token := subscriber.Connect(); token.Wait() && token.Error() != nil {
			t.Fatal(token.Error())
		}

		defer subscriber.Disconnect(250)

		topic := "sensors/"

		received := make(chan []byte)

		subscriber.Subscribe(topic, 1, func(client mqtt.Client, message mqtt.Message) {
			received <- message.Payload()
		})

		message := "test payload"
		publisher.Publish(topic, 1, false, message)

		select {
		case payload := <-received:
			if string(payload) != message {
				t.Errorf("Received payload %s, expected %s", payload, message)
			}
		case <-time.After(10 * time.Second):
			t.Error("Timeout: Did not receive the payload")
		}
	})

}