package firstworkshop

import (
	"sync"
)

func RevertTest() bool {
	invalid := []string{}
	valid := []string{}
	o := "word"
	mu := new(sync.Mutex)
	isLimitExceed := false
	ErrorStatusCodeLimitExceed := 1000
	var err error
	type Response struct {
		Success    bool
		StatusCode int
	}

	response := Response{}

	if !response.Success && response.StatusCode == ErrorStatusCodeLimitExceed {
		mu.Lock()
		invalid = append(invalid, o)
		isLimitExceed = true
		mu.Unlock()
	} else if err != nil {
		// log.Printf("[SubmitOrder] [ERROR] 'refID: %s' '[%+v]' create order through fillgoods failed %s\n", refID, req, err.Error())
		mu.Lock()
		invalid = append(invalid, o)
		mu.Unlock()
	} else {
		mu.Lock()
		valid = append(valid, o)
		mu.Unlock()
	}
	return isLimitExceed
}

// if mock > 2 {
// 	fmt.Println("if isLimitExceed")
// 	mu.Lock()
// 	invalid = append(invalid, o)
// 	// channelCheckIsLimitExceed <- true
// 	isLimitExceed = true
// 	mu.Unlock()
// } else if err != nil {
// 	fmt.Println("if invalid")
// 	// log.Printf("[SubmitOrder] [ERROR] 'refID: %s' '[%+v]' create order through fillgoods failed %s\n", refID, req, err.Error())
// 	mu.Lock()
// 	invalid = append(invalid, o)
// 	// channelCheckIsLimitExceed <- false
// 	isLimitExceed = false
// 	mu.Unlock()
// } else {
// 	fmt.Println("else valid")
// 	mu.Lock()
// 	valid = append(valid, o)
// 	isLimitExceed = false
// 	mock++
// 	// channelCheckIsLimitExceed <- false
// 	mu.Unlock()
// }
// return isLimitExceed
