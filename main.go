package main

import (
	"HW8/Arrays"
	"HW8/Consts"
	"fmt"
)

func main() {

	fmt.Println("Выберите задачу: 1)Arr2 2)Arr14 3)Arr14.2 4)Arr21")
	var ex int
	fmt.Scan(&ex)
	switch ex {
	case 1:
		Arrays.Arr2()
	case 2:
		Arrays.Arr14()
	case 3:
		Arrays.Arr142()
	case 4:
		Arrays.Arr21()
	default:
		fmt.Println(Consts.Error)

	}
}
