package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type tupla struct {
	llave string
	valor int
}

//Función que cuenta las palabras en un string (línea de un archivo)
//Nota: mezcla entre map y reduce
func palabrasString(linea string) map[string]int {

	//Inicializar contenedor de salida
	mapaPalabras := make(map[string]int)

	//Limpiar si hay caracteres al comienzo y al final
	linea = strings.TrimSpace(linea)

	//Separar las palabras de la línea o string recibido
	listadoPalabras := strings.Split(linea, " ")

	//Efectuar el conteo
	for _, palabra := range listadoPalabras {
		mapaPalabras[palabra] += 1
	}

	//Retornar el mapa con el procesamiento respectivo
	return mapaPalabras
}

//Función para lectura de un archivo
func leerLineas(path string) ([]string, error) {
	archivo, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer archivo.Close()

	//Arreglo que recibirá las líneas del archivo
	var lineas []string
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		lineas = append(lineas, scanner.Text())
	}

	//Retornar las líneas y el estado de la lectura
	return lineas, scanner.Err()
}

//Etapa de mapeo para la porción de trabajo recibida
func etapaMap(listadoPalabras []string) []tupla {

	/*
		//Limpiar si hay caracteres al comienzo y al final
		linea = strings.TrimSpace(linea)
		//Separar las palabras de la línea o string recibido
		listadoPalabras := strings.Split(linea, " ")
	*/

	//Contenedor con el resultado del mapeo
	contenedorLlaveValor := make([]tupla, len(listadoPalabras))

	//Realizar el mapeo
	for i, palabra := range listadoPalabras {
		contenedorLlaveValor[i] = tupla{llave: palabra, valor: 1}
	}

	//Retornar el resultado del mapeo (trabajo del mapper)
	return contenedorLlaveValor
}

//Etapa de "barajado" del mapeo
func shuffler(contenedorLlaveValor [][]tupla) map[string][]tupla {

	//Preprar el contenedor de salida del barajado
	apilados := make(map[string][]tupla)

	//Por cada una de las líneas o pedazos resultantes de mapeo
	for _, linea := range contenedorLlaveValor {

		//Apilar el resultado del mapeo por llaves
		for _, kv := range linea {

			apilados[kv.llave] = append(apilados[kv.llave], kv)

		}

	}

	//Retornar el resultado de apilar comunes
	return apilados

}

//Etapa reduce
func reducer(contenedorShuffler map[string][]tupla) map[string]int {

	//Preparar el contenedor de reducción
	contenedorReducer := make(map[string]int)

	//Reducir cada una de las pilas o montones
	for _, monton := range contenedorShuffler {
		for _, kv := range monton {
			contenedorReducer[kv.llave] += kv.valor
		}
	}

	//Retornar el resultado de la reducción
	return contenedorReducer

}

//El main realizará en la versión secuencial el Input, el Splitting y el Final result
func main() {

	//Input y splitting (distribución del trabajo y limpieza del input)

	//Splitting implícito por líneas
	var ruta string
	ruta = "texto.txt"
	//ruta = "foo.txt"
	lineas, _ := leerLineas(ruta)
	fmt.Println(lineas)

	//Tratamiento del input para enviar las porciones al flujo del mapReduce
	var acumuladorPartesTrabajo [][]string
	for _, linea := range lineas {

		//Limpiar si hay caracteres de movimientos de carro al comienzo y al final
		linea = strings.TrimSpace(linea)
		//Quitar puntuaciones
		linea = strings.ReplaceAll(linea, ":", "")
		linea = strings.ReplaceAll(linea, ";", "")
		linea = strings.ReplaceAll(linea, ",", "")
		linea = strings.ReplaceAll(linea, ".", "")
		//Separar las palabras de la línea o string recibido
		listadoPalabras := strings.Split(linea, " ")
		acumuladorPartesTrabajo = append(acumuladorPartesTrabajo, listadoPalabras)

	}

	/*
		//Utilización de función mezclada, no hace parte de la estructura original mapReduce
		fmt.Println("Resultado del mapeo con agregación")
		for i, linea := range lineas {
			fmt.Println("Línea ", i)
			fmt.Println(palabrasString(linea))
		}
	*/

	fmt.Println("---->Resultado etapa de mapeo")
	var contenedorMapeo [][]tupla
	for i, parteTrabajo := range acumuladorPartesTrabajo {
		fmt.Println("Parte ", i)
		contenedorMapeo = append(contenedorMapeo, etapaMap(parteTrabajo))
		fmt.Println(contenedorMapeo)
	}

	fmt.Println("---->Resultado etapa de shuffling")
	var contenedorShuffle map[string][]tupla
	contenedorShuffle = shuffler(contenedorMapeo)
	fmt.Println(contenedorShuffle)

	fmt.Println("---->Resultado etapa reduce")
	var contenedorReducer map[string]int
	contenedorReducer = reducer(contenedorShuffle)
	fmt.Println(contenedorReducer)

	fmt.Println("---->Resultado final")
	for llave, valor := range contenedorReducer {
		fmt.Printf("%s, %d\n", llave, valor)
	}

}
