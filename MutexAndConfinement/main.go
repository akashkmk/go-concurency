package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func purchaseTicket(wg *sync.WaitGroup, userId int, remainingTicket *int) {
	defer wg.Done()

	mutex.Lock()
	if *remainingTicket > 0 {
		*remainingTicket-- // User purchase a ticket
		fmt.Printf("User %d purchase a ticket and ramaining %d\n", userId, *remainingTicket)
	} else {
		fmt.Printf("User %d not found any ticket\n", userId)
	}
	mutex.Unlock()
}

func main() {
	var TotalTicket int = 500

	var wg sync.WaitGroup

	for userId := 0; userId < 2000; userId++ {
		wg.Add(1)
		go purchaseTicket(&wg, userId, &TotalTicket)
	}

	wg.Wait()
}
