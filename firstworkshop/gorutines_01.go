package firstworkshop

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"sync"
)

func ChanelOne(data chan int, wg *sync.WaitGroup) {
	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, i := range number {
		data <- i
	}
	close(data)
}

func Publisher(testTwoChan chan<- int) {
	for i := 0; i < 10; i++ {
		// fmt.Println("Publisher ->", i)
		testTwoChan <- i
	}
	close(testTwoChan)
}

func Subscriber(testTwoChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := range testTwoChan {
		fmt.Printf("Subscriber received: %d\n", i)
	}
}

func Gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func Sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			fmt.Println("Sq n ->", n)
			out <- n * n
		}
		close(out)
	}()
	return out
}

func Merge(cs ...<-chan int) (res <-chan int) {
	var wg sync.WaitGroup
	out := make(chan int)
	fmt.Println("cs lenght:", len(cs))

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))

	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

type data struct {
	Body struct {
		Message string `json:"message"`
		UserID  int    `json:"userId"`
		ID      int    `json:"id"`
		Title   string `json:"title"`
	}
	err error
}

func Request(method string, url string, body io.Reader) (<-chan data, <-chan string) {
	requestError := make(chan string)

	r, err := http.NewRequest(method, url, body)
	if err != nil {
		requestError <- err.Error()
		return nil, requestError
	}

	future := make(chan data)
	go func() {
		res, err := (&http.Client{}).Do(r)
		if err != nil {
			future <- data{err: err}
			return
		}
		defer res.Body.Close()

		var d data
		d.err = json.NewDecoder(res.Body).Decode(&d.Body)
		future <- d
		close(future)
	}()

	return future, nil
}

//higher Order Function
//func คืน func --->
func Sphere(num int) (func(radius float64) float64, func(text string) string) {

	resultFloat := func(radius float64) float64 {
		volume := 4 / 3 * math.Pi * radius * radius * radius
		return volume
	}

	resultString := func(text string) string {
		return text
	}

	return resultFloat, resultString
}

func MultipleCalculate() func(n int) func(c int) int {
	return func(n int) func(c int) int {
		return func(i int) int {
			return n * i
		}
	}
}

//closure function
func XX() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
