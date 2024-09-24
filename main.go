package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const hunger = 3

var philosophers = []string{"Platão", "Socrates", "Aristoteles", "Pascal", "Locke"}
var wg sync.WaitGroup
var sleepTime = 1 * time.Second
var eatTime = 2 * time.Second
var thinkTime = 1 * time.Second
var orderFinished = []string{}
var orderMutex sync.Mutex

func diningProblem(p string, leftFork, rightFork *sync.Mutex) {
	defer wg.Done()
	fmt.Println(p, "esta sentado.")
	// philosopher takes one second to seat.
	time.Sleep(sleepTime)

	// each philosopher is going to eat three times.
	for i := hunger; i > 0; i-- {
		fmt.Println(p, "esta com fome.")
		// philosopher is picking up both the forks. it takes one second.
		time.Sleep(sleepTime)
		// philosopher has picked up left fork. left fork is locked. no one can use it, unless it's unlocked
		leftFork.Lock()
		fmt.Printf("\t%s pegou o garfo a sua esquerda.\n", p)

		// philosopher has picked up right fork. right fork is locked. no one can use it, unless it's unlocked
		rightFork.Lock()
		fmt.Printf("\t%s pegou o garfo a sua direita.\n", p)

		fmt.Println(p, "esta com os dois garfos e comendo.")
		// philosopher starts eating. it takes him two seconds to eat.
		time.Sleep(eatTime)

		// give the philosopher some time to think
		fmt.Println(p, "esta pensando.")
		time.Sleep(thinkTime)

		rightFork.Unlock()
		fmt.Printf("\t%s devolveu o garfo direito.\n", p)
		leftFork.Unlock()
		fmt.Printf("\t%s devolveu o garfo esquerdo.\n", p)
		// philosopher takes one second to put down both the forks
		time.Sleep(sleepTime)
	}

	// current philosopher is done eating
	fmt.Println(p, "esta satisfeito.")
	// philosopher is leaving. takes one second
	time.Sleep(sleepTime)
	fmt.Println(p, "saiu da mesa.")
	orderMutex.Lock()
	defer orderMutex.Unlock()
	orderFinished = append(orderFinished, p)
}

func main() {
	fmt.Println("O problema do jantar dos filosofos")
	fmt.Println("-------------------------------")

	wg.Add(len(philosophers))

	forkLeft := &sync.Mutex{}

	for i := 0; i < len(philosophers); i++ {
		// mutex for the right fork
		forkRight := &sync.Mutex{}
		go diningProblem(philosophers[i], forkLeft, forkRight)

		forkLeft = forkRight

		// Sleep for 5 seconds before starting the next philosopher
		time.Sleep(5 * time.Second)
	}

	wg.Wait()

	fmt.Println("Mesa vazia.")
	fmt.Println("--------------------------")
	fmt.Printf("Ordem de encerramento: %s\n", strings.Join(orderFinished, ", "))
}
