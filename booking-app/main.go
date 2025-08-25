package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 5

var (
	conferenceName        = "Go Conference"
	remainingTickets uint = conferenceTickets
	bookings              = make([]UserData, 0)
	wg                    = sync.WaitGroup{}
)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	greetUser()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidUserTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidUserTickets {
			bookTicket(firstName, lastName, userTickets, email)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email) // 'go' creates a lightweight thread

			printFirstNames()

			noTicketsRemaining := remainingTickets == 0
			fmt.Printf("All bookings: %v\n", bookings)
			if noTicketsRemaining {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First/Last name is invalid, too short")
			}
			if !isValidEmail {
				fmt.Println("Email format is invalid")
			}
			if !isValidUserTickets {
				fmt.Println("Number of tickets entered is invalid")
			}
		}
	}
	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
}

func printFirstNames() {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	fmt.Printf("The first names of bookings are: %v\n", firstNames)
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

	fmt.Println("Enter your email name:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, userTickets uint, email string) {
	remainingTickets = remainingTickets - userTickets

	userData := UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("############################")
	fmt.Printf("Sending ticket: %v \nto email address %v\n", ticket, email)
	fmt.Println("############################")

	wg.Done()
}
