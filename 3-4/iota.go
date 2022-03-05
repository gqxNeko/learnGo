package main

import "fmt"

func main() {
	const (
		a = iota
		b = 1
		c
		d
		e = "hello world"
		f = 100
		g
		h
	)
	var url = "a=%d\nb=%d\nc=%d\nd=%d\ne=%s\nf=%d\ng=%d\nh=%d"
	var targetItem = fmt.Sprintf(url, a, b, c, d, e, f, g, h)
	fmt.Println(targetItem)

	const (
		i = 1 << iota
		j = 3 << iota
		k
		l
	)
	fmt.Println(i,j,k,l)
}
