package main

import (
	"sync"
	"testing"
)

func AddToMap(data *sync.Map, key int, value string, group *sync.WaitGroup) {
	defer group.Done()
	group.Add(1)
	data.Store(key, value)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go AddToMap(data, i, "value", group)
	}
	group.Wait()

	data.Range(func(key, value interface{}) bool {
		t.Logf("Key: %v, Value: %v\n", key, value)
		return true
	})
}
