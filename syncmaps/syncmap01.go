package syncmaps

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type SyncMaps struct {
	Sm *sync.Map
	Wg *sync.WaitGroup
	Mu *sync.Mutex
}

var (
	m sync.Map
	n map[int]string
)

func init() {
	m = sync.Map{}
	n = map[int]string{}
}

func (s *SyncMaps) TutorialOne() {

	maxRound := 10
	limit := 0

	for round := 0; round < maxRound; round++ {

		s.Wg.Add(1)
		go func() {

			defer func() {
				s.Wg.Done()
				s.Mu.Unlock()
			}()

			s.Mu.Lock()
			limit++
			// fmt.Println("limit ->", limit)

			rand.Seed(time.Now().UnixNano())
			key := rand.Intn(10)
			// fmt.Println("key ->", key)

			// if _, ok := m[key]; ok {
			// 	delete(m, key)
			// } else {
			// 	data := time.Now().Format(time.RFC3339)
			// 	fmt.Println("data ->", data)
			// 	m[key] = data
			// }

			if rand.Intn(2)%2 == 0 {
				data := time.Now().Format(time.RFC3339)
				fmt.Printf("key: %d || data: %v\n", key, data)
				n[key] = data
			} else {
				if _, ok := n[key]; ok {
					delete(n, key)
				}
			}
		}()

	}
	s.Wg.Wait()

}

func (s *SyncMaps) TutorialTwo() *sync.Map {
	n := 10

	for n > 0 {

		if n == 0 {
			break
		} else {
			n--
		}

		s.Wg.Add(1)
		go func() {
			defer func() {
				s.Wg.Done()
			}()
			rand.Seed(time.Now().UnixNano())
			key := rand.Intn(10)
			fmt.Println("key ->", key)

			if rand.Intn(2)%2 == 0 {
				m.Store(key, time.Now().Format(time.RFC1123))
			} else {
				if _, ok := m.Load(key); ok {
					m.Delete(key)
				}
			}
		}()
	}
	s.Wg.Wait()
	return &m
}

func (s *SyncMaps) ReadMap(m *sync.Map) {
	fmt.Printf("m -> %+v\n", m)
}
