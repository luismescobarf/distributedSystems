package main

//Paquetes requeridos (utilizados)
import (
	"fmt"
	"time"
)

//Buffered channels

//Envío de canal
func main() {

	canal := make(chan int, 1)
	fmt.Println(len(canal), cap(canal))

	go func() {
		canal <- 1
		fmt.Println("Ya envié!")
		fmt.Println(len(canal), cap(canal))
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("Recibí ", <-canal)
	fmt.Println(len(canal), cap(canal))
	//var pausa string;
	//fmt.Scan(&pausa);

} //Final de la sección principal
