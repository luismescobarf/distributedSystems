//Generar  el  conjunto  resultante del  producto  cartesiano  de  dos  conjuntosingresados.Los conjuntos estarán representados en vectores(slices o arrays).
package main

import (
	"fmt"
	"sync"
	"time"
)

//Mecanismo para contar el número de goroutines creadas y esperar
var wg sync.WaitGroup

var vector1 [5]int
var vector2 [5]int

var vector1Concurrente [5]int
var vector2Concurrente [5]int

var vector3 []int

func llenado() {
	for i := 0; i < 5; i++ {
		vector1[i] = i + 1
		vector2[i] = i + 2
	}
	fmt.Println(vector1, vector2)
	time.Sleep(3 * time.Second)
}

func productoC() {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		for j := 0; j < 5; j++ {
			aux := (vector1[i]) * (vector2[j])
			vector3 = append(vector3, aux)
		}
	}
	fmt.Println(vector3)
}

func main() {
	inicio := time.Now()
	llenado()
	productoC()
	fin := time.Now()
	fmt.Println(fin.Sub(inicio))
}
