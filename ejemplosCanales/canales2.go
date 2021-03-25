package main

//import ("fmt")

func doble(n int) int { return n * 2 }

func main(){
	go doble(2)
}