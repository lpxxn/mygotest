package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
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

	type t struct {
		a *int
		b int
	}
	chList := make(chan t, 3)
	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			chList <- t{a: nil, b: i}
		}(i)
	}
	wg.Wait()
	close(chList)
	for i := 0; i < 3; i++ {
		fmt.Println("value: ", <-chList)
	}
	length := 16
	snowflakeIDStr := strconv.FormatUint(uint64(262683231892309619), 16)
	fmt.Println("snowflakeIDStr: ", snowflakeIDStr, len(snowflakeIDStr))
	curLength := len(snowflakeIDStr)
	id, err := Rand(length - curLength)
	if err != nil {
		snowflakeIDStr += strings.Repeat("0", length-curLength)
		fmt.Println("snowflakeIDStr err: ", snowflakeIDStr)
	}
	fmt.Println("snowflakeIDStr +id: ", snowflakeIDStr+id)
}

var (
	numbers        = "0123456789"
	numberAndAlpha = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

// GenerateRandomBytes of n size
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func Rand(n int) (string, error) {
	bytes, err := GenerateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = numberAndAlpha[b%byte(len(numberAndAlpha))]
	}
	return string(bytes), nil
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
