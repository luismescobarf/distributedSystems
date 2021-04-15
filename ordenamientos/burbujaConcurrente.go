package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

//Estructura para fragmentación del vector
type FragmentoVector struct {
	valorMenor       int
	valorMayor       int
	tamanioFragmento int
	fragmentoVector  []int
}

func cargar(vec *[]int, aleatorio *rand.Rand) {
	//for f := 0; f < len(vec); f++ {
	for f := 0; f < 300000; f++ {
		//vec[f] = aleatorio.Intn(100)
		*vec = append(*vec, aleatorio.Intn(300000))
	}
}

func cargarDistintos(vec *[]int, aleatorio *rand.Rand) {

	var numeroElementos int = 20
	for i := 0; i < numeroElementos; i++ {
		*vec = append(*vec, i+1)
	}

}

//Función de ordenamiento versión Single-Thread
func ordenar(vec *[50000]int) {
	for k := 1; k < len(vec); k++ {
		for f := 0; f < len(vec)-k; f++ {
			if vec[f] > vec[f+1] {
				aux := vec[f]
				vec[f] = vec[f+1]
				vec[f+1] = aux
			}
		}
	}
	wg.Done()
}

//Función para ordenar los fragmentos estructurados de manera secuencial
func ordenarFragmentoSecuencial(fragmento FragmentoVector) {

	//Realizar ordenamiento burbuja
	for k := 1; k < len(fragmento.fragmentoVector); k++ {
		for f := 0; f < len(fragmento.fragmentoVector)-k; f++ {
			if fragmento.fragmentoVector[f] > fragmento.fragmentoVector[f+1] {
				aux := fragmento.fragmentoVector[f]
				fragmento.fragmentoVector[f] = fragmento.fragmentoVector[f+1]
				fragmento.fragmentoVector[f+1] = aux
			}
		}
	}

	//Una vez ordenados, actualizar los atributos de la estructura para la coordinación posterior a la concurrencia
	fragmento.valorMenor = fragmento.fragmentoVector[0]
	fragmento.valorMayor = fragmento.fragmentoVector[len(fragmento.fragmentoVector)-1]
	fragmento.tamanioFragmento = len(fragmento.fragmentoVector)

	////Salida de diagnóstico
	//fmt.Println("Diagnóstico de fragmento -> ",fragmento);
	//var pausa int;
	//fmt.Scanln(&pausa);

} //Fin de la función de ordenamiento de fragmento

//Función para ordenar los fragmentos estructurados
func ordenarFragmento(fragmento *FragmentoVector) {

	//Realizar ordenamiento burbuja
	for k := 1; k < len(fragmento.fragmentoVector); k++ {
		for f := 0; f < len(fragmento.fragmentoVector)-k; f++ {
			if fragmento.fragmentoVector[f] > fragmento.fragmentoVector[f+1] {
				aux := fragmento.fragmentoVector[f]
				fragmento.fragmentoVector[f] = fragmento.fragmentoVector[f+1]
				fragmento.fragmentoVector[f+1] = aux
			}
		}
	}

	//Una vez ordenados, actualizar los atributos de la estructura para la coordinación posterior a la concurrencia
	fragmento.valorMenor = fragmento.fragmentoVector[0]
	fragmento.valorMayor = fragmento.fragmentoVector[len(fragmento.fragmentoVector)-1]
	fragmento.tamanioFragmento = len(fragmento.fragmentoVector)

	////Salida de diagnóstico
	//fmt.Println("Diagnóstico de fragmento -> ",fragmento);
	//var pausa int;
	//fmt.Scanln(&pausa);

	//Avisar la terminación del ordenamiento del fragmento
	wg.Done()

} //Fin de la función de ordenamiento de fragmento

//Función para ordenar los fragmentos ordenados
func ordenarFragmentosOrdenados(sliceFragmentos []FragmentoVector) []FragmentoVector {

	//Realizar ordenamiento burbuja
	for k := 1; k < len(sliceFragmentos); k++ {
		for f := 0; f < len(sliceFragmentos)-k; f++ {
			//if sliceFragmentos[f].valorMayor > sliceFragmentos[f+1].valorMayor {
			if sliceFragmentos[f].valorMenor > sliceFragmentos[f+1].valorMenor {
				aux := sliceFragmentos[f]
				sliceFragmentos[f] = sliceFragmentos[f+1]
				sliceFragmentos[f+1] = aux
			}
		}
	}

	return sliceFragmentos

} //Fin de la función de ordenamiento de fragmento

