package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting batch")
	time.Sleep(3 * time.Second)
	fmt.Println("Fetching from  API")
	time.Sleep(7 * time.Second)
	fmt.Println("Processing")
	time.Sleep(2 * time.Second)
	fmt.Println("Error produyct XYZ is too cheap")
	time.Sleep(12 * time.Second)
	fmt.Println("Done")
}
