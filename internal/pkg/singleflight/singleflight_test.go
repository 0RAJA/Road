package singleflight

import (
	"github.com/stretchr/testify/require"
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
			_, err := redis.Do("redis", func() (interface{}, error) {
				time.Sleep(time.Second)
				return n, nil
			})
			require.NoError(t, err)
		}(i)
	}
	wg.Wait()
}
