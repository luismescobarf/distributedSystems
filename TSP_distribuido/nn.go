package main 

import ("fmt";
		"./tsp")


func main(){

	fmt.Println("Hola mundo");

	var contenedor tsp.Estructura
	fmt.Println(contenedor)

	fmt.Println("Cargando la longitud...")
	contenedor.CargarLongitud(float64(999))

	fmt.Println("Después de cargar la longitud:")
	fmt.Println(contenedor)

	fmt.Printf("Tipo del contenedor: %T\n Valor del contenedor %v\n",contenedor,contenedor)



}