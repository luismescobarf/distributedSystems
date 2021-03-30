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

//Función que cuenta las palabras en un string (línea de un archivo) (mezcla entre map y reduce)
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

//Etapa de mapeo para el conteo
func etapaMap(linea string) []tupla {

	//Limpiar si hay caracteres al comienzo y al final
	linea = strings.TrimSpace(linea)
	//Separar las palabras de la línea o string recibido
	listadoPalabras := strings.Split(linea, " ")

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
func shuffler(contenedorLlaveValor []tupla) map[string][]tupla {

	//Preprar el contenedor de salida del barajado
	apilados := make(map[string][]tupla)

	//Apilar el resultado del mapeo por llaves
	for _, kv := range contenedorLlaveValor {

		//Revisar si ya existe
		_, existe := apilados[kv.llave]
		//Si no existe adicionarlo
		if !existe {
			apilados[kv.llave] = []tupla{kv}
		} else { //Si existe, incrementar la pila correspondiente
			apilados[kv.llave] = append(apilados[kv.llave], kv)
		}

	}

	//Retornar el resultado de apilar comunes
	return apilados

}

func main() {

	var ruta string
	//ruta = "texto.txt"
	ruta = "foo.txt"
	lineas, _ := leerLineas(ruta)

	fmt.Println(lineas)
	//fmt.Println()
	//fmt.Println()
	//fmt.Println(lineas[0])
	//fmt.Println(lineas[1])
	//fmt.Println(lineas[2])

	/*
		fmt.Println("Resultado del mapeo con agregación")
		for i, linea := range lineas {
			fmt.Println("Línea ", i)
			fmt.Println(palabrasString(linea))
		}
	*/

	fmt.Println("Resultado etapa de mapeo")
	var contenedorMapeo []tupla
	for i, linea := range lineas {
		fmt.Println("Línea ", i)
		contenedorMapeo = etapaMap(linea)
		fmt.Println(contenedorMapeo)
	}

	fmt.Println("Resultado etapa de shuffling")
	var contenedorShuffle map[string][]tupla
	contenedorShuffle = shuffler(contenedorMapeo)
	fmt.Println(contenedorShuffle)

}
