package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	//var bookings [50]string
	var bookings []string
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available. \n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend \n")
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		fmt.Println("Enter your firt name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)
		// user inputs validation
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
		//isValidCity := city == "Singapore" || city == "London"
		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			//bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName+" "+lastName)
			// printing from the array
			/*
				fmt.Printf("the whole array: %v \n ", bookings)
				fmt.Printf("the first value: %v \n", bookings[0])
				fmt.Printf("Array type: %T \n", bookings)
				fmt.Printf("Array Length: %v \n", len(bookings))
			*/
			fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking) // return slice with the split elements
				firstNames = append(firstNames, names[0])
			}
			// display user names
			fmt.Printf("The first names of bookings: %v\n", firstNames)
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out, come back next year .")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @")
			}

			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			//fmt.Println("Your input data is invalid, try again")
			/*fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets \n", remainingTickets, userTickets)
			continue // this is will pass to next iteration*/
		}

	}

	// switch
	city := "London"
	switch city {
	case "New York":
		// code for new york
	case "Berlin":
		// code for Berlin
	case "value 1", "value2":
		// code
	default:
		// default code
	}

}
