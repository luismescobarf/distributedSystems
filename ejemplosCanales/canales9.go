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

func main(){

	dobles := make(chan int)

	go func() {		
		for i:=0;i<5;i++{
			time.Sleep(500 * time.Millisecond)
			dobles <- doble(i)			
		}		
	}()	

	for {
		i := <-dobles
		fmt.Printf("%d ", i)
	}


}