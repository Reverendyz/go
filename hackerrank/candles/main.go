package main

import "fmt"

func main() {
	candles := 0

	fmt.Scanf("%d", &candles)

	higher := 0
	count := 0

	for i := 0; i < candles; i++ {
		var nextCandle int
		fmt.Scanf("%d", &nextCandle)

		if nextCandle > higher {
			higher = nextCandle
			count = 0
		}
		if nextCandle == higher {
			count += 1
		}
	}
	fmt.Println(count)
}
