'''
    Сортировка выбором среди эл-тов дальше i-ого
	ищет минимальный и если он меньше i-ого меняет
	их местами. Дальше за i-ый эл-т берется эл-т
	справа от него и так же ищется минимальный эл-т
	дальше от i-ого. И так пока i не станет равен длине массива.
    * Сложность O(n^2)
    * Сложность по памяти O(1)
'''


def selection_sort(arr):
    length = len(arr)

    for i in range(length):
        smallest_index = i # Далее будем сравнивать с i-ым эл-том

        for j in range(i, length):
            if arr[j] < arr[smallest_index]:
                smallest_index = j # запоминаем индекс нового минимального эл-та
                
        arr[i], arr[smallest_index] = arr[smallest_index], arr[i] # Меняем местами i-ый эл-т и минимальный

    return arr


arr = [5, 4, 9, 7, 6, 1]
print(selection_sort(arr))
