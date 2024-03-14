package main

import "fmt"

func zoomIn(s int, w, h *int) {
	*w = *w * s
	*h = *h * s
}

func zoomOut(s int, w, h *int) {
	*w = *w / s
	*h = *h / s
}

func main() {
	var w, h, s int
	fmt.Scan(&w, &h)
	var t rune
	fmt.Scanf("%c %d", &t, &s)
	if t == '+' {
		zoomIn(s, &w, &h)
	} else if t == '-' {
		zoomOut(s, &w, &h)
	}
	fmt.Println(w, h)
}
