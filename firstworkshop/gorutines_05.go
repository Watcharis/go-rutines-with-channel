package firstworkshop

import (
	"fmt"
	"sync"
	"time"
)

// Wait for a key to appear on the syncMap
func waitSync(syncMap *sync.Map, key string) {
	fmt.Println("key ->", key)
	for true {
		if checkMap, ok := syncMap.Load(key); ok {
			fmt.Println("if")
			fmt.Println("checkMap ->", checkMap)
			syncMap.Delete(key)
			break
		} else {
			fmt.Println("else")
			time.Sleep(time.Duration(1) * time.Second)
		}
	}
}

func SyncMapTutorial() {

	var syncMap sync.Map

	waitChan := make(chan bool)

	go func() {
		var i int = 1
		for true {
			waitSync(&syncMap, "a")
			fmt.Printf("%d\n", i)
			i += 3
			syncMap.Store("b", true)
		}
	}()

	go func() {
		var i int = 2
		for true {
			waitSync(&syncMap, "b")
			fmt.Printf("%d\n", i)
			i += 3
			syncMap.Store("c", true)
		}
	}()

	go func() {
		var i int = 3
		for true {
			waitSync(&syncMap, "c")
			fmt.Printf("%d\n", i)
			i += 3
			syncMap.Store("a", true)
		}
	}()

	syncMap.Store("a", true) // Kick-start things

	<-waitChan
}
