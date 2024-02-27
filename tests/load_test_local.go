package tests

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestLoadLocal(t *testing.T) {
	var wg sync.WaitGroup
	zipcode := "08673000"
	totalRequests := 20
	requestsPerSecond := 1
	wg.Add(totalRequests)

	makeRequest := func() {
		defer wg.Done()

		url := fmt.Sprintf("http://localhost:8080/weather?zipcode=%s", zipcode)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Printf("Erro ao criar a requisição: %v\n", err)
			return
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("Erro na requisição: %v\n", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}

			message := string(body)
			fmt.Printf("Result: %s\n", message)
		}
	}

	for i := 0; i < totalRequests; i++ {
		go makeRequest()
		time.Sleep(time.Second / time.Duration(requestsPerSecond))
	}

	wg.Wait()
}
