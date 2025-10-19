package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func input() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	return reader.ReadString('\n')
}

func parseInput(str string) []int {
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

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func insertNode(root *Node, value int) *Node {
	if root == nil {
		root = &Node{}
		root.Data = value
	} else if root.Data > value {
		root.Left = insertNode(root.Left, value)
	} else {
		root.Right = insertNode(root.Right, value)
	}
	return root
}

func getMax(n *Node) *Node {
	if n.Right == nil {
		return n
	}
	return getMax(n.Right)
}

func findNode(n *Node, value int) *Node {
	if n == nil {
		return nil
	}

	if n.Data > value {
		return findNode(n.Left, value)
	}

	if n.Data < value {
		return findNode(n.Right, value)
	}

	return n
}

func delete(n *Node, value int) *Node {
	if n == nil {
		return nil
	} else if value < n.Data {
		n.Left = delete(n.Left, value)
	} else if value > n.Data {
		n.Right = delete(n.Right, value)
	} else {
		if n.Left == nil && n.Right != nil {
			n = n.Right
		} else if n.Left != nil && n.Right == nil {
			n = n.Left
		} else if n.Left == nil && n.Right == nil {
			n = nil
		} else {
			leftMax := getMax(n.Left)
			n.Data = leftMax.Data
			n.Left = delete(n.Left, leftMax.Data)
		}
	}
	return n
}

func (n *Node) toBracketNotation() string { // Распечатывание бинарного дерева линейно-скобочным способом
	if n == nil {
		return ""
	}

	if n.Left == nil && n.Right == nil {
		return fmt.Sprintf("%d", n.Data)
	}

	if n.Left != nil && n.Right == nil {
		return fmt.Sprintf("%d (%s,)", n.Data, n.Left.toBracketNotation())
	}

	if n.Left == nil && n.Right != nil {
		return fmt.Sprintf("%d (,%s)", n.Data, n.Right.toBracketNotation())
	}

	return fmt.Sprintf("%d (%s,%s)", n.Data, n.Left.toBracketNotation(), n.Right.toBracketNotation())

}

func printMenu(tree *Node) {
	Continue := true
	var n int
	for Continue {
		fmt.Print("1 - Insert a number\n2 - Delete a number\n3 - Search a number\n4 - Print tree\n5 - Close\nEnter an operation number: ")
		fmt.Scan(&n)
		switch n {
		case 1:
			var value int
			fmt.Print("Enter a number to insert: ")
			fmt.Scan(&value)
			insertNode(tree, value)
		case 2:
			var value int
			fmt.Print("Enter a number to insert: ")
			fmt.Scan(&value)
			delete(tree, value)
		case 3:
			var value int
			fmt.Print("Enter a number to insert: ")
			fmt.Scan(&value)
			if findNode(tree, value) == nil {
				fmt.Println("Can't find this number")
			} else {
				fmt.Println("Number successfully found")
			}
		case 4:
			fmt.Println(tree.toBracketNotation())
		case 5:
			fmt.Println(tree.toBracketNotation())
			Continue = false
		}
	}
}

func main() {
	fmt.Print("Enter a binary tree structure: ")

	input, err := input()
	if err != nil {
		log.Fatalln(err)
	}

	numbers := parseInput(input)
	var binaryTree *Node
	for _, n := range numbers {
		binaryTree = insertNode(binaryTree, n)
	}

	printMenu(binaryTree)
}
