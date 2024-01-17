package main

import (
	"fmt"
	"structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (DD/MM/YYYY): ")

	var appUser *user.User
	appUser, err := user.New(userFirstName, userLastName, userBirthDate)
	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@mail.com", "123456")
	admin.OutputUserDetails()

	//  ... do something awesome with that gathered dat!
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
