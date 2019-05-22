package mqttrouter

import (
	"github.com/eclipse/paho.mqtt.golang"
	"fmt"
)

func standardHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}
