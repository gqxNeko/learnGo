package main

import (
	"fmt"
	"unsafe"
	// "internal/itoa"
)

func main() {
	const (
		a = iota
		b = iota
		c
		d = iota
	)
	fmt.Println(a, b, c, d)

	const (
		e = iota
		f
		g
		h = 1
		i
		j = iota
	)
	fmt.Println(e, f, g, h, i, j)

	const (
		item = "hello"
		len  = len(item)
		// cap  = cap(item)
		size = unsafe.Sizeof(item)
	)
	fmt.Println(item, len /* cap */, size)

	const x int = 10
	const y int = 5
	fmt.Println(x, y)

	const (
		l = 0
		m
		n
	)
	fmt.Println(l, m, n)
}
