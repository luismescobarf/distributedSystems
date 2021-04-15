package main

//Paquetes requeridos (utilizados)
import (
    "fmt"
    "time"  
)

//Estructura Carrito de Compras
type carrito struct{
	subtotal 	float64;
	impuestos 	float64;
	envio 		float64;
	total 		float64;
}

//Función para calcular los cargos asociados a una venta
func calcCargos(entrada <-chan *carrito, salida chan<- *carrito){
	for c:= range entrada {
		time.Sleep(3 * time.Second);//Simular retrasos para realizar la operación
		c.impuestos = c.subtotal * 0.075;//7.5% de impuestos
		c.envio = c.subtotal * 0.10;//10% para los gastos de envío
		salida <- c;
	}
	close(salida);	
}


//Función para calcular el total del carrito de compras
func calcTotal(entrada <-chan  *carrito, salida chan<- *carrito){
	for c := range entrada{
		time.Sleep(3 * time.Second);//Simular retrasos para realizar la operación
		c.total = c.subtotal + c.impuestos + c.envio;
		salida <- c;
	}
	close(salida);
}


//Data Race - Simulación de carrito de compras
func main() {
	
	//Subtotales de cada uno de los carritos de compra que podrían generarse en el sistema distribuído
	subtotales := []float64{100.00,110.00,50.00,55.00};//Simulación de una entrada
	
	//Buffered channels
	carritos := make(chan *carrito, len(subtotales));//Espacio necesario para crear los carritos
	conCargos := make(chan *carrito, len(subtotales));
	conTotal := make(chan *carrito, len(subtotales));
	
	go calcCargos(carritos, conCargos);
	go calcTotal(conCargos, conTotal);
	
	
	//Recorrer cada uno de los carritos subtotalizados
	for _, st := range subtotales{
		
		carritos <- &carrito{subtotal: st}
			
	}//Fin del recorrido de los carritos
	close(carritos);
	
	//Recorrer el canal con los carritos totalizados (cargos y envíos)
	for c := range conTotal {
		
		//Mostrar en pantalla lo que se obtiene
		fmt.Printf("\n %#v \n",c)
		
	}	
	
}//Final de la sección principal
