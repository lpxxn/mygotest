package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

func main() {
	f := &FullJitterStrategy{}
	for i := 1; i < 30; i++ {
		t := f.Calculate(i)
		fmt.Printf("retires: %d, duration: %f \n", i, t.Seconds())
	}

	fmt.Println("---------")
	e := &ExponentialStrategy{}
	for i := 1; i < 30; i++ {
		t := e.Calculate(i)
		fmt.Printf("retires: %d, duration: %f \n", i, t.Seconds())
	}
}

var BackoffMultiplier = time.Second

// FullJitterStrategy implements http://www.awsarchitectureblog.com/2015/03/backoff.html
type FullJitterStrategy struct {
	rngOnce sync.Once
	rng     *rand.Rand
}

// Calculate returns a random duration of time [0, 2 ^ attempt]
func (s *FullJitterStrategy) Calculate(attempt int) time.Duration {
	// lazily initialize the RNG
	s.rngOnce.Do(func() {
		if s.rng != nil {
			return
		}
		s.rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	})

	backoffDuration := BackoffMultiplier *
		time.Duration(math.Pow(2, float64(attempt)))

	//if backoffDuration < 0 {
	//	return 0
	//}
	fmt.Print("---: ", backoffDuration, "  ")
	return time.Duration(s.rng.Intn(int(backoffDuration)))
}

// ---------------
// ExponentialStrategy implements an exponential backoff strategy (default)
type ExponentialStrategy struct {
}

// Calculate returns a duration of time: 2 ^ attempt
func (s *ExponentialStrategy) Calculate(attempt int) time.Duration {
	backoffDuration := BackoffMultiplier *
		time.Duration(math.Pow(2, float64(attempt)))
	return backoffDuration
}
