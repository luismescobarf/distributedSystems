package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//Mecanismo para contar el número de goroutines creadas y esperar
	var wg sync.WaitGroup

	operación := func(i int) {
		defer wg.Done() //Restar uno al WaitGroup, el defer garantiza la realización al final de la función
		fmt.Println("operación", i, time.Now().Format("5"))
		time.Sleep(3 * time.Second) //Simular carga computacional
		fmt.Println("operación", i, time.Now().Format("5"))
	}

	inicio := time.Now() //Inicio de la toma de tiempo

	for i := 0; i < 3; i++ {
		wg.Add(1) //Añadir un goroutine al conteo justo antes de que se cree
		go operación(i)
	}

	//Especificar que el gouroutine main espere a que la cantidad de goroutines sean 0
	wg.Wait()

	//Finalización y presentación de la toma de tiempo
	fmt.Println("tiempo total:", time.Since(inicio))

}
