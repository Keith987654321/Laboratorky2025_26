package main

import "fmt"

/*
	Работает аналогично сортировке вставками, но сравниваются
	эл-ты не последовательно, а на промежутке (gap),
	блягодаря чему большие эл-ты массива быстрее
	встают на свои места.
	Сложность в худшем случае O(n^2)
	Сложность по памяти O(1)
*/

func ShellSort(arr []int) []int {

	for gap := len(arr) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(arr); i++ {
			number := arr[i]
			j := i
			for j >= gap && arr[j-gap] > number {
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = number
		}
	}
	return arr
}

func main() {
	arr := []int{8, 4, 1, 56, 3, -44, 23, -6, 28, 0}
	fmt.Print(ShellSort(arr))
}
