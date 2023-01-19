package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')

	text = strings.ReplaceAll(text, " ", "")

	columns := math.Ceil(math.Sqrt(float64(len(text))))

	for i := 0; i < int(columns); i++ {
		for j := i; j < len(text); j += int(columns) {
			fmt.Printf("%c", text[j])
		}
		fmt.Printf(" ")
	}

}
