package main

/*
	Работает в целом так же, как и 13 лаба, но
	вместо пробирования используем вложенные списки.
	В каждой ячейке хранится массив эл-тов string.
	После того как мы посчитали hash, добавляем
	это слово в список под индексом hash.
*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

const TableSize = 30 // Всего 30 ячеек, где в каждой находится пока что пустой массив

func clearWord(word string) string {
	return strings.ToLower(strings.TrimFunc(word, func(r rune) bool {
		return !unicode.IsLetter(r)
	}))
}

func readWordsFromFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	words := []string{}
	for scanner.Scan() {
		word := clearWord(scanner.Text())
		if word == "" {
			continue
		}
		words = append(words, word)
	}

	return words, nil
}

func getHash(word string) int {
	hash := 0
	for _, r := range []rune(word) {
		hash += int(r)
	}
	return hash % TableSize
}

func addToHashtable(word string, table [][]string) {
	hash := getHash(word)
	table[hash] = append(table[hash], word) // Добавляем слово в массив под индексом hash
}

func main() {
	words, err := readWordsFromFile("input.txt")
	if err != nil {
		log.Fatalln("Can't open the file: input.txt")
	}

	hashTable := make([][]string, TableSize)
	for _, word := range words {
		addToHashtable(word, hashTable)
	}

	out, _ := os.Create("output.txt")

	counter := 1
	for _, words := range hashTable {
		for _, word := range words {
			fmt.Fprintf(out, "%d: %s\n", counter, word)
			counter++
		}
	}
}
