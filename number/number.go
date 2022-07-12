package number

import "fmt"

func Init() {
	DoPercentage()
}

func DoPercentage() {
	result := 1152.730 * 1.2 // (+20%)

	fmt.Printf("%v,%v\n", int(result), result)
}