package main

import (
	"fmt"
	"time"
)

func finder(theMine []string) []string {
	var foundOre []string
	for _, p := range theMine {
		if p == "ore1" || p == "ore2" || p == "ore3" {
			foundOre = append(foundOre, p)
		}
	}
	fmt.Println("Encontrado por el explorador: ", foundOre)
	return foundOre
}

func miner(foundOre []string) []string {
	var minedOre []string
	for i := 0; i < len(foundOre); i++ {
		minedOre = append(minedOre, "minedOre")
	}
	fmt.Println("Extraído por el minero: ", minedOre)
	return minedOre
}

func smelter(minedOre []string) []string {
	var smeltedOre []string
	for i := 0; i < len(minedOre); i++ {
		smeltedOre = append(smeltedOre, "smeltedOre")
	}
	fmt.Println("Fundido por el metalúrgico: ", smeltedOre)
	return smeltedOre
}

// //Secuencial (Single-Thread)
// //////////////////////////
// func main() {
// 	theMine := []string{"rock", "ore", "ore", "rock", "ore"}
// 	foundOre := finder(theMine)
// 	minedOre := miner(foundOre)
// 	fmt.Println(smelter(minedOre))
// }

// //Concurrencia sin coordinación (sólo explorador) (Multi-Thread)
// ////////////////////////////////////////////////////////////////
// func main() {
// 	theMine := []string{"rock", "ore1", "ore2", "rock", "ore3"}
// 	go finder(theMine)
// 	go finder(theMine)
// 	<-time.After(time.Second * 5) //you can ignore this for now
// }

// //Concurrencia coordinación explorador y extractor o minero
// ///////////////////////////////////////////////////////////
// func main() {
// 	theMine := []string{"rock", "ore1", "ore2", "rock", "ore3"}
// 	oreChan := make(chan string)

// 	//Finder
// 	go func(mine []string) {
// 		for _, item := range mine {
// 			if item == "ore1" || item == "ore2" || item == "ore3" {
// 				oreChan <- item //send
// 			}
// 		}
// 	}(theMine)

// 	//Ore Breaker
// 	go func() {
// 		for i := 0; i < 3; i++ {
// 			foundOre := <-oreChan //receive
// 			fmt.Println("Miner: Received " + foundOre + " from finder")
// 		}
// 	}()
// 	<-time.After(time.Second * 5) // Again, ignore this for now
// }

// //Buffered channels
// func main() {

// 	//Ejemplo funcionamiento cola buffered channel
// 	bufferedChan := make(chan string, 3)
// 	go func() {
// 		bufferedChan <- "first"
// 		fmt.Println("Sent 1st")
// 		bufferedChan <- "second"
// 		fmt.Println("Sent 2nd")
// 		bufferedChan <- "third"
// 		fmt.Println("Sent 3rd")
// 	}()

// 	<-time.After(time.Second * 1)

// 	go func() {
// 		firstRead := <-bufferedChan
// 		fmt.Println("Receiving..")
// 		fmt.Println(firstRead)
// 		secondRead := <-bufferedChan
// 		fmt.Println(secondRead)
// 		thirdRead := <-bufferedChan
// 		fmt.Println(thirdRead)
// 	}()

// }

//Minería concurrente (atando todos los cabos)

func main() {

	theMine := []string{"rock", "ore1", "ore2", "rock", "ore3"}
	oreChannel := make(chan string)
	minedOreChan := make(chan string)

	// Finder
	go func(mine []string) {
		for _, item := range mine {
			if item == "ore1" || item == "ore2" || item == "ore3" {
				oreChannel <- item //send item on oreChannel
			}
		}
	}(theMine)

	// Ore Breaker
	go func() {
		for foundOre := range oreChannel {
			fmt.Println("From Finder: ", foundOre)
			minedOreChan <- foundOre + "_minedOre" //send to minedOreChan
		}
	}()

	// Smelter
	go func() {
		for minedOre := range minedOreChan {
			fmt.Println("From Miner: ", minedOre)
			fmt.Println("From Smelter: Ore is smelted: " + minedOre + "smeltedOre")
		}
	}()

	<-time.After(time.Second * 5) // Again, you can ignore this

}
