package gaffe

import (
	"os"

	"github.com/gocarina/gocsv"
)

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
