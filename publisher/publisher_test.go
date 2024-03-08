package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	DefaultClient "mqtt/src/common"
)

var client = DefaultClient.CreateClient(DefaultClient.Broker, DefaultClient.IdPublisher, DefaultClient.Handler)

func TestMain(t *testing.T) {
	t.Run("Create new Sensor (freezer)", func(t *testing.T) {
		sensor := NewSensor("lj01f01","freezer", -30.0, "01/03/2024", 1)
		compare := &Sensor{Id: "lj01f01", Tipo: "freezer", Temperatura: -30.0, Dia: "01/03/2024", Hora: 1}
		if !reflect.DeepEqual(sensor, compare) {
			t.Errorf("The sensor was not created successfully...")
		}
	})

	t.Run("Create new Sensor (geladeira)", func(t *testing.T) {
		sensor := NewSensor("ab04g08", "geladeira", 20.0, "01/03/2024", 1)
		compare := &Sensor{Id: "ab04g08", Tipo: "geladeira", Temperatura: 20.0, Dia: "01/03/2024", Hora: 1}
		if !reflect.DeepEqual(sensor, compare) {
			t.Errorf("The sensor was not created successfully...")
		}
	})

	t.Run("Generating JSON file to payload (freezer)", func(t *testing.T) {
		sensor := NewSensor("lj01f01","freezer", -30.0, "01/03/2024", 1)
		got, err := sensor.ToJSON()
		var transformed map[string]interface{}
		json.Unmarshal([]byte(got), &transformed)
		if err != nil {
			t.Errorf("Error generating JSON: %v", err)
		}

		want := map[string]interface{}{
			"Id":			"lj01f01",
			"Tipo":			"freezer",
			"Temperatura":  -30.0,
			"Dia":			"01/03/2024",
			"Hora":			1,
		}

		if !(fmt.Sprint(transformed) == fmt.Sprint(want)) {
			t.Errorf("Unexpected JSON output.\nGot: %v\nWant: %v", transformed, want)
		}
	})

	t.Run("Generating JSON file to payload (geladeira)", func(t *testing.T) {
		sensor := NewSensor("ab04g08", "geladeira", 20.0, "01/03/2024", 1)
		got, err := sensor.ToJSON()
		var transformed map[string]interface{}
		json.Unmarshal([]byte(got), &transformed)
		if err != nil {
			t.Errorf("Error generating JSON: %v", err)
		}

		want := map[string]interface{}{
			"Id":			"ab04g08",
			"Tipo":			"geladeira",
			"Temperatura":  20.0,
			"Dia":			"01/03/2024",
			"Hora":			1,
		}

		if !(fmt.Sprint(transformed) == fmt.Sprint(want)) {
			t.Errorf("Unexpected JSON output.\nGot: %v\nWant: %v", transformed, want)
		}
	})

	t.Run("Test QoS - eg if the message was published by the broker", func(t *testing.T) {
		payload := "Hello, Broker!"
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			t.Error(token.Error())
		}

		token := client.Publish("sensors", 1, false, payload)

		if token.Wait() && token.Error() != nil {
			t.Error(token.Error())
		}

		t.Log("Broker received message with QoS 1!")
	})
}