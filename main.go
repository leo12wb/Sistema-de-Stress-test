package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	StatusCode int
}

func worker(url string, wg *sync.WaitGroup, results chan<- Result) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		results <- Result{StatusCode: 0}
		return
	}
	defer resp.Body.Close()

	results <- Result{StatusCode: resp.StatusCode}
}

func main() {
	// Parâmetros de CLI
	url := flag.String("url", "", "URL do serviço a ser testado")
	requests := flag.Int("requests", 1, "Número total de requests")
	concurrency := flag.Int("concurrency", 1, "Número de chamadas simultâneas")
	flag.Parse()

	// Validação de parâmetros
	if *url == "" || *requests <= 0 || *concurrency <= 0 {
		flag.Usage()
		return
	}

	// Canal para resultados
	results := make(chan Result, *requests)
	var wg sync.WaitGroup

	startTime := time.Now()

	// Execução das goroutines
	for i := 0; i < *requests; i++ {
		wg.Add(1)
		go worker(*url, &wg, results)
		if i % *concurrency == 0 {
			wg.Wait()
		}
	}

	wg.Wait()
	close(results)

	// Coleta dos resultados
	var totalRequests int
	var successRequests int
	statusCodeCounts := make(map[int]int)

	for result := range results {
		totalRequests++
		if result.StatusCode == http.StatusOK {
			successRequests++
		}
		statusCodeCounts[result.StatusCode]++
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)

	// Relatório
	fmt.Printf("Tempo total: %v\n", duration)
	fmt.Printf("Total de requests: %d\n", totalRequests)
	fmt.Printf("Requests com status 200: %d\n", successRequests)
	fmt.Println("Distribuição dos códigos de status HTTP:")
	for code, count := range statusCodeCounts {
		fmt.Printf("%d: %d\n", code, count)
	}
}
