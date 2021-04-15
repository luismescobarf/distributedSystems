package main

import ("fmt"
		"time"
	)

func doble(n int) int { return n * 2 }

func proceso(n int, dobles chan int){	
	//Procesar
	n = n * n * n
	//Enviar por el canal
	dobles<-n
}

func genera(salida chan<- int){
	for i:=1;i<5;i++{
		salida <- i
	}
	close(salida)
}

func cuadrado(entrada <-chan int, salida chan<- int){
	for i := range entrada {
		time.Sleep(500 * time.Millisecond)
		salida <- i * i
	}
	close(salida)
}

func main(){

	números := make(chan int)
	cuadrados := make(chan int)
	cuartas := make(chan int)
	
	go genera(números)
	go cuadrado(números, cuadrados)
	go cuadrado(cuadrados, cuartas)

	for i:= range cuartas{		
		fmt.Printf("%d ", i)
	}

}