package firstworkshop

import (
	"errors"
	"fmt"
	"sync"
)

func plus(lastItem int, item int) (int, *NumberError) {

	if item == lastItem {
		numError := errors.New("not process last item")
		resultError := NumberError{
			ItemErr:    item,
			ErrMessage: numError,
		}
		return 0, &resultError
	}

	mod := item%2 == 0
	if mod {
		multiply := item * 2
		return multiply, nil
	}

	return item, nil
}

func retry(item int, number chan int, errNumber chan NumberError) {

	maxRetry := 3
	i := 0
	for i < maxRetry {
		fmt.Println("round of retry :", i)
		resultNumber, errNum := plus(maxRetry, item)
		if errNum != nil {
			errNumber <- *errNum
		} else {
			fmt.Println("resultNumber ->", resultNumber)
			number <- resultNumber
		}
		i++
	}

	close(number)
	close(errNumber)
}

type NumberError struct {
	ItemErr    int
	ErrMessage error
}

func CheckDuplicateInArray(storageNumber []int, item int) bool {
	for _, v := range storageNumber {
		if v == item {
			return true
		}
	}
	return false
}

func BeforeTest12() ([]int, []NumberError) {

	wg := new(sync.WaitGroup)

	number := make(chan int, 1)
	errNumber := make(chan NumberError, 1)
	retryCH := make(chan NumberError, 1)
	storageNumber := []int{}
	// storageNumberCH := make(chan []int, 1)
	// storageErrorCH := make(chan []NumberError, 1)

	n := 10
	for i := 0; i <= n; i++ {

		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			resultNumber, errNum := plus(n, i)
			if errNum != nil {
				errNumber <- *errNum
			} else {
				number <- resultNumber
			}
		}(i)
	}

	wg.Add(1)
	go func() {

		defer func() {
			wg.Done()
		}()

		for val := range number {
			fmt.Printf("val -> %+v\n", val)
			if !CheckDuplicateInArray(storageNumber, val) {
				// fmt.Println("not duplicate")
				storageNumber = append(storageNumber, val)
			}
		}
		fmt.Println("------ number ----------")
	}()

	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			close(retryCH)
		}()

		for em := range errNumber {

			fmt.Printf("errNumber -> %+v\n", em.ErrMessage)
			retryCH <- em
		}
		fmt.Println("------ error ----------")
	}()

	wg.Add(1)
	go func() {

		defer func() {
			wg.Done()
		}()

		for rCH := range retryCH {
			retry(rCH.ItemErr, number, errNumber)
		}
		fmt.Println("------ retry ----------")
	}()

	wg.Wait()
	fmt.Println("storageNumber ->", storageNumber)

	return storageNumber, []NumberError{}
}
