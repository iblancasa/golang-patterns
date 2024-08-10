package main

import (
	"fmt"
	"log"
	"sync"
)

type Singleton struct {
	Value map[string]string
}

var (
	instance Singleton
	once     sync.Once
)

func GetSingleton() Singleton {
	once.Do(
		func() {
			log.Println("Create the object")
			instance = Singleton{
				Value: map[string]string{},
			}
		},
	)
	return instance
}

func main() {
	i := GetSingleton()
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		// Call multiple times the creation to show the object is not recreated
		go func() {
			defer wg.Done()
			GetSingleton()
		}()
	}
	wg.Wait()

	i.Value["key1"] = "value1"

	h := GetSingleton()
	if _, ok := h.Value["key1"]; ok {
		fmt.Println("The map is the same")
	}
}
