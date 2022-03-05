package main

import "fmt"

func main() {
	const (
		a = iota
		b = 1
		c
		d
		e = "hello world"
		f 
		g = 100
		h
	)
	var url = "a=%d\nb=%d\nc=%d\nd=%d\ne=%s\nf=%s\ng=%d\nh=%d"
	var targetItem = fmt.Sprintf(url, a, b, c, d, e, f, g, h)
	fmt.Println(targetItem)

	const (
		i = 1 << iota // iota = 0
		j = 3 << iota // 1 ; iota += 1
		k 						// 3<<iota 2 ; iota += 1
		l 						// 3<<iota 3 ; iota += 1
	)
	fmt.Println(i,j,k,l)
}
