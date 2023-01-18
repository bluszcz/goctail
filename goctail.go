package main

import (
	"flag"
	"fmt"
)

var flagvar int

func main() {
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
	flag.Parse()
	// fmt.Println("ip has value ", *ip)
	fmt.Println("flagvar has value ", flagvar)
	fmt.Println("Hello, World!")
}
