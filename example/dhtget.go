package main

import (
	"github.com/9nut/dht"
	"flag"
	"fmt"
)

func main() {
	sensor := flag.Int("t", 22, "Sensor type (11 or 22)")
	pin := flag.Int("p", 17, "RPi pin number")
	flag.Parse()

	temp, humid, err := dht.DHTRead(*sensor, *pin)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(temp, humid)
	}
}
