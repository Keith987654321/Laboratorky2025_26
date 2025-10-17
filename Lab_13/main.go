package main

/*
	Реализация хэш-таблицы "с наложением" (метод открытой адресации)
	испошльзует для хранения данных обычный массив данных, при добавлении
	нового эл-та считает его хэш и пытается добавить этот эл-т в ячейку
	с таким же номером, как и хэш. Если эта ячейка уже занята, то
	пытаемся вставить в следующую ячейку и так пока либо не найдём
	свободную ячейку, либо не поймём, что свободных мест нет.
*/

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

const HashTableSize = 30

func cleanWord(word string) string { // Чистим слова от знаков препинания
	return strings.ToLower(strings.TrimFunc(word, func(r rune) bool {
		return !unicode.IsLetter(r)
	}))
}

func readWordsFromFile(fileName string) ([]string, error) { // Читаем слова из файла и записываем в файл
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	words := []string{}

	for scanner.Scan() {
		word := cleanWord(scanner.Text())
		if word == "" {
			continue
		}
		words = append(words, word)
	}

	return words, nil
}

func getHash(word string) int { // Находим хэш слова
	hash := 0
	for _, r := range word {
		hash += int(r) % HashTableSize
	}
	return hash
}

func addToHashtable(word string, table []string) error { // Добавляем слово в хэш-таблицу
	hash := getHash(word)
	for i := 0; i < HashTableSize; i++ {
		if table[(hash+i)%HashTableSize] == "" {
			table[(hash+i)%HashTableSize] = word
			return nil
		}
	}
	return errors.New("Not enough space in table")
}

func main() {
	table := make([]string, HashTableSize)
	fileName := "input.txt"

	text, err := readWordsFromFile(fileName)
	if err != nil {
		log.Fatalf("Can't open the file: %s\n", fileName)
	}

	for _, word := range text {
		err := addToHashtable(word, table)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
	}

	out, _ := os.Create("output.txt")
	defer out.Close()

	for i, word := range table {
		fmt.Fprintf(out, "%d: %s\n", i, word)
	}
}
