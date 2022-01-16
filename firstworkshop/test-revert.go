package firstworkshop

import (
	"fmt"
	"log"
	"sync"
)

func RevertTest() bool {
	mock := 0
	invalid := []string{}
	valid := []string{}
	o := "word"
	mu := new(sync.Mutex)
	isLimitExceed := false
	var err error

	if mock > 2 {
		fmt.Println("if isLimitExceed")
		mu.Lock()
		invalid = append(invalid, o)
		// channelCheckIsLimitExceed <- true
		isLimitExceed = true
		mu.Unlock()
	} else if err != nil {
		fmt.Println("if invalid")
		log.Printf("[SubmitOrder] [ERROR] 'refID: %s' '[%+v]' create order through fillgoods failed %s\n", refID, req, err.Error())
		mu.Lock()
		invalid = append(invalid, o)
		// channelCheckIsLimitExceed <- false
		isLimitExceed = false
		mu.Unlock()
	} else {
		fmt.Println("else valid")
		mu.Lock()
		valid = append(valid, o)
		isLimitExceed = false
		mock++
		// channelCheckIsLimitExceed <- false
		mu.Unlock()
	}
	return isLimitExceed
}
