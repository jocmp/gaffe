package gaffe

import (
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

func FetchBenchmarks(path string) ([]Benchmark, error) {
	file, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	benchmarks := []Benchmark{}

	if err := gocsv.UnmarshalFile(file, &benchmarks); err != nil {
		panic(err)
	}

	if _, err := file.Seek(0, 0); err != nil {
		panic(err)
	}

	return benchmarks, nil
}
