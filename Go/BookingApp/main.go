package main

import (
	"fmt"
	"time"
	"sync"
)

type UserData struct {
	firstName string
	lastName string
	age int
}
func PrintFirstNames(bookings []UserData) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		name := booking.firstName
		firstNames = append(firstNames, name)
	}
	return firstNames
}

var wg = sync.WaitGroup{}

func hello() {
    user := UserData {
		firstName: "vinh",
		lastName: "Hoang",
		age: 29,
	}
	var names = []UserData{user,user, user}
	var firstNames = PrintFirstNames(names)
	for _, name := range firstNames {
		fmt.Println(name)
	}
}

func sendingEmail() {
	fmt.Println("Sending email")
	time.Sleep(5 * time.Second)
	fmt.Println("Success!")
	wg.Done()
}

func main() {
	var name string
	fmt.Printf("Enter your name: ")
	fmt.Scan(&name)
	hello()
	wg.Add(1)
	go sendingEmail()
	fmt.Println("Goodbye!")
	wg.Wait()
}
