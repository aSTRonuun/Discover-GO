package main

import (
	"fmt"
	"time"

	"github.com/aSTRonuun/html"
)

func theFastet(url1, url2, url3 string) string {
	c1 := html.Title(url1)
	c2 := html.Title(url2)
	c3 := html.Title(url3)

	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(500 * time.Millisecond):
		return "Overyone Lose!"
		// default:
		// 	return "No anwsers yet..."
	}
}

func main() {

	champion := theFastet(
		"https://www.cod3r.com.br",
		"https://youtube.com",
		"https://www.google.com",
	)

	fmt.Println(champion)
}
