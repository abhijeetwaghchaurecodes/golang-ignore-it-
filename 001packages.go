package main

//simple SYNTAX to import packages , good to interpret , easy to write
import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, Golang! , here's Abhijeet")
	fmt.Printf("and the sq root of 4 is %g", math.Sqrt(4))
}

//println - prints sentences with a NEW LINE character at the end, less customizable ,
//printf - prints with more formatting options , very much CUSTOMIZABLE
