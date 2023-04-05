package Arrays

import (
	"fmt"
	"log"
)

func Arr2() {

	fmt.Println("Введите размер масива:")
	var arr []int
	var N int
	var err error
	_, err = fmt.Scan(&N)
	if err != nil {
		log.Println("Input error")
		return
	}
	if N <= 0 {
		log.Println("Input error")
		return
	} else {

	}
	for i := 2; len(arr) <= N; i *= 2 {
		arr = append(arr, i)
	}
	fmt.Println(arr)

}
