package main

//Paquetes requeridos
import (
	"fmt"
	"sync"
)

//Estructura de sincronización de goroutines
var wg sync.WaitGroup

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
func gcd_distribuido_seccion1(x int, y int, canalGCD chan int, maximoComunDivisor *int) {

	//Una vez se completa esta función o contexto, se cierra este hilo o rutina go
	defer wg.Done()

	//Relizar el proceso hasta que se satisface el condicional general
	for {

		//Criterio de Parada
		if x > 0 && y == 0 { //Sección 1

			////Salida de diagnóstico
			//fmt.Println("Estado x cuando se cumple criterio de parada = ",x)
			//fmt.Println("Estado y cuando se cumple criterio de parada = ",y)
			//var pausa string;
			//fmt.Scanln(&pausa);

			//Enviar resultado por referencia al hilo principal (main)
			*maximoComunDivisor = x

			//Salida de diagnóstico
			//fmt.Println("El resultado obtenido en la goroutine es = ",x);

			//Terminar el ciclo general del teorema euclidiano
			break

		} else { //Sección 2 Distribuída (el cómputo no se realiza en esta goroutine o hilo)

			////Salida de diagnóstico
			//fmt.Println("Enviando x de la Sección 1 a la 2 = ",x)
			//fmt.Println("Enviando y de la Sección 1 a la 2 = ",y)
			//var pausa string;
			//fmt.Scanln(&pausa);

			//Enviar 'x' por el canal
			canalGCD <- x

			//Enviar 'y' por el canal
			canalGCD <- y

			//Recibir los nuevos valores de 'x' y 'y' procesados por el otro hilo o goroutine
			x, y = <-canalGCD, <-canalGCD

			////Salida de diagnóstico
			//fmt.Println("Recibiendo x en la Sección 1 después del cómputo de la Sección 2 = ",x)
			//fmt.Println("Recibiendo y en la Sección 1 después del cómputo de la Sección 2 = ",y)
			//fmt.Scanln(&pausa);

		}

	}

	//Avisar la terminación de la goroutine (alternativa al defer)
	//wg.Done();

} //Fin de la Sección 1 gcd distribuído

//Sección 2
func gcd_distribuido_seccion2(canalGCD chan int) {

	//Una vez se completa esta función o contexto, se cierra este hilo o rutina go
	defer wg.Done()

	////Variable para diagnóstico
	//var pausa string;

	for {

		//Recibir los nuevos valores de 'x' y 'y' procesados por el otro hilo o goroutine
		x, y := <-canalGCD, <-canalGCD

		////Salida de diagnóstico
		//fmt.Println("x recibido en la Sección 2 = ",x)
		//fmt.Println("y recibido en la Sección 2 = ",y)
		//fmt.Println("r generado en la Sección 2 = ",x % y)
		//fmt.Scanln(&pausa);

		//Obtener el residuo de 'x' y 'y'
		r := x % y

		//Enviar 'y' por el canal en la posición de 'x' de comunicación
		canalGCD <- y

		//Enviar 'r' por el canal en la posición de 'y' de comunicación
		canalGCD <- r

		//Terminar el ciclo del "servicio" para obtener los residuos
		if r == 0 {
			break
		}

	}

	//Avisar la terminación del proceso
	//wg.Done();

} //Fin de la Sección 2 gcd distribuído

//Sección principal llamado a implementaciones GCD Single-Thread y GCD Multi-Thread
func main() {

	//Argumentos para la función máximo común divisor
	var x, y int
	x = 45
	y = 27

	//Mostrar en pantalla el resultado del gcd (single thread)
	fmt.Println("GCD Single-Thread (", x, ",", y, ") = ", gcd_st(x, y))

	//Sección gcd distribuído
	canalGCD := make(chan int) //Canal de comunicación entre los hilos o rutinas go donde se está distribuyendo el cómputo
	//canalGCD := make(chan int,4);//Esta declaración del canal, en lugar de la anterior, cambia la interacción a asincrónica
	var maximoComunDivisor int //Variable que recibe el resultado
	wg.Add(2)
	//Iniciar los dos procesos independientes
	go gcd_distribuido_seccion1(x, y, canalGCD, &maximoComunDivisor)
	go gcd_distribuido_seccion2(canalGCD)
	wg.Wait()

	//Mostrar el resultado obtenido después de que se acaba la sección distribuída
	fmt.Println("GCD Sincrónico = ", maximoComunDivisor)

} //Final de la sección principal
