package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var total float64

	data, err := os.ReadFile("./prueba.csv")
	if err != nil {
		panic(err)
	}

	res := strings.Split(string(data), "\n")

	for _, line := range res {
		if len(line) > 0 {
			finalLine := strings.Split(string(line), ";")
			fmt.Printf("%v\t\t%v %v\n", finalLine[0], finalLine[1], finalLine[2])

			// if i > 0 {
			// 	price, err := strconv.ParseFloat(finalLine[1], 64)
			// 	quantity, err2 := strconv.ParseFloat(finalLine[2], 64)
			// 	if err != nil || err2 != nil {
			// 		panic(err)
			// 	}
			// 	total += price * quantity
			// }
		}
	}

	fmt.Printf("\t\t %.2f\n", total)
}
