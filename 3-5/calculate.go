package main

import "fmt"

func main() {
	var a uint = 60 // 60 = 32 + 16 + 8 + 4  	0011 1100
	var b uint = 13 // 13 = 8 + 4 + 1 				0000 1101
	var c = a & b 	// 4 + 8									0000 1100
	var d = a ^ b 	// 1 +	16 + 32						0011 0001
	var e = a | b		// 1 + 4 + 8	+ 16 + 32		0011 1101
	fmt.Println(a, b, c, d, e) // 60 13 12 49 61
}
