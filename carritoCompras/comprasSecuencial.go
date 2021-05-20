package main

//Paquetes requeridos (utilizados)
import (
	"fmt"
	"time"
)

//Estructura Carrito de Compras
type carrito struct {
	subtotal  float64
	impuestos float64
	envio     float64
	total     float64
}

//Función para calcular los cargos asociados a una venta
func calcCargos(c *carrito) {
	time.Sleep(3 * time.Second)      //Simular retrasos para realizar la operación
	c.impuestos = c.subtotal * 0.075 //7.5% de impuestos
	c.envio = c.subtotal * 0.10      //10% para los gastos de envío
}

//Función para calcular el total del carrito de compras
func calcTotal(c *carrito) {
	time.Sleep(3 * time.Second) //Simular retrasos para realizar la operación
	c.total = c.subtotal + c.impuestos + c.envio
}

//Data Race - Simulación de carrito de compras
func main() {

	//Subtotales de cada uno de los carritos de compra que podrían generarse en el sistema distribuído
	subtotales := []float64{100.00, 110.00, 50.00, 55.00} //Simulación de una entrada

	//Recorrer cada uno de los carritos subtotalizados
	for _, st := range subtotales {

		//Instanciar cada uno de los carritos de compra
		c := carrito{subtotal: st} //Constructor que brinda Go

		//Calcular los cargos de la compra registrada en el carrito
		calcCargos(&c)

		//Calcular el valor totar a pagar por la compra
		calcTotal(&c)

		//Mostrar en pantalla lo que se obtiene
		fmt.Printf("\n %#v \n", c)

	} //Fin del recorrido de los carritos

} //Final de la sección principal
