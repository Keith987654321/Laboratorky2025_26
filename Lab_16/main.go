package main

import (
	"fmt"
	"math"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func readInput(str string) []int { // Считываем ввод и возвращаем массив чисел для создания бинарного дерева
	result := []int{}
	var number, numberLen int
	for _, ch := range []byte(str) {
		if ch == ' ' || ch == '(' || ch == ')' || ch == ',' {
			if numberLen != 0 {
				result = append(result, number)
				number = 0
				numberLen = 0
			}
			continue
		}
		if '0' <= ch && ch <= '9' {
			number = int(ch-'0') + number*int(math.Pow10(numberLen))
			numberLen++
		}
	}
	return result
}

func (n *Node) insertNode(data int) *Node { // Так же, как и в 15 лабе
	if n == nil {
		n = &Node{}
		n.Data = data
		return n
	}

	if n.Data > data {
		n.Left = n.Left.insertNode(data)
	} else {
		n.Right = n.Right.insertNode(data)
	}
	return n
}

func (n *Node) treeTraversal() {
	stack := []*Node{} // Создаём стек
	stack = append(stack, n)
	for len(stack) > 0 {
		v := stack[len(stack)-1]     // Вытаскиваем эл-т из стека
		stack = stack[:len(stack)-1] // Удаляем этот эл-т из стека

		if v != nil {
			fmt.Printf("%d ", v.Data)      // Печатаем данные из ноды
			stack = append(stack, v.Right) // Добавляем правую ноду
			stack = append(stack, v.Left)  // И левую
		}
	}
}

func main() {
	input := "8 (3 (1, 6 (4,7)), 10 (, 14(13,)))"
	arr := readInput(input)

	var BinaryTree *Node
	for _, num := range arr {
		BinaryTree = BinaryTree.insertNode(num)
	}

	BinaryTree.treeTraversal()
}
