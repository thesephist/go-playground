package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func avg(nums ...float64) float64 {
	var sum float64 = 0.0
	for _, n := range nums {
		sum += n
	}
	return sum / float64(len(nums))
}

func runEstimate(iterations int, sink chan<- float64) {
	inCount := 0

	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	for i := 0; i < iterations; i++ {
		if math.Pow(r.Float64(), 2)+math.Pow(r.Float64(), 2) < 1 {
			inCount++
		}
	}

	sink <- (float64(inCount) / float64(iterations)) * 4
}

func main() {
	iterCountPtr := flag.Int("iterations", 100000, "number of Monte-Caro iterations")
	concurrencyCountPtr := flag.Int("concurrency", 1, "number of goroutines to use")

	flag.Parse()

	iterPerSubroutine := *iterCountPtr / *concurrencyCountPtr
	sink := make(chan float64)

	for i := 0; i < *concurrencyCountPtr; i++ {
		go runEstimate(iterPerSubroutine, sink)
	}

	var piEstimates []float64
	for i := 0; i < *concurrencyCountPtr; i++ {
		est := <-sink
		piEstimates = append(piEstimates, est)
		fmt.Printf("Estimate from thread #%d: %.10f\n", i, est)
	}

	piApprox := avg(piEstimates...)

	fmt.Printf("Estimate of Pi after %d iterations: %.10f\n",
		iterPerSubroutine*(*concurrencyCountPtr), piApprox)
}
