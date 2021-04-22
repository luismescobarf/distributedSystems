package main

//Importar librerías
import (
	"fmt"
	"sync"
)

//Porción de trabajo para distribuír y recoger el resultado
type porcionTrabajo struct {
	x       int
	dobleX  int
	tripleX int
}

//Sincronización global
var wg sync.WaitGroup

//Función para generar partes del trabajo
func generarSlicePorcionesTrabajo() []porcionTrabajo {

	//Contenedor de retorno
	var coleccionPorcionesTrabajo []porcionTrabajo

	var n int = 10
	for i := 0; i < n; i++ {
		//Si es par, generar el fragmento
		if i%2 == 0 {
			coleccionPorcionesTrabajo = append(coleccionPorcionesTrabajo, porcionTrabajo{x: i, dobleX: 0, tripleX: 0})
		}

	}

	//Retornar la colección de fragmentos para distribuir concurrentemente
	return coleccionPorcionesTrabajo

}

//Realizar los procesos de forma concurrente
func trabajadoresConcurrentes(entrada []chan porcionTrabajo, salida []chan porcionTrabajo) {

	//Sincronización de los trabajadores que se van a generar en este contexto
	//var wg sync.WaitGroup

	//Un trabajador que va a escuchar por cada uno de los canales
	wg.Add(len(entrada))

	//Recorrer todos los canales
	for i, c := range entrada {

		//Por cada canal, generar una rutina go que esté escuchando el canal y procesando
		go func(canal chan porcionTrabajo, i int) {

			//Cuando el trabajador termina su trabajo con lo que recibió por el canal
			defer wg.Done()

			//Rutina go escuchando el canal
			for p := range canal {

				//Realizar el cómputo que le corresponde
				p.dobleX = p.x * 2
				p.tripleX = p.x * 3

				//Hacer el envío
				salida[i] <- p

			}

			//Cerrar el canal de salida cuando termina de escuchar y escribir
			close(salida[i])

		}(c, i) //Enviamos el canal a la rutina go para que lo escuche

	}

}

//Recopilador (una rutina go recopila de cada canal)
func recopiladorTrabajoConcurrente(entrada []chan porcionTrabajo) {

	//Sincronización de los trabajadores que van a realizar la lectura
	//var wg sync.WaitGroup

	//Un trabajador que va a escuchar por cada uno de los canales
	wg.Add(len(entrada))

	//Recorrer todos los canales de entrada
	for _, c := range entrada {

		//Generar una rutina go encargada de cada canal
		go func(canal chan porcionTrabajo) {

			//Goroutine escuchando todo lo que llega por el canal
			for p := range canal {
				//fmt.Println("Fragmento recibido lectura concurrente: ", <-canal)
				fmt.Println("Fragmento recibido lectura concurrente: ", p)
			}

			//Cuando el trabajador termina de recopilar
			wg.Done()

		}(c)

	}

}

//Sección principal del ejemplo
func main() {

	//Finalidad del ejemplo en pantalla
	//fmt.Println("Transmisión por colecciones de canales")

	//Generar porciones para procesar
	porcionesTrabajo := generarSlicePorcionesTrabajo()

	//Salida de diagnóstico
	fmt.Println("Porciones de trabajo")
	fmt.Println(porcionesTrabajo)

	//Arreglos de canales para trabajo simultáneo
	var envioArregloCanales []chan porcionTrabajo
	var recepcionArregloCanales []chan porcionTrabajo
	//Proceso para que sean canales con valor cero, y no valor nulo
	for i := 0; i < len(porcionesTrabajo); i++ {
		canalAuxiliar1 := make(chan porcionTrabajo)
		envioArregloCanales = append(envioArregloCanales, canalAuxiliar1)
		canalAuxiliar2 := make(chan porcionTrabajo)
		recepcionArregloCanales = append(recepcionArregloCanales, canalAuxiliar2)
	}

	//Se realizarán tantos envíos como porciones de trabajo se tengan
	wg.Add(len(porcionesTrabajo))

	//Distribuir el trabajo concurrentemente
	for i := 0; i < len(envioArregloCanales); i++ {
		go func(i int) {
			defer wg.Done()
			envioArregloCanales[i] <- porcionesTrabajo[i]
			close(envioArregloCanales[i])
		}(i)
	}

	//Generar los trabajadores que escuchan los canales y escriben su cómputo
	go trabajadoresConcurrentes(envioArregloCanales, recepcionArregloCanales)

	//Lectura concurrente de las respuestas
	go recopiladorTrabajoConcurrente(recepcionArregloCanales)

	//Esperar terminación de la recopilación
	wg.Wait()

	//Forma alternativa de sincronización
	//<-time.After(time.Second * 10)

}
