package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const hunger = 3

var philosophers = []string{"Plato", "Socrates", "Aristotle", "Pascal", "Locke"}
var wg sync.WaitGroup
var sleepTime = 1 * time.Second
var eatTime = 2 * time.Second
var thinkTime = 1 * time.Second
var orderFinished = []string{}
var orderMutex sync.Mutex

func diningProblem(p string, leftFork, rightFork *sync.Mutex) {
	defer wg.Done()
	fmt.Println(p, "is seated.")
	// philosopher takes one second to seat.
	time.Sleep(sleepTime)

	// each philosopher is going to eat three times.
	for i := hunger; i > 0; i-- {
		fmt.Println(p, "is hungry.")
		// philosopher is picking up both the forks. it takes one second.
		time.Sleep(sleepTime)
		// philosopher has picked up left fork. left fork is locked. no one can use it, unless it's unlocked
		leftFork.Lock()
		fmt.Printf("\t%s has picked up the fork to his left.\n", p)

		// philosopher has picked up right fork. right fork is locked. no one can use it, unless it's unlocked
		rightFork.Lock()
		fmt.Printf("\t%s has picked up the fork to his right.\n", p)

		fmt.Println(p, "has both forks, and is eating.")
		// philosopher starts eating. it takes him two seconds to eat.
		time.Sleep(eatTime)

		// give the philosopher some time to think
		fmt.Println(p, "is thinking.")
		time.Sleep(thinkTime)

		rightFork.Unlock()
		fmt.Printf("\t%s has put down the fork on his right.\n", p)
		leftFork.Unlock()
		fmt.Printf("\t%s has put down the fork on his left.\n", p)
		// philosopher takes one second to put down both the forks
		time.Sleep(sleepTime)
	}

	// current philosopher is done eating
	fmt.Println(p, "is satisfied.")
	// philosopher is leaving. takes one second
	time.Sleep(sleepTime)
	fmt.Println(p, "has left the table.")
	orderMutex.Lock()
	defer orderMutex.Unlock()
	orderFinished = append(orderFinished, p)
}

func main() {
	fmt.Println("The Dining Philosophers Problem")
	fmt.Println("-------------------------------")

	wg.Add(len(philosophers))

	forkLeft := &sync.Mutex{}

	for i := 0; i < len(philosophers); i++ {
		// mutex for the right fork
		forkRight := &sync.Mutex{}
		go diningProblem(philosophers[i], forkLeft, forkRight)

		forkLeft = forkRight
	}

	wg.Wait()

	fmt.Println("The table is empty.")
	fmt.Println("--------------------------")
	fmt.Printf("Order finished: %s\n", strings.Join(orderFinished, ", "))
}
