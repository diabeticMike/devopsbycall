package main

import (
	"fmt"
)

type user struct {
	name  string
	age   uint8
	email string
}

func main() {
	users := []user{
		{
			name:  "Mike",
			age:   32,
			email: "mike@gmail.com",
		},
		{
			name:  "John",
			age:   54,
			email: "john@gmail.com",
		},
		{
			name:  "Abobus",
			age:   2,
			email: "abobus@gmail.com",
		},
	}

	//fmt.Println(users[len(users)-1])
	//fmt.Println(len(users))

	var clientInfo string = "John"
	fmt.Println("Введенное имя для поиска: " + clientInfo)
	//fmt.Println(users[0].name)

	for index := 0; index < len(users); index++ {
		if users[index].name == clientInfo {
			fmt.Println("****************Информация об клиенте****************")
			fmt.Println("Имя клиента: " + users[index].name)
			fmt.Println("Email клиента: " + users[index].email)
			fmt.Println("Возвраст клиента: " + fmt.Sprint(users[index].age))
		}
	}

}
