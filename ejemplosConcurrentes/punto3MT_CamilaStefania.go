//Generar  el  conjunto  resultante del  producto  cartesiano  de  dos  conjuntosingresados.Los conjuntos estarán representados en vectores(slices o arrays).
package main

import (
	"fmt"
	"sync"
	"time"
)

type fragmento struct {
	n   int
	vec []int
}

//Mecanismo para contar el número de goroutines creadas y esperar
var wg sync.WaitGroup

func llenado(salida chan []int) {

	var vector2 []int
	var vector1 []int

	var tamanio int = 5

	for i := 1; i <= tamanio; i++ {
		vector2 = append(vector2, i+2)
		vector1 = append(vector1, i+1)
	}
	time.Sleep(3 * time.Second)

	salida <- vector2
	salida <- vector1

	close(salida)
}

func productoFragmento(canalFragmento chan fragmento, canalFragmentoSalida chan fragmento) {

	defer wg.Done()

	var resultadoProducto []int

	fragmentoRecibido := <-canalFragmento

	for i := 0; i < len(fragmentoRecibido.vec); i++ {
		resultadoProducto = append(resultadoProducto, fragmentoRecibido.n*fragmentoRecibido.vec[i])
	}

	fmt.Println("Para enviar -> ", resultadoProducto)

	canalFragmentoSalida <- fragmento{n: fragmentoRecibido.n, vec: resultadoProducto}

	time.Sleep(500 * time.Millisecond)

}

func main() {

	inicio := time.Now()
	lleno := make(chan []int, 2) //Memoria!!!!

	go llenado(lleno)

	sliceGenerado := <-lleno
	sliceGenerado2 := <-lleno

	fmt.Println("Recibimos 2 vectores")
	fmt.Println(sliceGenerado)
	fmt.Println(sliceGenerado2)

	var productoC [][]int

	canalFragmento := make(chan fragmento, len(sliceGenerado2))
	canalFragmentoSalida := make(chan fragmento, len(sliceGenerado2))

	for i := 0; i < len(sliceGenerado); i++ {
		canalFragmento <- fragmento{n: sliceGenerado[i], vec: sliceGenerado2}
	}

	for i := 0; i < len(sliceGenerado); i++ {
		wg.Add(1)
		go productoFragmento(canalFragmento, canalFragmentoSalida)
		// for j := 0; j <= 4; j++ {
		// 	productoC[i][j] = sliceGenerado[i] * sliceGenerado2[j]
		// }
	}
	wg.Wait()

	for i := 0; i < len(sliceGenerado); i++ {
		e := <-canalFragmentoSalida
		productoC = append(productoC, e.vec)

	}

	fmt.Println(productoC)

	fin := time.Now()
	fmt.Println(fin.Sub(inicio))
}
