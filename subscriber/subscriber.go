package main

import (
	"fmt"
	DefaultClient "mqtt/src/common"
)

func main() {
	
	client := DefaultClient.CreateClient(DefaultClient.Broker, DefaultClient.IdSubscriber, DefaultClient.Handler)
	
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe("sensors/", 1, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}

	fmt.Println("Subscriber done!")
	select {}
}