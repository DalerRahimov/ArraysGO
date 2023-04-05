package Arrays

import (
	"HW8/Consts"
	"fmt"
	"log"
)

func Arr14() {
	fmt.Println("Введите размер масива:")
	var N int
	var err error
	_, err = fmt.Scan(&N)
	if err != nil {
		log.Println(Consts.Error)
		return
	}

	//Создаём срез и заполняем его
	Arr := make([]int, N)
	fmt.Printf("Заполните масив:\n")
	for i := 0; i < N; i++ {
		fmt.Printf(" Arr[%v] ", i+1)
		_, err := fmt.Scan(&Arr[i])
		if err != nil {
			log.Println(Consts.Error)
			return
		}
	}
	fmt.Println(Arr)

	//Выводим элементы с четными номерами
	for i := 1; i < N; i += 2 {
		fmt.Printf("%v ", Arr[i])

	}
	//Выводим элементы с нечетными номерами
	for i := 0; i < N; i += 2 {
		fmt.Printf("%v ", Arr[i])

	}

}
