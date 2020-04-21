package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

const userTable = "database/user.yaml"

type User struct {
	Id int
	Name string
	Password string
}

func readFile(path string) []byte {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return data
}

func writeFile(path string, data []byte) error {
	err := ioutil.WriteFile(path, data, 0644)

	if err != nil {
		log.Fatalf("error: %v", err)
		return err
	}

	return nil
}

func updateUserPassword(userId int, password string) {
	users := getUsers()
	for i, user := range users {
		if user.Id == userId {
			users[i].Password = password
			saveUsers(users)
			return
		}
	}
}

func deleteUser(userId int) {
	users := getUsers()
	for i, user := range users {
		if user.Id == userId {
			users[i] = users[len(users) - 1]
			users[len(users)-1] = User{}
			users = users[:len(users) - 1]
			saveUsers(users)
			return
		}
	}
}

func saveUsers(users []User) {
	d, err := yaml.Marshal(&users)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = writeFile(userTable, d)

	if err != nil {
		log.Fatal(err)
	}
}

func getUsers() []User {
	var users []User
	data := readFile(userTable)
	err := yaml.Unmarshal(data, &users)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return users
}

func main()  {
	deleteUser(3)
	//go updateUserPassword(3, "Hasan61")
}
