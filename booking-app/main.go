package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available. \n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend \n")
	var firstName string
	var lastName string
	var email string
	var userTickets int
	//fmt.Print(remainingTickets) // this will displays the value of the variable .
	//fmt.Print(conferenceName)   // this will displays  the address memory of the variable.
	//var userTickets int
	/*fmt.Printf("conferenceTickets is %T, remainingTickets is %T , conferenceName is %T \n", conferenceTickets, remainingTickets, conferenceName) // %T is placeholder for the type of the variable
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available. \n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend")
	var userName string
	var userTickets int
	// ask the user for their name
	userName = "Jamil"
	userTickets = 2
	fmt.Printf("User %v bocked %v tickets.\n", userName, userTickets)*/
	/*fmt.Println("userName value: ",userName)
	fmt.Println("userName address: ",&userName)*/
	// User Input
	fmt.Println("Enter your firt name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)
	fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
}
