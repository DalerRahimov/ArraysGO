package Arrays

import (
	"HW8/Consts"
	"fmt"
	"log"
	"sort"
)

func Arr142() {

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

	var ArrEven []int //Срез для четных номеров для сортировки
	var ArrOdd []int  //Срез с четными номеров для сортировки

	//Выводим элементы с четными номерами
	q := 0
	for i := 1; i < N; i += 2 {
		fmt.Printf("%v ", Arr[i])
		//Заполняем четный срез
		ArrEven = append(ArrEven, Arr[i])
		q++
	}
	//Сортируем срез с четными
	sort.Slice(ArrEven, func(i, j int) bool {
		return ArrEven[i] < ArrEven[j]
	})
	fmt.Println(ArrEven)

	//Выводим элементы с нечетными номерами
	j := 0
	for i := 0; i < N; i += 2 {
		fmt.Printf("%v ", Arr[i])
		//Заполняем нечетный срез
		ArrOdd = append(ArrOdd, Arr[i])
		j++
	}
	//Сортируем срез с нечетными
	sort.Slice(ArrOdd, func(i, j int) bool {
		return ArrOdd[i] < ArrOdd[j]
	})
	fmt.Println(ArrOdd)

}
