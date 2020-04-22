package main

import (
	"sync"
)

func runConcurrentOperations(userId int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		deleteUser(userId)
	}()

	go func() {
		defer wg.Done()
		updateUserPassword(userId, "pass123")
	}()

	wg.Wait()
}
