package main

import (
	"fmt"
	"strings"
	"time"
)

const hunger = 3

var philosophers = []string{"Platão", "Socrates", "Aristoteles", "Pascal", "Locke"}
var orderFinished = []string{}

func diningProblem(p string) {
	fmt.Println(p, "está sentado.")
	// O filósofo leva um segundo para se sentar.
	time.Sleep(1 * time.Second)

	// Cada filósofo vai comer três vezes.
	for i := hunger; i > 0; i-- {
		fmt.Println(p, "está com fome.")
		// O filósofo pega ambos os garfos. Isso leva um segundo.
		time.Sleep(1 * time.Second)

		fmt.Printf("\t%s pegou o garfo à sua esquerda.\n", p)
		fmt.Printf("\t%s pegou o garfo à sua direita.\n", p)

		fmt.Println(p, "está com os dois garfos e comendo.")
		// O filósofo começa a comer. Isso leva dois segundos.
		time.Sleep(2 * time.Second)

		fmt.Println(p, "está pensando.")
		// Dê ao filósofo um tempo para pensar.
		time.Sleep(1 * time.Second)

		fmt.Printf("\t%s devolveu o garfo direito.\n", p)
		fmt.Printf("\t%s devolveu o garfo esquerdo.\n", p)
		// O filósofo leva um segundo para colocar ambos os garfos de volta.
		time.Sleep(1 * time.Second)
	}

	// O filósofo atual terminou de comer.
	fmt.Println(p, "está satisfeito.")
	// O filósofo está saindo. Leva um segundo.
	time.Sleep(1 * time.Second)
	fmt.Println(p, "saiu da mesa.")
	orderFinished = append(orderFinished, p)
}

func main() {
	fmt.Println("O problema do jantar dos filósofos")
	fmt.Println("-------------------------------")

	for _, philosopher := range philosophers {
		diningProblem(philosopher)
	}

	fmt.Println("Mesa vazia.")
	fmt.Println("--------------------------")
	fmt.Printf("Ordem de encerramento: %s\n", strings.Join(orderFinished, ", "))
}
