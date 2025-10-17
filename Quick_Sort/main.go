package main

/*
	Быстрая сортировка берёт эл-т по середине массива, как опорный,
	и сравнивает остальные эл-ты с ним. По окончании работы функции
	partOfQuickSort все эл-ты слева от опорного должны быть <= ему,
	а все эл-ты справа >= опорного. Таким образом мы разбили массив
	на 2 подмассива - меньшый (левый) и больший (правый). Потом
	с помощью рекурсии мы запускаем тот же алгоритм уже отдельно
	для каждого подмассива и алгоритм повторяется, пока длина подмассивов
	не станет равна 1. Когда длины всех подмассивов = 1, то исходный
	массив считается отсортированным.
*/

import "fmt"

func partOfQuickSort(arr []int, left, right int) int {
	pivot := arr[(left+right)/2] // Запоминаем опорный эл-т
	for left <= right {          // Сортируем подмассивы, пока указатели left и right не пересекутся друг с другом
		for ; arr[left] < pivot; left++ { // Двигаем указатель left вправо, пока не найдём эл-т больше опорного
		}
		for ; arr[right] > pivot; right-- { // Двигаем right, пока не найдём эл-т меньше опорного
		}

		if left <= right {
			arr[left], arr[right] = arr[right], arr[left] // Меняем эл-ты в нужные подмассивы
			left++
			right--
		}
	}
	return left // По окончанию ф-ции указатель left указывает на начало правого (большего) подмассива
}

func QuickSort(arr []int, start, end int) {
	if start >= end { // Условие выхода из рекурсии
		return
	}

	rightStart := partOfQuickSort(arr, start, end) // Делим массив на 2 подмассива
	QuickSort(arr, start, rightStart-1)            // Рекурсивно сортируем левый подмассив
	QuickSort(arr, rightStart, end)                // и правый
}

func Sort(arr []int) {
	QuickSort(arr, 0, len(arr)-1)
}

func main() {
	arr := []int{4, 7, 2, 5, 9, 4, 3}

	Sort(arr)

	fmt.Print(arr)
}
