package main

import (
	"a_alg"
	"log"
)

func main() {
	err, road := a_alg.SearchStandardAStar2D(
			". . . . # . . . . . . . # . . . . . .\n" +
			". . O . # . . . . . . . # . . . # . .\n" +
			". # # . # . . . . . . . # . . . # . .\n" +
			". . # # # . . . . . . . # . . . # . .\n" +
			". . . # . . # . . . . . # . . . # . .\n" +
			". . . . . . . . . # . . # . . . # . .\n" +
			". . . # . . . # . . # . # . . . # . .\n" +
			". . . # # # # # # . # . # . . . # . .\n" +
			". . . # . . . . . . # . # . . . # . .\n" +
			". . . # . . . . # # # . # . . . # . .\n" +
			". . . # # # # . . # . . # . . . # . .\n" +
			". # . . . . # . . . . . . . . . # . .\n" +
			". . . # # # . # . # . # # # # # # . .\n" +
			". # . . . # . # . . . # . . . # . . .\n" +
			". . # # . . . . # # # # . . X # . . .\n" +
			". # . . . . . . . . . # . . . . . . .\n")
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("\n%v", road)
	}
}