//Función para calcular el tiempo transcurrido
func diferenciaTiempo(hora1, hora2 time.Time) time.Duration {
	diferencia := hora2.Sub(hora1)
	return diferencia
}

//Sección principal ordenamiento burbuja distribuído
func main() {

	//Obtener el número de procesadores (capacidad de cómputo donde se ejecute la implementación)
	numeroProcesadores := runtime.NumCPU()
	//numeroProcesadores = numeroProcesadores-1;//Dejar libre un procesador para el fragmento excedente

	//Generar el vector que se va a ordenar
	var vec1 []int
	aleatorio := rand.New(rand.NewSource(time.Now().UnixNano()))
	//cargar(&vec1, aleatorio)
	cargarDistintos(&vec1, aleatorio)
	fmt.Println(vec1)
	rand.Shuffle(len(vec1), func(i, j int) {
		vec1[i], vec1[j] = vec1[j], vec1[i]
	})

	//Salida de diagnóstico
	fmt.Println(vec1)
	var pausaCarga string
	fmt.Scanln(&pausaCarga)

	//fmt.Println("Número de Procesadores = ", numeroProcesadores)

	var tamanioFragmentos int
	tamanioFragmentos = len(vec1) / numeroProcesadores

	residuoRelacion := len(vec1) % numeroProcesadores
	if residuoRelacion != 0 {
		tamanioFragmentos += 1
	}

	//Salidas de diagnóstico
	fmt.Println("Tamaño del vector = ", len(vec1))
	fmt.Println("Número de procesadores = ", numeroProcesadores)
	fmt.Println("Tamaño de los fragmentos = ", tamanioFragmentos)
	fmt.Println("Residuo relación = ", residuoRelacion)

	//Generar slice de fragmentos
	var sliceFragmentos []FragmentoVector

	//Recorrer el vector que se va a ordenar para fragmentarlo
	i := 0 //Subíndice Contador general por inferencia
	for i < len(vec1) {

		////Campos estructura fravmento
		//type FragmentoVector struct{
		//	valorMayor int
		//	valorMenor int
		//	fragmentoVector[] int
		//}

		//Adicionar elementos al fragmento
		var EstructuraAuxiliar FragmentoVector
		var j int
		for j = 0; j < tamanioFragmentos && i+j < len(vec1); j++ {
			EstructuraAuxiliar.fragmentoVector = append(EstructuraAuxiliar.fragmentoVector, vec1[i+j])
		}

		//Adicionar el fragmento al vector de estructuras
		sliceFragmentos = append(sliceFragmentos, EstructuraAuxiliar)

		//Actualizar el índice del vector general
		i = j + i

	} //Fin del ciclo de fragmentación del vector

	//Salida de diagnóstico
	fmt.Println("Slice o vector de estructuras: ")
	fmt.Println(sliceFragmentos)

	//Estructuras par ala toma de tiempos
	var hora1, hora2 time.Time

	//Realizar ordenamientos
	hora1 = time.Now() //Tiempo inicial
	for _, fragmento := range sliceFragmentos {

		//Ordenar cada uno de los fragmentos
		ordenarFragmentoSecuencial(fragmento)

		//sliceFragmentos[indice] = fragmento

	}
	hora2 = time.Now() //Tiempo final
	di := diferenciaTiempo(hora1, hora2)
	fmt.Println("Tiempo Secuencial = ", di.Seconds())
	var pausa string
	fmt.Scanln(&pausa)

	//Generar hilos de ordenamiento y asignarles carga computacional
	hora1 = time.Now() //Tiempo inicial
	wg.Add(numeroProcesadores)
	for s := 0; s < len(sliceFragmentos); s++ {

		//Ordenar cada uno de los fragmentos
		go ordenarFragmento(&sliceFragmentos[s])

	}
	wg.Wait()
	hora2 = time.Now() //Tiempo final
	di = diferenciaTiempo(hora1, hora2)
	fmt.Println("Tiempo Concurrente = ", di.Seconds())
	fmt.Scanln(&pausa)

	//Salida de diagnóstico
	fmt.Println("Fragmentos ordenados: ")
	fmt.Println(sliceFragmentos)

	////Salida de diagnóstico
	//fmt.Println("Index slice ");
	//fmt.Println(sliceFragmentos[0]);

	sliceFragmentos = ordenarFragmentosOrdenados(sliceFragmentos)

	//Salida de diagnóstico
	fmt.Println("Final: ")
	fmt.Println(sliceFragmentos)

	//Realizar mezcla de fragmentos para producir el vector final ordenado
	//Realizar mezcla de fragmentos para producir el vector final ordenado
	//Realizar mezcla de fragmentos para producir el vector final ordenado

}
