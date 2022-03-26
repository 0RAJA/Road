package singleflight

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGroup_Do(t *testing.T) {
	redis := NewGroup()
	wg := new(sync.WaitGroup)
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(n int) {
			defer wg.Done()
			result, err := redis.Do("redis", func() (interface{}, error) {
				time.Sleep(time.Second)
				return n, nil
			})
			if err != nil {
				fmt.Println("err:", err)
				return
			}
			fmt.Println(result)
		}(i)
	}
	wg.Wait()
}
