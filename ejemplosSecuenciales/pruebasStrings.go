package main //Nombre del paquete (primer ejemplo de comentario)

//Sección donde se importan otros paquetes o librerías
import (
	"fmt" //Paquete que contiene funciones de entrada y salida básicas
	"strconv"
	"strings"
)

//Definición del main, donde se inicia la ejecición
func main(){

	fmt.Println("Hola mundo");//El uso de ; es opcional

	var cadena string = "prueba"

	fmt.Printf("%s \n",string(cadena[len(cadena)-2]))

	var segundaCadena string = "-->"

	segundaCadena = segundaCadena + "*-->*"

	fmt.Printf("Previo a reemplazos: %s \n", segundaCadena)

	segundaCadena = strings.ReplaceAll(segundaCadena, "*", "")
	segundaCadena = strings.ReplaceAll(segundaCadena, "P", "")

	fmt.Printf("Posterior: %s \n", segundaCadena)

	//Trim(s string, cutset string) string

	//res = strings.ReplaceAll("abcdabxyabr", "a", "")





}//Se utilizan llaves para el manejo de bloques y contextos

/*
 * Complete the timeConversion function below.
 */
 func timeConversion(s string) string {
    /*
     * Write your code here.
     */
     
     //Output string declaration
     var militaryTime string
     
     //Intermediate array to convert the time     
     arrayTime := strings.Split(s, ":")
     
     //Convert the hours to numeric format     
     hh, _ := strconv.ParseInt(arrayTime[0], 0, 16)   
     
     //Get AM/PM
     var am_pm string = s[8] + s[9]
     
     //Depending on the AM/PM move the hours
     
     
     
     
     
     
     
          
     
     //func Split(s, sep string) []string
     
     
     
     //Return the time in militar format
     return militaryTime

}