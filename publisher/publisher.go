package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
	DefaultClient "mqtt/src/common"
)

const maxTemperaturaFreezer = -15.0
const minTemperaturaFreezer = -25.0
const maxTemperaturaGeladeira = 10.0
const minTemperaturaGeladeira = 2.0

type Sensor struct {
	Id        	string
	Tipo	    string
	Temperatura float64
	Dia			string
	Hora		float64
}

func NewSensor(
	id string,
	tipo string,
	temperatura float64,
	dia string,
	hora float64) *Sensor {

	s := &Sensor{
		Id:				id,
		Tipo:			tipo,
		Temperatura:	temperatura,
		Dia:			dia,
		Hora:			hora,
	}

	return s

}

func (s *Sensor) ToJSON() (string, error) {
	jsonData, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func main() {

	sensor := NewSensor("lj01f01","freezer", -30.0, "01/03/2024", 1) //freezer

	sensor2 := NewSensor("ab04g08", "geladeira", 20.0, "01/03/2024", 1) //geladeira

	var sensors []Sensor
	sensors = append(sensors, *sensor, *sensor2)

	client := DefaultClient.CreateClient(DefaultClient.Broker, DefaultClient.IdPublisher, DefaultClient.Handler)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for {
		for _, sensor := range sensors {
			topic := "sensors/" + sensor.Id
			sensor.Temperatura = (rand.Float64() * (maxTemperaturaFreezer - minTemperaturaFreezer)) + minTemperaturaFreezer
			payload, _ := sensor.ToJSON()
			token := client.Publish(topic, 0, false, payload)
			token.Wait()
			fmt.Printf("Temperatura Baixa Freezer: %s\n", payload)
			time.Sleep(time.Duration(sensor.Hora) * time.Second)
		}

		for _,sensor := range sensors {
			topic := "sensors/" + sensor.Id
			sensor.Temperatura = (rand.Float64() * (maxTemperaturaGeladeira - minTemperaturaGeladeira)) + minTemperaturaGeladeira
			payload, _ := sensor.ToJSON()
			token := client.Publish(topic, 0, false, payload)
			token.Wait()
			fmt.Printf("Temperatura Alta Geladeira: %s\n", payload)
			time.Sleep(time.Duration(sensor.Hora) * time.Second)
		}
	}
}