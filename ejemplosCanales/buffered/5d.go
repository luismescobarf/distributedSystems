package main

//Paquetes requeridos (utilizados)
import (
    "bufio"
    "fmt"
    "log"
    "os"
    "sync"  
)

var libros = map[string]string{
	"Alice" : "libros/AliceWonderland.txt",
	"Iliad" : "libros/Iliad.txt",
	"Quijote" : "libros/DonQuijote.txt",
}

var wg sync.WaitGroup;
var mu sync.Mutex //Protección de palabras
var palabras = make(map[string]int)

//Función de conteo de palabras
func contarPalabras(archivo string){
	
	defer wg.Done()
	
	f,err := os.Open(archivo);
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	
	scanner := bufio.NewScanner(f);
	scanner.Split(bufio.ScanWords);
	
	for scanner.Scan() {
		mu.Lock();
		palabras[scanner.Text()]++
		mu.Unlock();
	}
	if err := scanner.Err(); err != nil {
		log.Println("leyendo input:",err);
	}
	
}

//Sección principal
func main() {
	
	for i:=0; i<100; i++ {
		
		for _,archivo := range libros {
			
			wg.Add(1);
			go contarPalabras(archivo);
			
		}
	}
	
	wg.Wait();
	
	for p, c := range palabras {
		fmt.Printf("palabra: %s, veces: %d\n",p,c);
	}
		
	
}//Final de la sección principal
