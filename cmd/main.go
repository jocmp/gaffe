package main

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type Benchmark struct {
	ServerID   int64   `csv:"Server ID"`
	Sponsor    string  `csv:"Sponsor"`
	ServerName string  `csv:"Server Name"`
	Timestamp  string  `csv:"Timestamp"`
	Distance   string  `csv:"Distance"`
	Ping       float64 `csv:"Ping"`
	Download   float64 `csv:"Download"` // Download speed in bits
	Upload     float64 `csv:"Upload"`   // Upload speed in bits
	IPAddress  string  `csv:"IP Address"`
}

func main() {
	file, err := os.OpenFile("speedtest.csv", os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	benchmarks := []*Benchmark{}

	if err := gocsv.UnmarshalFile(file, &benchmarks); err != nil {
		panic(err)
	}

	if _, err := file.Seek(0, 0); err != nil { // Go to the start of the file
		panic(err)
	}

	content, err := gocsv.MarshalString(&benchmarks)

	if err != nil {
		panic(err)
	}

	fmt.Println(content) // Display all clients as CSV string
}
