package main

import (
	"fmt"
	"strings"
)

func main() {

	concert := "music concert"
	const Totaltickets = 45
	var remainingtickets uint = 45

	bookings := []string{}

	fmt.Println("Welcome to our", concert, "ticket booking website")
	fmt.Println("Our", concert, "has a total of", Totaltickets, "tickets available and remaininng tickets are", remainingtickets)
	fmt.Println("Please book your tickets for the event")

	for {

		var firstname string
		var lastname string
		var email string
		var tickets uint
		var mobileno int
		// ask user for their name and no. of tickets needed
		fmt.Println("Enter your first name : ")
		fmt.Scan(&firstname)
		fmt.Println("Enter your last name :")
		fmt.Scan(&lastname)

		fmt.Println("Enter your mobile number :")
		fmt.Scan(&mobileno)

		fmt.Println("Enter your email address :")

		fmt.Scan(&email)
		fmt.Println("Number of tickets")

		fmt.Scan(&tickets)

		if tickets <= remainingtickets {

			remainingtickets = remainingtickets - tickets

			bookings = append(bookings, firstname+" "+lastname)

			fmt.Printf("%v %v booked %v number of tickets. You will get an confirmation email to %v along with the booking ID to your mobile numnber %v.\n", firstname, lastname, tickets, email, mobileno)

			fmt.Printf("Thanks for booking the tickets for our %v.\n", concert)
			fmt.Printf("Remaining tickets for %v are %v.\n", concert, remainingtickets)
			fmt.Printf("First value in bookings is %v.\n", bookings[0])
			Firstnames := []string{}
			for _, element := range bookings {
				var name = strings.Fields(element)
				Firstnames = append(Firstnames, name[0])
			}
			fmt.Printf("the first names of the bookings are : %v.\n", Firstnames)
			fmt.Printf("the contents of the bookings are : %v.\n", bookings)

			if remainingtickets == 0 {
				fmt.Printf("The tickets for our %v are booked out. Pleasecomeback Next year.\n", concert)
				break

			}

		} else {
			fmt.Printf("There are only %v tickets available for our %v. You cannot book %v tickets. please try again with new number of tickets that are less than %v tickets.\n", remainingtickets, concert, tickets, remainingtickets)

		}

	}

}
