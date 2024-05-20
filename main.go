package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	TotalRequests        int
	SuccessfulRequests   int
	FailedRequests       int
	StatusDistribution   map[int]int
	TotalTime            time.Duration
}

func runLoadTest(url string, totalRequests, concurrency int) Result {
	start := time.Now()

	results := Result{
		TotalRequests:      totalRequests,
		StatusDistribution: make(map[int]int),
	}

	var wg sync.WaitGroup
	ch := make(chan int, concurrency)

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- 1

			resp, err := http.Get(url)
			if err != nil {
				results.FailedRequests++
			} else {
				results.SuccessfulRequests++
				results.StatusDistribution[resp.StatusCode]++
			}

			<-ch
		}()
	}

	wg.Wait()
	results.TotalTime = time.Since(start)
	return results
}

func generateReport(results Result) {
	fmt.Println("Load Test Report:")
	fmt.Println("Total Time:", results.TotalTime)
	fmt.Println("Total Requests:", results.TotalRequests)
	fmt.Println("Successful Requests:", results.SuccessfulRequests)
	fmt.Println("Failed Requests:", results.FailedRequests)
	fmt.Println("Status Distribution:")
	for status, count := range results.StatusDistribution {
		fmt.Printf("%d: %d\n", status, count)
	}
}

func main() {
	var url string
	var totalRequests, concurrency int

	flag.StringVar(&url, "url", "", "URL do serviço a ser testado.")
	flag.IntVar(&totalRequests, "requests", 0, "Número total de requests.")
	flag.IntVar(&concurrency, "concurrency", 1, "Número de chamadas simultâneas.")

	flag.Parse()

	if url == "" || totalRequests == 0 || concurrency == 0 {
		flag.PrintDefaults()
		return
	}

	results := runLoadTest(url, totalRequests, concurrency)
	generateReport(results)
}
