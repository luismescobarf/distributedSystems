package main

//Paquetes requeridos (utilizados)
import (
    "fmt"
    "time"  
)

//Counting semaphore
func main() {
	
	//sema := make(chan struct{},20);
	canal := make(chan int);
	
	for i:=0; i<10000; i++ {//Creación de 10000 Goroutines
		
		go func(i int){
			
			//sema <- struct{}{} //Adquiere un token
			time.Sleep(1*time.Second);
			canal <- i
			//<-sema //Devuelve el token
			
		}(i)
		
	}
	
	for i:=0; i<10000; i++ {
		
		fmt.Printf("%d ", <-canal)
		
	}
	fmt.Println();
		
	
}//Final de la sección principal
