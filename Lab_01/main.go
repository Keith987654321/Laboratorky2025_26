package main

/*
	Идея решения: Складываем все открывающие скобки в голову односвязного списка,
	а когда встречаем закрывающую скобку, проверяем подходят ли они друг дружке.
*/

import "fmt"

type Node struct { //Односвязный список
	Value string
	Next  *Node
}

func (head *Node) addToHead(val string) {
	node := &Node{
		Value: val,
	}
	node.Next = head.Next
	head.Next = node
}

func (head *Node) get() string {
	if head.Next == nil {
		return "stoy, mne ne priyatno!"
	}

	result := head.Next.Value
	head.Next = head.Next.Next
	return result
}

func isValid(s string) bool {
	head := &Node{}
	arr := []byte(s)
	for _, char := range arr {
		if (char == byte('(')) || (char == byte('{')) || (char == byte('[')) {
			head.addToHead(string(char))
			continue
		}
		str := head.get()
		char1 := string(char)
		switch char1 {
		case ")":
			if (str + char1) != "()" {
				return false
			}
		case "}":
			if (str + char1) != "{}" {
				return false
			}
		case "]":
			if (str + char1) != "[]" {
				return false
			}
		}
	}
	if head.Next == nil {
		return true
	}
	return false
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("[{()}]"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("({)"))
	fmt.Println(isValid("({)}"))
	fmt.Println(isValid(""))
}
