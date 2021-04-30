package main

import (
	"fmt"
	"math"
	"strconv"
	"time"

	//"math"
	"io/ioutil"
	"os"
	"strings"
)

// 0 <= index <= len(a)
func insert(a []int, index int, value int) []int {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func ordenarMayor(listNum [][]int, Cant int) [][]int {
	var tmp []int
	for x := 0; x < Cant; x++ {
		for y := 0; y < Cant; y++ {

			if listNum[x][0] > listNum[y][0] {
				//fmt.Println(listNum[x])
				tmp = listNum[x]
				listNum[x] = listNum[y]
				listNum[y] = tmp
				//fmt.Println(y)
			}
		}
	}
	//fmt.Print("\nArray dinamico ordenado: ")
	//fmt.Println(listNum)
	//fmt.Println()
	return listNum
}
func ordenarMenor(listNum [][]int, Cant int) [][]int {
	var tmp []int
	for x := 0; x < Cant; x++ {
		for y := 0; y < Cant; y++ {

			if listNum[x][2] < listNum[y][2] {
				//fmt.Println(listNum[x])
				tmp = listNum[x]
				listNum[x] = listNum[y]
				listNum[y] = tmp
				//fmt.Println(y)
			}
		}
	}
	//fmt.Print("\nArray dinamico ordenado: ")
	//fmt.Println(listNum)
	//fmt.Println()
	return listNum
}

func cubrirNodo(nodo int, nodosCubiertos map[int]struct{}, nodosSinCubrir map[int]struct{}) {

	//Agregar al cunjunto de nodos cubiertos
	nodosCubiertos[nodo] = struct{}{}

	//Eliminar del conjunto de nodos cubiertos
	delete(nodosSinCubrir, nodo)
	println("ELIMINADO")

}

func main() {
	var numeroNodos int
	var matrizAdyacencia [][]int
	var matrizAdyacenciaDIST [][]int
	var listadoInsercioens [][]int
	var tour []int
	var costoOriginalArista int

	var posiblesVerticesEnvolvente [][]int
	//var p_1 []int
	//var p_2 []int
	//			########OBTENER INFO DEL ARCHIVO, CREAR ARRAY CON DATOS################
	var nombreArchivo = "./eil51.tsp"
	//var nombreArchivo = "./st70.tsp"

	//fmt.Println(nombreArchivo)

	var coordenadasPuntos [][]int
	var manejadorArchivo, err = os.Open(nombreArchivo)
	if err != nil {
		panic(err)
	}
	//leo el documento completo
	c, err := ioutil.ReadAll(manejadorArchivo)
	if err != nil {
		panic(err)
	}

	// sao linea a linea del documento leido
	for _, line := range strings.Split(strings.TrimRight(string(c), "\n"), "\n") {

		//fmt.Println("###linea####")
		// un array de cada linea
		arrayst := strings.Split(line, " ")
		var t2 = []int{}
		//for encargado de convertir el array String a array Int
		for _, cha := range arrayst {
			j, err := strconv.Atoi(cha)
			if err != nil {
				panic(err)
			}
			t2 = append(t2, j)
		}
		coordenadasPuntos = append(coordenadasPuntos, t2)

	}
	//fmt.Println(coordenadasPuntos)

	//fmt.Printf("### Contenido de archivo: ###\n%s\n", string(c))
	manejadorArchivo.Close()
	numeroNodos = len(coordenadasPuntos)
	//fmt.Println(numeroNodos)

	// ############## #Construir matriz de adyacencia ########################
	for i := 0; i < numeroNodos; i++ {
		var arraynt []int
		for j := 0; j < numeroNodos; j++ {
			arraynt = append(arraynt, 0)
		}
		matrizAdyacencia = append(matrizAdyacencia, arraynt)
	}

	matrizAdyacenciaDIST = matrizAdyacencia
	//fmt.Println(matrizAdyacencia)

	for i := 0; i < numeroNodos; i++ {
		for j := 0; j < numeroNodos; j++ {
			if i != j {
				p_1 := []int{coordenadasPuntos[i][1], coordenadasPuntos[i][2]}
				///fmt.Println(p_1)
				//fmt.Println("valor P_1")
				p_2 := []int{coordenadasPuntos[j][1], coordenadasPuntos[j][2]}
				//fmt.Println(p_2)
				//fmt.Println("valor P_2")
				part1 := math.Pow(float64((p_1[0] - p_2[0])), 2)
				part2 := math.Pow(float64((p_1[1] - p_2[1])), 2)
				part3 := part1 + part2
				part4 := int(math.Sqrt(part3) + 0.5)
				//sumacua := math.Pow(math.Pow(float64((p_1[0]-p_2[0])), 2)+math.Pow(float64((p_1[1]-p_2[1])), 2), (1 / 2))
				distanciaEuclidiana_p1_p2 := part4
				//distanciaEuclidiana_p1_p2 := int(math.Pow(float64(math.Pow(float64((p_1[0] - p_2[0])),2)+math.Pow(float64((p_1[1] - p_2[1]))),2)),2)
				//(math.Pow(float64((math.Pow(float64((p_1[0] - p_2[0])),2)+math.Pow(float64((p_1[1] - p_2[1]))),2)),(1/2))+0.5)
				//distanciaEuclidiana_p1_p2 := int((((p_1[1] - p_2[1])**2 + (p_1[1] - p_2[1])**2)**(1/2))+0.5)
				//matrizAdyacencia[i][j] = distanciaEuclidiana_p1_p2
				matrizAdyacenciaDIST[i][j] = distanciaEuclidiana_p1_p2
			}
		}
	}

	nodosSinCubrir := make(map[int]struct{})
	nodosCubiertos := make(map[int]struct{})

	for i := 0; i < numeroNodos; i++ {
		nodosSinCubrir[i] = struct{}{}
	}

	//Mostrar conjuntos antes de actualización
	fmt.Println("Nodos sin cubrir:")
	fmt.Println(nodosSinCubrir)
	fmt.Println("Nodos cubiertos:")
	fmt.Println(nodosCubiertos)

	//Realizar cobertura de un nodo
	//var nodoCubierto int = 4
	//cubrirNodo(nodoCubierto, nodosCubiertos, nodosSinCubrir)
	nodoCubierto := 0
	//Mostrar conjuntos después de actualización
	fmt.Println("Nodo seleccionado = ", nodoCubierto)
	fmt.Println("Nodos sin cubrir:")
	fmt.Println(nodosSinCubrir)
	fmt.Println("Nodos cubiertos:")
	fmt.Println(nodosCubiertos)

	//fmt.Println(nodosSinCubrir)

	cubrirNodo(nodoCubierto, nodosCubiertos, nodosSinCubrir)
	tour = append(tour, nodoCubierto)
	//nodosCubiertos[nodoInicial] = nodoInicial
	//delete(nodosSinCubrir, nodoInicial)
	//tour = append(tour, nodoInicial)

	// ############Envolvente convexa más sencilla (triángulo) ################

	for i := 0; i < numeroNodos; i++ {
		if matrizAdyacencia[nodoCubierto][i] != 0 && i != nodoCubierto {
			//arrayi :=  matrizAdyacencia[nodoCubierto][i]
			var arraynt []int
			arraynt = append(arraynt, matrizAdyacencia[nodoCubierto][i])
			arraynt = append(arraynt, i)
			posiblesVerticesEnvolvente = append(posiblesVerticesEnvolvente, arraynt)
			//fmt.Println(arraynt)
			//time.Sleep(5 * time.Second)
			//posiblesVerticesEnvolvente.append( [ matrizAdyacencia[ nodoInicial ][i] , i ] )
		}
	}

	///////        #Ordenar por el criterio (más cercano o menor)    /////////

	//fmt.Println(posiblesVerticesEnvolvente)
	//fmt.Println(posiblesVerticesEnvolvente[0][0])
	//fmt.Println(len(posiblesVerticesEnvolvente))
	posiblesVerticesEnvolvente = ordenarMayor(posiblesVerticesEnvolvente, len(posiblesVerticesEnvolvente))
	//fmt.Println(posiblesVerticesEnvolvente)

	////// ####### AÑADIR A LA SOLUCION ################### ///////////
	tour = append(tour, posiblesVerticesEnvolvente[0][1]) //#Primera esquina más lejana de la envolvente
	cubrirNodo(posiblesVerticesEnvolvente[0][1], nodosCubiertos, nodosSinCubrir)

	tour = append(tour, posiblesVerticesEnvolvente[1][1]) //#Segunda esquina más lejana de la envolvente
	cubrirNodo(posiblesVerticesEnvolvente[1][1], nodosCubiertos, nodosSinCubrir)

	////////////// ######Criterio del constructivo###### \\\\\\\\\\\\\\\\\\\\\\\\\\
	cont := 0
	for len(nodosSinCubrir) != 0 {

		cont += 1
		fmt.Println("TAMAÑO DEL NODOSINCUBRIR ", len(nodosSinCubrir))
		fmt.Println("TOUR", tour)

		//una Iteracion de la insercion
		for i := 0; i < len(tour)-1; i++ {
			println("TAMAÑO DEL TOUR ", len(tour))
			fmt.Println("TAMAÑO DEL NODOSINCUBRIR ", len(nodosSinCubrir))
			println("VALOR I ", i)

			//tour[i], tour[i+1]
			costoOriginalArista = matrizAdyacencia[tour[i]][tour[i+1]]

			for k, _ := range nodosSinCubrir {

				costoInsercionK := 0
				costoInsercionK += matrizAdyacencia[tour[i]][k]
				costoInsercionK += matrizAdyacencia[k][tour[i+1]]
				diferenciaInsercion := costoInsercionK - costoOriginalArista
				var arrayin []int
				arrayin = append(arrayin, k)
				arrayin = append(arrayin, i+1)
				arrayin = append(arrayin, diferenciaInsercion)

				listadoInsercioens = append(listadoInsercioens, arrayin)

			}
		}
		println("SALIIIIIIIIIIIIII")

		costoOriginalArista = matrizAdyacencia[tour[len(tour)-1]][tour[0]]

		for k, _ := range nodosSinCubrir {

			costoInsercionK := 0
			costoInsercionK += matrizAdyacencia[tour[len(tour)-1]][k]
			costoInsercionK += matrizAdyacencia[k][tour[0]]
			diferenciaInsercion := costoInsercionK - costoOriginalArista
			var arrayin []int
			arrayin = append(arrayin, k)
			arrayin = append(arrayin, 0)
			arrayin = append(arrayin, diferenciaInsercion)

			listadoInsercioens = append(listadoInsercioens, arrayin)

		}

		println("SALIIIIIIIIIIIIII 2")
		listadoInsercioens = ordenarMenor(listadoInsercioens, len(listadoInsercioens))

		////  ##############3#Añadir a la solución###########\\\\\\\\\\\
		tour = insert(tour, listadoInsercioens[cont][1], listadoInsercioens[cont][0])
		cubrirNodo(listadoInsercioens[0][0], nodosCubiertos, nodosSinCubrir)
		println("LISTA INSERIONES ", listadoInsercioens[cont][0])
		time.Sleep(3 * time.Second)

	}
	//fmt.Println()

}
