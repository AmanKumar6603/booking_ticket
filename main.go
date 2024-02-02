package main

import (
	"booking_ticket/common"
	"fmt"
	"time"
)

// package level variables
var trainName = "Rajdhani Express"

const trainTickets uint = 50

var remainingTickets uint = 50
var booking = make([]Userdata, 0)

type Userdata struct {
	firstName       string
	Lastname        string
	email           string
	numberOfTickets uint
}

// main function
func main() {
	for remainingTickets > 0 {
		greetUsers() // greeting function
		// get user input functions
		firstname, Lastname, userTickets, useremail := getUserinput()
		//validate user inputs
		isValidName, isValidEmail, isvalidticketNo := common.ValidateInputParameters(firstname, Lastname, useremail, userTickets, remainingTickets)
		if isValidName && isValidEmail && isvalidticketNo {
			bookingTickets(firstname, Lastname, userTickets, useremail)
			go sendTicket(firstname, Lastname, userTickets, useremail)
			firstNames := getFirstnames()
			fmt.Printf("The first names of bookings are %v\n", firstNames)
			if remainingTickets == 0 {
				//end the loop
				fmt.Println("All tickets are booked, Come next day...😊👍⏭️")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First or last name you enter is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you enter doesn't contain @ sign")
			}
			if !isvalidticketNo {
				fmt.Println("Number of tickets you enter is invalid")
			}
			fmt.Println("Try again....!")
		}
	}
}

// greeting functions
func greetUsers() {
	fmt.Println("(❁´◡`❁)(❁´◡`❁)(❁´◡`❁)(❁´◡`❁)(❁´◡`❁)(❁´◡`❁)(❁´◡`❁)(❁´◡`❁)(❁´◡`❁)(❁´◡`❁)")
	fmt.Println("Welcome to our Train 🚆 tickets booking App...!")
	fmt.Printf("The train name is %v ", trainName)
	fmt.Printf("we have total of %v tickets and %v are still available\n", trainTickets, remainingTickets)
	fmt.Println("Get your ticket to get into the", trainName)
}

// get first name
func getFirstnames() []string {
	firstNames := []string{}
	for _, bookings := range booking {
		firstNames = append(firstNames, bookings.firstName)
	}
	return firstNames
}

// get User input
func getUserinput() (string, string, uint, string) {
	var firstname string
	var Lastname string
	var userTickets uint
	var useremail string
	fmt.Println("Enter the first name ")
	fmt.Scan(&firstname)
	fmt.Println("Enter the last name ")
	fmt.Scan(&Lastname)
	fmt.Println("Enter the number of tickets want to book")
	fmt.Scan(&userTickets)
	fmt.Println("Enter your email")
	fmt.Scan(&useremail)
	return firstname, Lastname, userTickets, useremail
}

// booking tickets
func bookingTickets(firstname string, Lastname string, userTickets uint, useremail string) {
	remainingTickets = remainingTickets - userTickets
	var userData = Userdata{
		firstName:       firstname,
		Lastname:        Lastname,
		email:           useremail,
		numberOfTickets: userTickets,
	}
	booking = append(booking, userData)
	fmt.Printf("List of bookings is %v\n", booking)
	fmt.Printf("%v buy %v tickets\n", firstname, userTickets)
	fmt.Printf("Thankyou %v %v for booking %v tickets....😊😊😊\n", firstname, Lastname, userTickets)
	fmt.Printf("Remaining tickets are %v\n", remainingTickets)
}

// sending tickets function
func sendTicket(firstName string, lastname string, userTickets uint, email string) {
	time.Sleep(10 * time.Second)
	var tickets = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastname)
	fmt.Println("😊😊😊😊TICKET BOOKED😊😊😊😊😊😊😊😊😊")
	fmt.Printf("Sending %v to email address %v\n", tickets, email)
	fmt.Println("📩📩📩📩📩 SENT SENT SENT ...📩📩📩📩📩📩")
}
