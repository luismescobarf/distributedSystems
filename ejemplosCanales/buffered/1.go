package main

//Paquetes requeridos (utilizados)
import (
	"fmt"
	"time"
)

func doble(n int) int {
	time.Sleep(500 * time.Millisecond)
	return n * 2
}

func triple(n int) int {
	time.Sleep(800 * time.Millisecond)
	return n * 3
}

//Envío de canal
func main() {

	//Creación de canales
	dobles := make(chan int)
	triples := make(chan int)

	//Canal de cancelación
	var cancela chan struct{}

	go func() {
		time.Sleep(3 * time.Second)
		cancela = make(chan struct{})
		cancela <- struct{}{}
	}()

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
		case <-cancela:
			fmt.Println("cancelado!")
			break loop
		case i, ok := <-dobles:
			if !ok {
				break loop
			}
			fmt.Println(i, " ")
		case i, ok := <-triples:
			if !ok {
				break loop
			}
			fmt.Println(i, " ")
			//default:
			//Instrucciones por defecto
			//time.Sleep(250*time.Millisecond);
			//fmt.Println("Por defecto mientras se completan los otros procesos!");
		}
		fmt.Println()
	}

} //Final de la sección principal
