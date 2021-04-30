package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"time"
)

type posibleNodo struct {
	id int
	x  int
}

type posibleInsercion struct {
	a          int
	b          int
	diferencia int
}

type solucion struct {
	ruta  []int
	costo int
}

func insertionTSP(matrizAdyacencia [][]int, nodoInicial int) solucion {

	var tour []int

	mapaCubierto := make(map[int]struct{})
	mapaSinCubrir := make(map[int]struct{})
	// fmt.Print(numeroNodos)

	numeroNodos := len(matrizAdyacencia)

	for i := 0; i < numeroNodos; i++ {
		mapaSinCubrir[i] = struct{}{}
	}

	//nodoInicial := 0
	mapaCubierto[nodoInicial] = struct{}{}
	delete(mapaSinCubrir, nodoInicial)
	tour = append(tour, nodoInicial)

	var posiblesVerticesEnvolvente []posibleNodo

	for i := 0; i < numeroNodos; i++ {
		if matrizAdyacencia[nodoInicial][i] != 0 && i != nodoInicial {

			posiblesVerticesEnvolvente = append(posiblesVerticesEnvolvente, posibleNodo{id: i, x: matrizAdyacencia[nodoInicial][i]})
		}
	}

	sort.SliceStable(posiblesVerticesEnvolvente, func(i, j int) bool {
		return posiblesVerticesEnvolvente[i].x > posiblesVerticesEnvolvente[j].x
	})

	tour = append(tour, posiblesVerticesEnvolvente[0].id)
	mapaCubierto[posiblesVerticesEnvolvente[0].id] = struct{}{}
	delete(mapaSinCubrir, posiblesVerticesEnvolvente[0].id)
	tour = append(tour, posiblesVerticesEnvolvente[1].id)
	mapaCubierto[posiblesVerticesEnvolvente[1].id] = struct{}{}
	delete(mapaSinCubrir, posiblesVerticesEnvolvente[1].id)

	//fmt.Println(" Tour inicial -> ", tour)

	v := len(mapaSinCubrir)
	for v > 0 {

		var listadoInserciones []posibleInsercion

		for i := 0; i < len(tour)-1; i++ {
			costoOriginalArista := matrizAdyacencia[tour[i]][tour[i+1]]

			for k := range mapaSinCubrir {
				costoInsercionK := 0
				costoInsercionK = costoInsercionK + matrizAdyacencia[tour[i]][k]
				costoInsercionK = costoInsercionK + matrizAdyacencia[k][tour[i+1]]
				diferenciaInsercion := costoInsercionK - costoOriginalArista

				listadoInserciones = append(listadoInserciones, posibleInsercion{a: k, b: i + 1, diferencia: diferenciaInsercion})
			}
		}

		aux := tour[len(tour)-1:]
		costoOriginalArista := matrizAdyacencia[aux[0]][tour[0]]

		for k := range mapaSinCubrir {
			costoInsercionK := 0
			costoInsercionK = costoInsercionK + matrizAdyacencia[aux[0]][k]
			costoInsercionK = costoInsercionK + matrizAdyacencia[k][tour[0]]
			diferenciaInsercion := costoInsercionK - costoOriginalArista

			listadoInserciones = append(listadoInserciones, posibleInsercion{a: k, b: 0, diferencia: diferenciaInsercion})
		}

		sort.SliceStable(listadoInserciones, func(i, j int) bool {
			return listadoInserciones[i].diferencia < listadoInserciones[j].diferencia
		})

		que := listadoInserciones[0].a
		donde := listadoInserciones[0].b
		tour = append(tour[:donde+1], tour[donde:]...)
		tour[donde] = que
		mapaCubierto[que] = struct{}{}
		delete(mapaSinCubrir, que)
		//fmt.Println(" Tour ->", tour)

		v = len(mapaSinCubrir)
	}

	cont := len(tour) - 1
	fo := 0
	for i := 0; i < cont; i++ {
		fo = fo + matrizAdyacencia[tour[i]][tour[i+1]]
	}
	fo = fo + matrizAdyacencia[tour[cont]][tour[0]]
	// fmt.Print()
	// fmt.Print("Función Objetivo->", fo)
	// fmt.Print(tour)
	return solucion{ruta: tour, costo: fo}

}

