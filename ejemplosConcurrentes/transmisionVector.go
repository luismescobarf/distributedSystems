//Ejemplo de transmisión de vectores por canales
package main

import (
	"fmt"
)

func main() {

	//Declaramos el canal de comunicación, por donde se va a transmitir el slice
	canal := make(chan []int)

	//Goroutine que genera el contenido
	go func() {

		//Declaramos el slice que vamos a transmitir
		var contenedorDatos []int

		//Límite superior de la secuencia
		var limSuperior int = 10

		//Llenarlo con números pares
		for i := 1; i <= limSuperior; i++ {
			if i%2 == 0 {
				contenedorDatos = append(contenedorDatos, i)
			}

		}

		//Transmitir el slice lleno
		canal <- contenedorDatos

		//Cerrar el canal después de transmitir
		close(canal)

	}()

	//Recibir el vector
	sliceGenerado := <-canal

	//Mostrar en pantalla lo que se recibió
	fmt.Println("Vector generado en una goroutine externa:")
	fmt.Println(sliceGenerado)

	//Realizar una transformación (volver todos impares)
	for posicion, numero := range sliceGenerado {
		sliceGenerado[posicion] = numero + 1
	}

	//Mostrar en pantalla el resultado local
	fmt.Println("Vector actualizado en la goroutine main:")
	fmt.Println(sliceGenerado)

}
