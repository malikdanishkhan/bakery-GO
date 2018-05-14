package main

import	"os"
import "fmt"
import	"time"
import	"strconv"
import	"math/rand"

func main() {
	start := time.Now()

	//get arguments
	arguments := os.Args[1:]

	numberServers, ns := strconv.Atoi(arguments[0])
	numberCustomers, nc := strconv.Atoi(arguments[1])

	//check if arguments are valid
	if (ns != nil) || (nc != nil) {
		fmt.Println("Arguments must be integers")
		os.Exit(2)
	}

	//start server channel
	serverChannel := make(chan Customer, numberCustomers)

	//start ticket channel
	ticketChannel := make(chan int, numberCustomers) 

	//make customers
	for i := 1; i <= numberCustomers; i++ {
		go open(i, serverChannel, ticketChannel)
		fmt.Println("Creating customer #", i)
	}

	//make servers
	for i := 1; i <= numberServers; i++ {
		go manage(serverChannel)
		fmt.Println("Creating server #", i)
	}

	fmt.Println("*-----------------------*")
	fmt.Println("Done creating customers and servers")

	//process tickets
	serve(numberCustomers, ticketChannel)
  
	totalTime := time.Since(start)
	fmt.Println("Total run-time ", totalTime)
}

func fib(n int) int {
    if n == 0 {
        return 0
    } else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

type Customer struct {
	ticket int
	value int
}

func open(customerNumber int, serverChannel chan Customer, ticket chan int) {
	arrivalTime := (time.Duration(rand.Intn(7000)) * time.Millisecond)

	//put customer to sleep
	time.Sleep(arrivalTime)

	fmt.Println("*-----------------------*") 
	fmt.Println("A customer just arrived!")

	//place order for customer
	currentOrder := Customer{ticket: customerNumber, value: rand.Intn(40)} 

	fmt.Println("Customer with ticket #", currentOrder.ticket,"placed an order for Fib of:", currentOrder.value)
    
    //send current customer order to server
	serverChannel <- currentOrder 

	//send customer ticket number
	ticket <- customerNumber 
}

func manage(serverChannel chan Customer) {
	for {
		// server is listening for order
		server := <-serverChannel
		// server gives order to customer
		fmt.Println("Served customer ticket #", server.ticket, "with Fib:", fib(server.value)) 
	}
}

func serve(n int, t chan int){
	for i := 0; i < n; i++ {
		x := <- t
		fmt.Println("Processing customer with ticket number ", x)
		fmt.Println("*-----------------------*")
	}
}