func main() {
	// Lee el archivo tsp
	//nombre_archivo := "eil51.tsp"
	//nombre_archivo := "wi29.tsp"
	nombre_archivo := "qa194.tsp"
	//nombre_archivo := "uy734.tsp"
	//nombre_archivo := "zi929.tsp"
	f, err := os.Open(nombre_archivo)
	if err != nil {
		fmt.Println(err)
		return
	}
	//Slice que recibe los datos
	var slice []float32
	for {
		//var a int
		var a float32
		_, err := fmt.Fscan(f, &a)

		if err != nil {
			break
		}
		slice = append(slice, a)
		defer f.Close()
	}

	// Numero de filas de coordenadas
	numRows := len(slice) / 3

	// slice vacio
	grid := make([][]float32, numRows)

	//  organiza la estructura del slice de 3 x n
	for i := 0; i < numRows; i++ {
		grid[i] = make([]float32, 3)
		// fmt.Println(grid[i])

	}
	//se llena el slice con las cordenadas
	c := 0
	for i := 0; i < numRows; i++ {
		for j := 0; j < 3; j++ {
			grid[i][j] = slice[c]
			c++
		}
	}
	// / print out slices
	// fmt.Println(grid)

	// Tamaño de la lista de coordenadas
	numeroNodos := len(grid)
	// fmt.Print(numeroNodos)

	// slice vacio
	matrizAdyacencia := make([][]int, numeroNodos)
	matrizAdyacenciaDIST := make([][]int, numeroNodos)
	//  organiza la estructura del slice de 3 x n
	for i := 0; i < numeroNodos; i++ {
		matrizAdyacencia[i] = make([]int, numeroNodos)
		matrizAdyacenciaDIST[i] = make([]int, numeroNodos)
	}

	for i := 0; i < numeroNodos; i++ {
		for j := 0; j < numeroNodos; j++ {
			if i != j {
				a1 := grid[i][1]
				a2 := grid[i][2]
				b1 := grid[j][1]
				b2 := grid[j][2]

				distanciaEuclidiana := int(math.Sqrt(float64((a1-b1)*(a1-b1)+(a2-b2)*(a2-b2))) + 0.5)
				// fmt.Print("   A1 ", a1, " a2 ", a2, " b1 ", b1, " b2 ", b2, " distancia ", distanciaEuclidiana)
				matrizAdyacencia[i][j] = distanciaEuclidiana
				matrizAdyacenciaDIST[i][j] = distanciaEuclidiana
			}
		}
	}

	// //Ejemplo de un solo arranque
	// nodoInicial := 0
	// tourSolucion := insertionTSP(matrizAdyacencia, nodoInicial)
	// fmt.Println("Solución Recibida -> ", tourSolucion)

	inicio := time.Now() //Inicio de la toma de tiempo

	//Multiarranque
	var contenedorInsertions []solucion
	for i := 0; i < len(matrizAdyacencia); i++ {
		contenedorInsertions = append(contenedorInsertions, insertionTSP(matrizAdyacencia, i))
	}

	//Ordenamos los arranques
	sort.SliceStable(contenedorInsertions, func(i, j int) bool {
		return contenedorInsertions[i].costo < contenedorInsertions[j].costo
	})

	//Finalización y presentación de la toma de tiempo
	fmt.Println("tiempo total insertion multiarranque:", time.Since(inicio))

	//Resultado de estrategia multiarranque
	mejorSolucion := contenedorInsertions[0]
	fmt.Println()
	fmt.Println("Mejor Obtenido-> ", mejorSolucion)

}
