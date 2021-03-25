package main

import ("fmt")

func doble(n int) int { return n * 2 }

func proceso(n int, dobles chan int){
	
	//Procesar
	n = n * n * n

	//Enviar por el canal
	dobles<-n


}

func main(){

	dobles := make(chan int)

	go func() {		
		dobles <- doble(2)
	}()

	go proceso(2,dobles)
	
	fmt.Println("Leyendo desde el main primer valor: ",<-dobles)
	fmt.Println("Leyendo desde el main segundo valor: ",<-dobles)


}