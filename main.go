package main

import (
	"bookingApp/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go conference"

// conferenceName := "Go conference"
const conferenceTicket = 50

var remainingTicket uint = 50
var bookingName = make([]UserData, 0) // slice

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	// fmt.Print("Hello, World!")

	greetUser(conferenceName, conferenceTicket, remainingTicket)

	//for {
	// var firstName string
	// var lastName string
	// var email string
	// var tickets uint
	// var bookingName [50]string  // array

	firstName, lastName, email, tickets := getUserInput()

	// bookingName[0] = firstName + " " + lastName // array adds

	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, tickets, remainingTicket)

	if isValidEmail && isValidName && isValidTicketNumber {
		// var userData = UserData{
		// 	firstName:       firstName,
		// 	lastName:        lastName,
		// 	email:           email,
		// 	NumberOfTickets: tickets,
		// }
		// // var userData = make(map[string]string)
		// // userData["fistName"] = firstName
		// // userData["lastData"] = lastName
		// // userData["email"] = email
		// // userData["numberOfTicket"] = strconv.FormatUint(uint64(tickets), 10)
		// bookingName = append(bookingName, userData)

		// fmt.Printf("Thank you %v for booking %v ticket. You will get your ticket at your email %v\n", bookingName[0], tickets, email)
		// remainingTicket = remainingTicket - tickets
		// fmt.Println("Toatal: ", remainingTicket, "left")

		// firstNames := printFirstName()

		bookTicket(tickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(tickets, firstName, lastName, email)

		firstNames := printFirstName()
		fmt.Println("The first name of the booking are: ", firstNames)
		if remainingTicket == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
			//break
		}
	} else {
		//fmt.Printf("We have only %v ticket left and your are trying to book %v tickets", remainigTicket, tickets)
		if !isValidName {
			fmt.Println("first name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("number of tickets you entered is invalid")
		}

	}
	//}
	wg.Wait()
}

func greetUser(conferenceName string, conferenceTicket int, remainingTicket uint) {
	fmt.Printf("Welcome to our  %v booking App\n", conferenceName)
	fmt.Printf("We have total of %v ticket and %v tickets are still available\n", conferenceTicket, remainingTicket)
	fmt.Println("Get your ticket here to attend")

}

func printFirstName() []string {
	firstNames := []string{}
	for _, name := range bookingName {
		firstNames = append(firstNames, name.firstName)
	}
	// fmt.Printf("These are all ours booking: %v\n", firstNames)
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var tickets uint
	fmt.Println("Enter your firstName:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your lastName:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email:")
	fmt.Scan(&email)
	fmt.Println("Enter the number of tickets:")
	fmt.Scan(&tickets)
	return firstName, lastName, email, tickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTicket = remainingTicket - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookingName = append(bookingName, userData)
	fmt.Printf("List of bookings is %v\n", bookingName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTicket, conferenceName)
}

func sendTicket(numberOfTicket uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", numberOfTicket, firstName, lastName)
	fmt.Println("###################################")
	fmt.Printf("Sending ticket: %v\n to to email address %v\n", ticket, email)
	fmt.Println("###################################")
	wg.Done()
}
