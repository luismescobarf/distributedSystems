package main

import "fmt"

//Función Máximo Común Divisor de dos números (Teorema de Euclides)
//GCD -> Greatest Common Divisor (Máximo Común Divisor)
//Versión -> Single Thread
func gcd_st(x int, y int) int {

	//Relizar el proceso hasta que se satisface el condicional general
	for {

		//Criterio de Parada
		if x > 0 && y == 0 { //Sección 1

			return x

		} else { //Sección 2

			//Obtener el residuo de 'x' y 'y'
			r := x % y
			//Actualizar los parámetros de entrada (búsqueda del gcd)
			x = y
			y = r

		}

	} //Fin del ciclo que busca el gcd

} //Fin de la función gcd (Single Thread)

//Funciones Greatest Common Divisor Distribuído
///////////////////////////////////////////////

//Sección 1
func gcd_distribuido_seccion1(x int, y int, canalGCD chan int) {

	//Relizar el proceso hasta que se satisface el condicional general
	for {

		//Criterio de Parada
		if x > 0 && y == 0 { //Sección 1

			//Dejar resultado en el canal (retorno distribuído)
			canalGCD <- x

		} else { //Sección 2

			//Enviar 'x' por el canal
			canalGCD <- x

			//Enviar 'y' por el canal
			canalGCD <- y

			//Recibir los nuevos valores de 'x' y 'y'
			//procesados por el otro hilo o goroutine
			x, y = <-canalGCD, <-canalGCD

		}

	}

} //Fin de la Sección 1 gcd distribuído

//Sección 2
func gcd_distribuido_seccion2(canalGCD chan int) {

	//Recibir los nuevos valores de 'x' y 'y'
	//procesados por el otro hilo o goroutine
	x, y := <-canalGCD, <-canalGCD

	//Obtener el residuo de 'x' y 'y'
	r := x % y

	//Enviar 'y' por el canal en la posición de 'x' de comunicación
	canalGCD <- y

	//Enviar 'r' por el canal en la posición de 'y' de comunicación
	canalGCD <- r

} //Fin de la Sección 2 gcd distribuído

//Sección principal gcd
func main() {

	//Mostrar en pantalla el resultado del gcd de 9 y 6 (single thread)
	//fmt.Println( gcd_st(45,27) );

	//Sección gcd distribuído
	canalGCD := make(chan int, 4)
	var maximoComunDivisor int //Variable que recibe el resultado
	go gcd_distribuido_seccion1(9, 3, canalGCD)
	go gcd_distribuido_seccion2(canalGCD)

	//Obtener resultado almacenado en el canal
	maximoComunDivisor = <-canalGCD
	fmt.Println("GCD Sincrónico = ", maximoComunDivisor)

} //Final de la sección principal
