package main

 import "fmt"

 type person struct {
	firstName string
	lastName string
	contactInfo
 }

 type contactInfo struct {
	phone string
	email string
 }


func main() {
	vinh := person {
		firstName: "vinh",
		lastName: "Hoang",
		contactInfo: contactInfo{
			phone: "0123456",
			email: "drvinhht@gmail.com",
		},
	}
	vinh.changeName("james")
	fmt.Printf("%+v",vinh)
	fmt.Println(vinh.info())
	
}

func (p person) info() string {
	return "My name is " + p.firstName + p.lastName  
}

func (pointer *person) changeName(name string) {
	(*pointer).firstName = name
}


