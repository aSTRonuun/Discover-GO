package main

import "fmt"

func buy(work1, work2 bool) (bool, bool, bool) {
	buyTv50 := work1 && work2
	buyTv32 := work1 != work2 // exclusive or
	buyIceCream := work1 || work2

	return buyTv50, buyTv32, buyIceCream
}

func main() {
	tv50, tv32, iceCream := buy(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, IceCream: %t, Healthy: %t",
		tv50, tv32, iceCream, !iceCream)
}
