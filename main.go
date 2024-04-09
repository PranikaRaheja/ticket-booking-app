package main

import (
	"fmt"
	"strings"
	"time"
)

	const conferenceTickets int=50
	var conferenceName = "Go Conference"
	var remainingTickets uint=50
	var bookings =make([]UserData, 0)

type UserData struct{
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

func main(){
	
	
	greetUsers()

	
	
	for {

	firstName, lastName,email,userTickets :=getUserInput()
	isValidName,isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, uint(userTickets),remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		
		bookTicket( uint(userTickets), firstName, lastName, email)
		go sendTicket(uint(userTickets), firstName, lastName, email)
		//call function print first names
		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are : %v\n",firstNames)

		if remainingTickets==0{
		//end program
		fmt.Println("Our conferece is booked out, Come back next year.")
		break
		}
	}else{
		if !isValidName {
			fmt.Println("First name or last name you entered is too short ")
		}
		if !isValidEmail{
			fmt.Println("Email address you entered doesn't contain @ sign")

		}
		if !isValidTicketNumber{
			fmt.Println("Number of tickets you entered is invalid")
		}
		fmt.Println("Your input data is invalid, try again")
		
	}

	}
}
func greetUsers(){
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n",conferenceTickets,remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
func getFirstNames()[] string{
	firstNames := []string{}
		for _, booking := range bookings{
		
		firstNames=append(firstNames, booking.firstName)
		}
		return firstNames
}


func getUserInput()(string, string, string, int){
	var firstName string
	var lastName string
	var email string
	var userTickets int
	// ask user for their name
	fmt.Println("Enter your First name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your Last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint,remainingTickets uint)(bool, bool ,bool){
	isValidName :=len(firstName)>=2 && len(lastName) >=2
	isValidEmail := strings.Contains(email,"@")
	isValidTicketNumber := userTickets >0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}


func bookTicket( userTickets uint, firstName string, lastName string, email string){
		remainingTickets=remainingTickets-uint(userTickets)
		
		//create a map for a user
		var userData = UserData{
			firstName:firstName,
			lastName: lastName,
			email: email,
			numberOfTickets: userTickets,
		}


		bookings=append(bookings, userData)
		fmt.Printf("List of bookings is %v\n",bookings)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",firstName,lastName,userTickets,email)
		fmt.Printf("%v tickets remaining for %v\n",remainingTickets,conferenceName)
}

func sendTicket(userTickets uint,firstName string, lastName string,email string){
	time.Sleep(10*time.Second)
	var ticket=fmt.Sprintf("%v tickets for %v %v",userTickets,firstName, lastName)
	fmt.Println("###########")
	fmt.Printf("Sending ticket:\n %v to email address %v\n",ticket,email)
	fmt.Println("###########")
}