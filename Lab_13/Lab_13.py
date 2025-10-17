"""
Реализация хэш-таблицы "с наложением" (метод открытой адресации)

Использует обычный список для хранения данных.
При добавлении нового элемента считается его хэш и попытка вставить
в ячейку с таким же индексом. Если она занята — ищем следующую
свободную, пока не найдём или не убедимся, что таблица заполнена.
"""

import string

HASH_TABLE_SIZE = 30


def clean_word(word: str) -> str:
    """Очистка слова от знаков препинания и приведение к нижнему регистру"""
    word = word.strip(string.punctuation + "«»–…")
    return word.lower()


def read_words_from_file(filename: str) -> list[str]:
    """Чтение слов из текстового файла"""
    words = []
    try:
        with open(filename, "r", encoding="utf-8") as file:
            for line in file:
                for raw_word in line.split():
                    word = clean_word(raw_word)
                    if word:
                        words.append(word)
    except FileNotFoundError:
        print(f"Ошибка: не удалось открыть файл {filename}")
        exit(1)
    return words


def get_hash(word: str) -> int:
    """Вычисление хэша слова"""
    return sum(ord(ch) % HASH_TABLE_SIZE for ch in word)


def add_to_hashtable(word: str, table: list[str]) -> None:
    """Добавление слова в хэш-таблицу (линейное пробирование)"""
    h = get_hash(word)
    for i in range(HASH_TABLE_SIZE):
        index = (h + i) % HASH_TABLE_SIZE
        if table[index] == "":
            table[index] = word
            return
    raise RuntimeError("Недостаточно места в таблице")



table = ["" for _ in range(HASH_TABLE_SIZE)]
filename = "input.txt"

words = read_words_from_file(filename)
for word in words:
    try:
        add_to_hashtable(word, table)
    except RuntimeError as e:
        print(e)
        break

# Запись результата
with open("output.txt", "w", encoding="utf-8") as out:
    for i, word in enumerate(table):
        out.write(f"{i}: {word}\n")



