package main

import (
	"fmt"
	"time"
)

func main() {

	operación := func(i int) {
		fmt.Println("operación", i, time.Now().Format("5"))
		time.Sleep(3 * time.Second) //Simular carga computacional
		fmt.Println("operación", i, time.Now().Format("5"))
	}

	inicio := time.Now() //Inicio de la toma de tiempo

	for i := 0; i < 3; i++ {
		operación(i)
	}

	//Finalización y presentación de la toma de tiempo
	fmt.Println("tiempo total:", time.Since(inicio))

}
