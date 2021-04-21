package main

import (
	"bufio"
	"fmt"
	"os"
)

type nodo struct {
	id int
	x  float32
	y  float32
}

type arista struct {
	id string
	i  nodo
	j  nodo
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

//Sección principal
func main() {

	fmt.Println("Carga con la información de la instancia")

	//Llamado a la función de lectura del archivo
	var ruta string
	ruta = "./eil51.tsp"
	lineas, _ := leerLineas(ruta)

	fmt.Println("Líneas cargadas", lineas)

	fmt.Println("Salida con líneas diferenciadas")
	for i, linea := range lineas {
		fmt.Println(i, ") ", linea)
	}

	fmt.Printf("Tipo del contenedor generado -> %T\n", lineas)

}
