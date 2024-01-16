package main

import "fmt"

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthdate (DD/MM/YYYY): ")

	//  ... do something awesome with that gathered dat!
	fmt.Println(firstName, lastName, birthDate)
}

func outputUserDetails() {

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
