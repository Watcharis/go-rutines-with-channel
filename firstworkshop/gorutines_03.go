package firstworkshop

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

func RaceCondition01() map[string]string {

	// Example RaceConditions

	// Reffernce https://go.dev/doc/articles/race_detector

	// func main() {
	// 		c := make(chan bool)
	// 		m := make(map[string]string)
	// 		var mu sync.Mutex

	// 		go func() {
	// 			mu.Lock()
	// 			m["1"] = "a" // First conflicting access.
	// 			c <- true
	// 			mu.Unlock()
	// 		}()
	// 		m["2"] = "b" // Second conflicting access.
	// 		<-c
	// 		for k, v := range m {
	// 			fmt.Println(k, v)
	// 		}
	// }

	start := time.Now()

	c := make(chan bool, 2)
	m := make(map[string]string)

	wg := new(sync.WaitGroup) // var wg sync.WaitGroup
	mu := new(sync.Mutex)     // var m sync.Mutex

	wg.Add(1)
	go func(wg *sync.WaitGroup, mu *sync.Mutex) {
		defer wg.Done()
		mu.Lock()
		m["1"] = "a" // First conflicting access.
		c <- true    // Write channel
		mu.Unlock()
	}(wg, mu)

	isTrue := <-c
	if isTrue {
		m["2"] = "b"
	}

	wg.Add(1)
	go func(wg *sync.WaitGroup, mu *sync.Mutex, isTrue bool) {
		defer wg.Done()
		mu.Lock()
		m["3"] = "c" // First conflicting access.
		c <- isTrue  // Write channel
		mu.Unlock()
	}(wg, mu, isTrue)

	wg.Add(1)
	go func(wg *sync.WaitGroup, mu *sync.Mutex, isTrue bool) {
		defer wg.Done()
		mu.Lock()
		m["4"] = "d" // First conflicting access.
		c <- isTrue  // Write channel
		mu.Unlock()
	}(wg, mu, isTrue)

	wg.Wait()

	for k, v := range m {
		fmt.Printf("key: %s || val: %s\n", k, v)
	}

	end := time.Since(start)
	fmt.Printf("end: %.10f\n", float64(end.Seconds()))
	return m
}

func CheckSpeedDelareMap() map[string]string {
	start := time.Now()
	m := make(map[string]string)
	m["1"] = "a"
	m["2"] = "b"
	m["3"] = "c"
	m["4"] = "d"
	end := time.Since(start)
	fmt.Printf("end: %.10f\n", float64(end.Seconds()))
	return m
}

func FindBirthDay() string {
	// str := "1996-01-20T11:45:26.371Z"
	str := "1996-01-20T15:04:05+07:00"
	t, err := time.Parse(time.RFC3339, str)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("BirthDay ->", t.Unix())
	fmt.Println("ShortDay : ", t.Format(time.ANSIC))
	return fmt.Sprintf("BirthDay -> %+v", t.Unix())
}

func CheckSwitchCase(grade int) string {
	// switch {
	// case grade >= 80:
	// 	return "A"
	// case grade >= 70 && grade < 80:
	// 	return "B"
	// case grade >= 60 && grade < 70:
	// 	return "C"
	// case grade >= 50 && grade < 60:
	// 	return "D"
	// default:
	// 	return "F"
	// }
	a := 1
	t := reflect.ValueOf(a).Kind()
	fmt.Println("ty ->", t)
	switch t := reflect.ValueOf(a).Kind(); t {
	case reflect.Int:
		return "int"
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "boolean"
	default:
		return t.String()
	}
	// return "ok"
}
