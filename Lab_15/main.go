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

func (n *Node) insertNode(data int) *Node { // Рекурсивно добавляем новую ноду в бинарное дерево
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

func (n *Node) preorderTraversal() { // Распечатывание бинарного дерева "прямым" обходом
	if n == nil { // Проверяем, не является ли эта нода пустой
		return
	}

	fmt.Printf("%d ", n.Data)   // Сначала печатаем содержимое ноды
	n.Left.preorderTraversal()  // Потом для левой ветки
	n.Right.preorderTraversal() // И для правой
}

func (n *Node) inorderTraversal() { // Распечатывание дерева "центральным" обходом
	if n == nil {
		return
	}

	n.Left.inorderTraversal()  // Печатаем левую ветку
	fmt.Printf("%d ", n.Data)  // Печатаем содержимое этой ноды уже после распечатывания левой ветки
	n.Right.inorderTraversal() // Печатаем правую
}

func (n *Node) postorderTraversal() { // Распечатываем дерево "концевым" обходом
	if n == nil {
		return
	}

	n.Left.postorderTraversal()
	n.Right.postorderTraversal()
	fmt.Printf("%d ", n.Data)
}

func main() {
	input := "8 (3 (1, 6 (4,7)), 10 (, 14(13,)))"

	arr := readInput(input)
	var BinaryTree *Node
	for _, num := range arr {
		BinaryTree = BinaryTree.insertNode(num)
	}

	BinaryTree.preorderTraversal()
	fmt.Println()
	BinaryTree.inorderTraversal()
	fmt.Println()
	BinaryTree.postorderTraversal()
	fmt.Println()
}
