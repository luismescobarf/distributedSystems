package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	dat, err := ioutil.ReadFile("./eil51.tsp")
	if err != nil {
		panic(err)
	}
	//fmt.Print(string(dat))
	fmt.Printf("%T", dat)
	fmt.Println()
	var coordenadasPuntos [][]int
	lineas := strings.Split(string(dat), "\n")
	for _, s := range lineas {
		coordenadasTexto := strings.Split(s, " ")
		var lineaCoordenadas []int
		for i := 0; i < len(coordenadasTexto); i++ {
			cn, _ := strconv.Atoi(coordenadasTexto[i])
			lineaCoordenadas = append(lineaCoordenadas, cn)
		}
		coordenadasPuntos = append(coordenadasPuntos, lineaCoordenadas)
	}

	fmt.Println(coordenadasPuntos)

}
