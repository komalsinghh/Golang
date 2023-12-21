package main

import (
	"booking-app/common"
	"fmt"
	"strings"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50

var bookings []string

func main() {

	greetUsers()
	fmt.Printf("Conference tickets is %T, remaining tickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	// fmt.Printf(&remainingTickets)

	//var bookings [50]string

	for {

		firstName, lastName, email, userTickets := getUserInput()

		//isValidCity := city == "Pune" || "Goa"

		isValidName, isValidEmail, isValidTicketNumber := common.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			//fmt.Printf("The Whole array: %v\n", bookings)
			fmt.Printf("The whole slice: %v\n", bookings)
			fmt.Printf("The first value: %v\n", bookings[0])
			fmt.Printf("Array type: %T\n", bookings)
			fmt.Printf("Array lenth: %v\n", len(bookings))

			fmt.Printf("All the tickets are: %v\n", bookings)

			//call a function to print first name
			firstNames := getFirstName()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			//var noTicketsRemaining bool= remainingTickets==0
			// noTicketsRemaining := remainingTickets==0
			// 	if noTicketsRemaining{
			if remainingTickets == 0 {
				//end program
				fmt.Println("Our conference is booked out. Come back next year!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Name is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Invalid Ticket Number")
			}
		}
	}

	// city := "London"

	// switch city {
	// case "New York":
	// 	//execute code for booking New York conference tickets
	// case "Singapore" , "Hong Kong":
	// 	//execute code for booking Singapore conference tickets
	// case "London", "Berlin":
	// 	//execute code for booking London conference tickets
	// case "Mexico City":
	// 	//execute code for booking New York conference tickets
	// default:
	// 	fmt.Println("No valid city selected")
	// }
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func getFirstName() []string {
	// print only the first name of user who booked list
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	// updating tickets when ticket is booked
	remainingTickets = remainingTickets - userTickets
	// bookings[0]=firstName+" "+lastName

	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation mail at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("Remaining Tickets %v\n", remainingTickets)

}
