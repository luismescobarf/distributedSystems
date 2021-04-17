//Generar  el  conjunto  resultante del  producto  cartesiano  de  dos  conjuntosingresados.Los conjuntos estarán representados en vectores(slices o arrays).

package main

import (
	"fmt"

	"sync"

	"time"
)

var vector1 [5]int

var vector2 [5]int

var vector3 []int

var wg sync.WaitGroup

func llenado() {

	for i := 0; i < 5; i++ {

		vector1[i] = i + 1

		vector2[i] = i + 2

	}

	fmt.Println(vector1, vector2)

}

func productoC() {

	for i := 0; i < 5; i++ {

		for j := 0; j < 5; j++ {

			aux := (vector1[i]) * (vector2[j])

			vector3 = append(vector3, aux)

		}

	}

	fmt.Println(vector3)

	wg.Done()

}

func main() {

	inicio := time.Now()

	wg.Add(1)

	llenado()

	go productoC()

	wg.Wait()

	fin := time.Now()

	fmt.Println(fin.Sub(inicio))

}
