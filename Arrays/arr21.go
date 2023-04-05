package Arrays

import (
	"HW8/Consts"
	"fmt"
	"log"
)

func Arr21() {
	fmt.Println("Введите размер масива:")
	var (
		N, K, L int
	)
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

	//Получаем два индекса массива для вычисления суммы индексов между ними
	fmt.Println("Введите начальный индекс  Arr[i]:")
	_, err = fmt.Scan(&K)
	//fmt.Println(Arr[K])
	if err != nil {
		log.Println(Consts.Error)
		return
	}
	fmt.Println("Введите конечный индекс Arr[i]:")
	_, err = fmt.Scan(&L)
	//fmt.Println(Arr[L])
	if err != nil {
		log.Println(Consts.Error)
		return
	}
	if K >= 0 && L > K && N >= L {
		sum := 0
		j := 0
		for i := K - 1; i <= L-1; i++ {
			sum += Arr[i]
			j++
		}
		var ress float64
		ress = float64(sum) / float64(j)
		fmt.Println(ress)

	} else {
		log.Println(Consts.Error)
		return
	}

}
