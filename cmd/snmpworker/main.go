package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("SNMP worker starting...")

	for {
		fmt.Println("Querying SNMP Devices...")
		time.Sleep(10 * time.Second) // sleep for 10 seconds, then "query" again
	}

}
