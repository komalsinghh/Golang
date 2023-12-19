package main
 import "fmt"

func main(){
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint= 50
	
	fmt.Printf("Conference tickets is %T, remaining tickets is %T, conferenceName is %T\n",conferenceTickets,remainingTickets,conferenceName)

	fmt.Printf("Welcome to %v application\n",conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n",conferenceTickets,remainingTickets)
	fmt.Println("Get your tickets here to attend")
	// fmt.Printf(&remainingTickets)

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

	//updating tickets when ticket is booked
	remainingTickets=remainingTickets-userTickets
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation mail at %v\n",firstName,lastName,userTickets,email)
	fmt.Printf("Remaining Tickets %v",remainingTickets)
}