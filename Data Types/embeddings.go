package main

import "fmt"

// Processor struct
type Processor struct {
	processorName string
	cores         int
}

// Memory struct
type Memory struct {
	memorySize int
	memoryType string
}

// Computer struct
type Computer struct {
	Processor
	Memory
	price int
}

func main() {
	computer := Computer{} //Initialize with default values

	fmt.Println(computer)

	computer.processorName = "Intel i5"
	computer.cores = 6
	computer.memorySize = 8
	computer.memoryType = "DDR4"
	computer.price = 50000

	fmt.Println(computer)

	// Another method of initializing
	computer1 := Computer{
		Processor: Processor{
			processorName: "Ryzen 5",
			cores:         8,
		},
		Memory: Memory{
			memorySize: 8,
			memoryType: "DDR4",
		},
		price: 60000,
	}
	fmt.Println(computer1)
	//Get and set single property
	computer.memorySize = 16
	fmt.Println(computer.memorySize)
	fmt.Println(computer1.Processor)

}
