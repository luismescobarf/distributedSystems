package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

/////////////
//Nodos
/////////////
type nodo struct {
	id int
	x  float32
	y  float32
}

/////////////
//Aristas
/////////////
type arista struct {
	id    string
	i     nodo
	j     nodo
	costo int
}

//Crear una nueva arista
func (a *arista) cargarNodos(i nodo, j nodo) {
	a.id = strconv.FormatInt(int64(i.id-1), 10) + "-" + strconv.FormatInt(int64(j.id-1), 10)
	a.i = i
	a.j = j
	a.calcularCosto()
}

//Distancia Euclidiana 2D de la arista
func (a *arista) calcularCosto() {
	a.costo = int(math.Sqrt(float64((a.i.x-a.j.x)*(a.i.x-a.j.x)+(a.i.y-a.j.y)*(a.i.y-a.j.y))) + 0.5)
	//int((((p_1[0]-p_2[0])**2 + (p_1[1]-p_2[1])**2) * *(1 / 2)) + 0.5)
}

/////////////
//Red o grafo
/////////////
type grafo struct {
	red              map[string]arista //Representación por tabla hash o arreglo asociativo
	matrizAdyacencia [][]int           //Representación por matriz de adyacencia del grafo
	aristas          []arista          //Slice de aristas
	nodos            []nodo            //Slice de nodos
	n                int               //Número de nodos del grafo
}

//Realizarlo como receiver de la estructura (actualmente está en el código)
func (g *grafo) generarMatrizAdyacencia() {

}

//Calcular costo de una ruta sobre el grafo
func (g grafo) costoTour(tour []int) int {
	//Costo de la ruta
	var costo int = 0

	//Recorrer todo el tour y acumular costos
	for i := 0; i < len(tour)-1; i++ {
		costo += g.red[etiquetaArista(tour[i], tour[i+1])].costo
	}

	//Incorporar costo del retorno
	costo += g.red[etiquetaArista(tour[len(tour)-1], tour[0])].costo

	//Retornar el costo del tour
	return costo
}

///////////////////////////////////////////Funciones generales

//Función para generar etiquetas a partir del subíndice
func etiquetaArista(i int, j int) string {
	return strconv.FormatInt(int64(i), 10) + "-" + strconv.FormatInt(int64(j), 10)
}

//Función para lectura del archivo que contiene la instancia tsp
func leerLineas(path string) ([]nodo, error) {

	//Abrir el archivo y retornar error si falla
	archivo, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer archivo.Close() //Cerrar el archivo cuando acabe la función

	//Arreglo que recibirá los nodos contenidos en las líneas del archivo
	var nodos []nodo

	//Cargar y recorrer el archivo
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {

		//Obtener cada línea del archivo en formato string
		lineaNodo := strings.Split(scanner.Text(), " ")

		//Convertir cada línea string a slice float32
		var lineaNodoNumerica []float32
		for _, c := range lineaNodo {
			cn, _ := strconv.Atoi(c)
			lineaNodoNumerica = append(lineaNodoNumerica, float32(cn))
		}

		//Preparar el nodo
		nodoAux := nodo{id: int(lineaNodoNumerica[0]), x: lineaNodoNumerica[1], y: lineaNodoNumerica[2]}

		//Agregar el nodo a la colección
		nodos = append(nodos, nodoAux)

	}

	//Retornar la colección de nodos y el estado de la lectura
	return nodos, scanner.Err()
}

//Sección principal
func main() {

	//Llamado a la función de lectura del archivo con el caso de prueba
	var ruta string
	ruta = "./eil51.tsp" //Especificación del caso de prueba
	nodos, _ := leerLineas(ruta)

	// //Salida de diagnóstico
	// fmt.Println("Nodos cargados:")
	// fmt.Println(nodos)

	//Generar las aristas
	var aristas []arista           //Arreglo de aristas
	red := make(map[string]arista) //Hash de aristas
	for i, nodo_i := range nodos {
		for j, nodo_j := range nodos {

			//Si es una conexión consigo mismo, colocar un valor muy alto
			if i != j {
				//Adicionar al arreglo de aristas
				aristaAux := &arista{}
				aristaAux.cargarNodos(nodo_i, nodo_j)
				aristas = append(aristas, *aristaAux)
				//Adicionar al hash de aristas
				red[aristaAux.id] = *aristaAux
			}

		}
	}

	// //Salida de diagnóstico
	// fmt.Println("Aristas cargadas:")
	// fmt.Println(aristas[0])

	//Generar matriz de adyacencia
	var matrizAdyacencia [][]int
	for i := 0; i < len(nodos); i++ {
		//Preparar el contenedor de la fila
		var fila []int
		for j := 0; j < len(nodos); j++ {
			fila = append(fila, math.MaxInt64)
		}
		//Llenar la matriz de adyacencia con los costos correspondientes
		for j := 0; j < len(nodos); j++ {
			if i != j {
				fila[j] = red[etiquetaArista(i, j)].costo
			}
		}
		//Adicionar la fila a la matriz de adyacencia
		matrizAdyacencia = append(matrizAdyacencia, fila)

	}

	//Construir la red con la información generada
	instanciaTSP := &grafo{
		red:              red,
		nodos:            nodos,
		aristas:          aristas,
		n:                len(nodos),
		matrizAdyacencia: matrizAdyacencia,
	}

	// //Salida de diagnóstico
	// fmt.Println("Mostrar caso TSP cargado:")
	// fmt.Println(instanciaTSP)

	//Obtener costo de solución óptima (validación de cálculo de la función objetivo)
	optimaEil51 := []int{1, 22, 8, 26, 31, 28, 3, 36, 35, 20, 2, 29, 21, 16, 50, 34, 30, 9, 49, 10, 39, 33, 45, 15, 44, 42, 40, 19, 41, 13, 25, 14, 24, 43, 7, 23, 48, 6, 27, 51, 46, 12, 47, 18, 4, 17, 37, 5, 38, 11, 32}
	//Ajustar id's
	for i := 0; i < len(optimaEil51); i++ {
		optimaEil51[i] = optimaEil51[i] - 1
	}

	//Salida de diagnóstico
	fmt.Println("Ruta con índices actualizados:")
	fmt.Println(optimaEil51)
	fmt.Println("Costo óptimo literal -> ", instanciaTSP.costoTour(optimaEil51))

	//Prueba de ordenamiento de un slice de aristas
	seleccionPrimerasAristas := []arista{
		aristas[0],
		aristas[1],
		aristas[2],
		aristas[3],
		aristas[4],
		aristas[5],
	}
	fmt.Println("Aristas antes de ordenamiento:")
	fmt.Println(seleccionPrimerasAristas)

	//Criterio de ordenamiento
	sort.SliceStable(seleccionPrimerasAristas, func(i, j int) bool {
		return seleccionPrimerasAristas[i].costo < seleccionPrimerasAristas[j].costo
	})

	fmt.Println("Aristas después de ordenamiento:")
	fmt.Println(seleccionPrimerasAristas)

}
