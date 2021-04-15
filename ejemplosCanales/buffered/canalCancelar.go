package main

import (
	"fmt"
	"time"
)

//func doble(n int) int { return n * 2 }

func proceso(n int, dobles chan int) {
	//Procesar
	n = n * n * n
	//Enviar por el canal
	dobles <- n
}

func genera(salida chan<- int) {
	for i := 1; i < 5; i++ {
		salida <- i
	}
	close(salida)
}

func cuadrado(entrada <-chan int, salida chan<- int) {
	for i := range entrada {
		time.Sleep(500 * time.Millisecond)
		salida <- i * i
	}
	close(salida)
}

func doble(n int) int {
	time.Sleep(500 * time.Millisecond)
	return n * 2
}

func triple(n int) int {
	time.Sleep(800 * time.Millisecond)
	return n * 3
}

func main() {

	dobles := make(chan int)
	triples := make(chan int)
	var cancela chan struct{} //Canal de eventos

	go func() {
		for i := 0; i < 10; i++ {
			dobles <- doble(i)
			triples <- triple(i)
		}
		close(dobles)
		close(triples)

	}()

loop:
	for {
		select {
		case i, ok := <-dobles:
			if !ok {
				break loop
			}
			fmt.Print(i, " ")

		case i, ok := <-triples:
			if !ok {
				break loop
			}
			fmt.Print(i, " ")

		default:
			//Código...

		}
	}

}
