package main

import "fmt"

func main(){
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
  var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n",conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.",conferenceTickets,remainingTickets)
	fmt.Println("Get your tickets here ")

	var bookings []string
	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint
		
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		
		// bookings[0] = firstName+" "+lastName
		bookings = append(bookings,firstName+" "+lastName)
	
		fmt.Println("Enter your email: ")
		fmt.Scan(&email)
	
		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)
	
		remainingTickets = remainingTickets - userTickets
		fmt.Printf("User %v booked %v ticket.\n",bookings[0],userTickets)
		fmt.Printf("%v tickets remaining",remainingTickets)
		fmt.Printf("There are all our bookings: %v\n",bookings)
	}
	
}