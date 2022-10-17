package firstworkshop

import (
	"fmt"
	"sync"
)

func Test10() {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	store := []map[int]int{}

	testCH := make(chan int)

	wg.Add(1)
	go func(wg *sync.WaitGroup, mu *sync.Mutex) {
		defer wg.Done()

		ii := 0
		for ii < 5 {

			n := 10
			i := 0
			for i < n {

				wg.Add(1)
				go func(i int) {
					defer wg.Done()
					test := map[int]int{i: i}
					mu.Lock()
					store = append(store, test)
					mu.Unlock()
					testCH <- test[i]

				}(i)
				i++
			}
			ii++
		}
	}(&wg, &mu)

	// wg.Add(1)
	go func(wg *sync.WaitGroup) {

		// defer wg.Done()

		i := 0
		for data := range testCH {
			i++

			go func(data int, i int) {

				fmt.Printf("data -> %d || i -> %d\n", data, i)
			}(data, i)
		}
	}(&wg)

	wg.Wait()
	close(testCH)
	// fmt.Println("store ->", store)
}
