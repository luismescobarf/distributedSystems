package main

import (	
	"fmt"	
)

//Función para actualizar conjuntos de control en el avance de una heurística
func cubrirNodo(nodo int, nodosCubiertos map[int]struct{}, nodosSinCubrir map[int]struct{}){

	//Agregar al cunjunto de nodos cubiertos
	nodosCubiertos[nodo] = struct{}{}

	//Eliminar del conjunto de nodos cubiertos
	delete(nodosSinCubrir,nodo)

}


//Ejemplo manejo de ítems en un mapa para implementación TSP
func main() {
	
	//Conjuntos	
	nodosSinCubrir := make(map[int]struct{}) 
	nodosCubiertos := make(map[int]struct{}) 

	//Número de nodos del grafo
	var n int = 10

	//Inicializar conjunto de nodos cubiertos
	for i:=0;i<n;i++{
		nodosSinCubrir[i] = struct{}{}
	}

	//Mostrar conjuntos antes de actualización
	fmt.Println("Nodos sin cubrir:")
	fmt.Println(nodosSinCubrir)
	fmt.Println("Nodos cubiertos:")
	fmt.Println(nodosCubiertos)

	//Realizar cobertura de un nodo
	var nodoCubierto int= 4
	cubrirNodo(nodoCubierto,nodosCubiertos,nodosSinCubrir)
	
	//Mostrar conjuntos después de actualización
	fmt.Println("Nodo seleccionado = ",nodoCubierto)
	fmt.Println("Nodos sin cubrir:")
	fmt.Println(nodosSinCubrir)
	fmt.Println("Nodos cubiertos:")
	fmt.Println(nodosCubiertos)

}
