package main

/*
	Сортировка выбором среди эл-тов дальше i-ого
	ищет минимальный и если он меньше i-ого меняет
	их местами. Дальше за i-ый эл-т берется эл-т
	справа от него и так же ищется минимальный эл-т
	дальше от i-ого. И так пока i не станет равен длине массива.
*/

import "fmt"

func SelectionSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		smallestIndex := i // Далее будем сравнивать с i-ым эл-том
		for j := i; j < length; j++ {
			if arr[j] < arr[smallestIndex] {
				smallestIndex = j // запоминаем индекс нового минимального эл-та
			}
		}
		arr[i], arr[smallestIndex] = arr[smallestIndex], arr[i] // Меняем местами i-ый эл-т и минимальный
	}
	return arr
}

func main() {
	arr := []int{5, 4, 9, 7, 6, 1}
	fmt.Print(SelectionSort(arr))
}
