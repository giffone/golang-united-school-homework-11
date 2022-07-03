package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	mu := sync.Mutex{}
	ch := make(chan struct{}, pool)
	wg := sync.WaitGroup{}
	for i := int64(0); i < n; i++ {
		wg.Add(1)
		ch <- struct{}{}
		go func(id int64) {
			user := getOne(id)
			mu.Lock()
			res = append(res, user)
			mu.Unlock()
			<-ch
			wg.Done()
		}(i)
	}
	wg.Wait()
	return
}
