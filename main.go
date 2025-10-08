package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const totalTickets = 100
var conferenceName = "International Conference on Go"
var remainingTickets uint = 100
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	noOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUser()

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicket := isValidUser(firstName, lastName, email, userTickets)

	if isValidName && isValidEmail && isValidTicket {
		bookTickets(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTickets(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("First names of bookings are: %v\n", firstNames)

		if remainingTickets==0 {
			fmt.Println("All tickets are booked out...")
		}
	} else {
		if !isValidName {
			fmt.Println("First or last name is too short")
		}
		if !isValidEmail {
			fmt.Println("Invalid email address")
		}
		if !isValidTicket {
			fmt.Println("Number of tickets are invalid")
		}
	}
	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to %v ticket booking application\n", conferenceName)
	fmt.Printf("We have total %v tickets and %v tickets remaining\n", totalTickets, remainingTickets)
	fmt.Println("Book your ticket now to attend conference")
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string 
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets you want to book: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func isValidUser(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName)>=2 && len(lastName)>=2 
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := userTickets>0 && userTickets<=remainingTickets

	return isValidName, isValidEmail, isValidTicket
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets-userTickets

	var user = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		noOfTickets: userTickets,
	}

	bookings = append(bookings, user)
	fmt.Printf("List of booking is %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking ticket ...you will get confirmational email at %v\n", firstName, lastName, email)
	fmt.Printf("%v tickets still remaining\n", remainingTickets)
}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(30 * time.Second)

	var ticket = fmt.Sprintf("%v tickets is sent for %v %v\n", userTickets, firstName, lastName);
	fmt.Println("-------------------")
	fmt.Printf("Sending ticket:\n %v \n to email %v\n", ticket, email)
	fmt.Println("-------------------")

	wg.Done()
}
