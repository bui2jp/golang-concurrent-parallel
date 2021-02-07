package main

import( 
	"fmt"
	"log"
	"runtime"
   )

func main() {
    fmt.Println("Concurrent Example start ...")

	log.Println(runtime.NumGoroutine())

	//fmt.Println()
	//go myProc(1, 1000000)	
	// go myProc(2, 1000000)	
	// go myProc(3, 1000000)
	// go myProc(4, 1000000)
	// go myProc(5, 1000000)
	// go myProc(6, 1000000)
	// go myProc(7, 1000000)
	// go myProc(8, 1000000)
	// go myProc(9, 1000000)
	myProc(10, 1000000)

    fmt.Println("Concurrent Example finished.")	
}

func myProc(x, y int) int {
	//return x + y
	sum := x
	count := 0
	//for i := 0; i < y; i++ {
	for {
		log.Println("myProc loop", x, ": sum :", sum, " : cnt : ", count)
		sum += x
		count += 1
	}

	return sum
}
