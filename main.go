//concorente
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
	// filosofos demoram 1seg para sentar.
	time.Sleep(sleepTime)

	// cada filosofo vai comer 3x.
	for i := hunger; i > 0; i-- {
		fmt.Println(p, "esta com fome.")
		// esta pegando os garfos, demorando 1seg
		time.Sleep(sleepTime)
		
		// filosofo pegou o garfo esquerdo. garfo esquerdo travado. ninguem mais usa, ate ser desbloqueado.
		leftFork.Lock()
		fmt.Printf("\t%s pegou o garfo a sua esquerda.\n", p)

		// filosofo pegou o garfo direito. garfo direito travado. ninguem mais usa, ate ser desbloqueado.
		rightFork.Lock()
		fmt.Printf("\t%s pegou o garfo a sua direita.\n", p)

		fmt.Println(p, "esta com os dois garfos e comendo.")
		// filosofo comeca a comer, 2seg para comer.
		time.Sleep(eatTime)

		// dar temo para o filosofo pensar
		fmt.Println(p, "esta pensando.")
		time.Sleep(thinkTime)

		rightFork.Unlock()
		fmt.Printf("\t%s devolveu o garfo direito.\n", p)
		leftFork.Unlock()
		fmt.Printf("\t%s devolveu o garfo esquerdo.\n", p)
		// filosofo demora 1seg para botar os garfos na mesa 
		time.Sleep(sleepTime)
	}

	// filosofo atual comeu
	fmt.Println(p, "esta satisfeito.")
	// filosofo saindo, demora 1seg
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
		// mutex para o garfo direito
		forkRight := &sync.Mutex{}
		go diningProblem(philosophers[i], forkLeft, forkRight)

		forkLeft = forkRight

		// dormir 5seg para o proximo filosofo
		time.Sleep(5 * time.Second)
	}

	wg.Wait()

	fmt.Println("Mesa vazia.")
	fmt.Println("--------------------------")
	fmt.Printf("Ordem de encerramento: %s\n", strings.Join(orderFinished, ", "))
}
