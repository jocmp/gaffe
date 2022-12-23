package gaffe

import (
	"fmt"
	"sort"
)

type BenchmarkPresenter struct {
	Timestamp string
	Download  string
	Upload    string
}

type BenchmarksPresenter struct {
	Benchmarks []BenchmarkPresenter
}

func PresentBenchmarks(benchmarks []Benchmark) BenchmarksPresenter {
	presentedBenchmarks := []BenchmarkPresenter{}
	var presentedBenchmark BenchmarkPresenter
	sorted := sortTimeDescending(benchmarks)

	for _, benchmark := range sorted {
		presentedBenchmark = BenchmarkPresenter{
			Timestamp: benchmark.ChicagoTime().Format("2006-01-02 03:04:05 PM MST"),
			Download:  fmt.Sprintf("%f", bitsToMegabits(benchmark.Download)),
			Upload:    fmt.Sprintf("%f", bitsToMegabits(benchmark.Upload)),
		}
		presentedBenchmarks = append(presentedBenchmarks, presentedBenchmark)
	}

	presented := BenchmarksPresenter{
		Benchmarks: presentedBenchmarks,
	}

	return presented
}

func sortTimeDescending(benchmarks []Benchmark) []Benchmark {
	destination := make([]Benchmark, len(benchmarks))
	copy(destination, benchmarks)

	sort.Slice(destination, func(i, j int) bool {
		return destination[i].Timestamp > destination[j].Timestamp
	})

	return destination
}

func bitsToMegabits(value float64) float64 {
	megabits := value / (1000.0 * 1000.0)
	return megabits
}
