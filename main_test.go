package main

import (
	"io/ioutil"
	"log"
	"testing"
)

func restore() {
	users, err := ioutil.ReadFile(readonlyUserTable)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(userTable, users, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func TestConcurrentUserActions(t *testing.T) {
	restore()
	// given
	userId := 3

	// when
	runConcurrentOperations(userId)

	// then
	users := getUsers()
	if isUserExists(users, userId) {
		t.Errorf("Test Failed: User exists even though delete method runned first.")
	}
}
