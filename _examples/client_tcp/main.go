package main

import (
	"fmt"
	"time"

	modbus "github.com/torykit/go-modbus"
)

func main() {
	p := modbus.NewTCPClientProvider("10.211.55.3:502", modbus.WithEnableLogger())
	client := modbus.NewClient(p)
	err := client.Connect()
	if err != nil {
		fmt.Println("connect failed, ", err)
		return
	}
	defer client.Close()

	fmt.Println("starting")
	for {
		results, err := client.ReadHoldingRegisters(1, 0, 10)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(results)

		//	fmt.Printf("ReadDiscreteInputs %#v\r\n", results)

		time.Sleep(time.Second * 2)
	}
}
