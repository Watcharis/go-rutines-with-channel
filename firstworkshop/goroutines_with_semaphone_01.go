package firstworkshop

import (
	"fmt"
	"sync"
	"time"
)

func PrepareData() [][]string {
	tokenTest := "TEST_TOKEN"
	dataStore := [][]string{}
	storeToken := []string{}
	n := 500
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if len(storeToken) <= n {
				storeToken = append(storeToken, tokenTest)
			}
		}
		if len(storeToken) == n {
			dataStore = append(dataStore, storeToken)
		}
		storeToken = nil
	}
	return dataStore
}

func logData(index int, value []string, sem chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("index: %+v | token: %+v\n ", index, value)
	<-sem
}

func Semaphone() error {
	startTime := time.Now()
	fmt.Println("startTime ->", startTime)

	sem := make(chan int, 10)
	initToken := PrepareData()
	fmt.Println("initToken ->", len(initToken))

	wg := sync.WaitGroup{}

	for i, v := range initToken {
		sem <- 1
		wg.Add(1)
		go logData(i, v, sem, &wg)
	}
	wg.Wait()

	endTime := time.Since(startTime)
	fmt.Println("endTime ->", endTime)

	return nil
}
