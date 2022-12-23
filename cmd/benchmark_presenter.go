package cmd

import (
	"fmt"

	"github.com/jocmp/gaffe"
)

type BenchmarkPresenter struct {
	Timestamp string
	Download  string
	Upload    string
}

type BenchmarksPresenter struct {
	Benchmarks []BenchmarkPresenter
}

func PresentBenchmarks(benchmarks []gaffe.Benchmark) BenchmarksPresenter {
	presentedBenchmarks := []BenchmarkPresenter{}
	var presentedBenchmark BenchmarkPresenter

	for _, benchmark := range benchmarks {
		presentedBenchmark = BenchmarkPresenter{
			Timestamp: benchmark.Timestamp,
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

func bitsToMegabits(value float64) float64 {
	megabits := value / (1000.0 * 1000.0)
	return megabits
}
