"""
Реализация хэш-таблицы "со списками" (метод цепочек)

Работает аналогично лабе №13, но вместо пробирования
используются вложенные списки: каждая ячейка таблицы
— это список строк. После вычисления hash слово
добавляется в соответствующий список.
"""

import string

TABLE_SIZE = 30  # Всего 30 ячеек (каждая — список слов)


def clear_word(word: str) -> str:
    """Очистка слова от пунктуации и приведение к нижнему регистру"""
    return word.strip(string.punctuation + "«»–…").lower()


def read_words_from_file(filename: str) -> list[str]:
    """Чтение слов из текстового файла"""
    words = []
    try:
        with open(filename, "r", encoding="utf-8") as file:
            for line in file:
                for raw_word in line.split():
                    word = clear_word(raw_word)
                    if word:
                        words.append(word)
    except FileNotFoundError:
        print(f"Ошибка: не удалось открыть файл {filename}")
        exit(1)
    return words


def get_hash(word: str) -> int:
    """Вычисление хэша слова"""
    return sum(ord(ch) for ch in word) % TABLE_SIZE


def add_to_hashtable(word: str, table: list[list[str]]) -> None:
    """Добавление слова в хэш-таблицу (метод цепочек)"""
    h = get_hash(word)
    table[h].append(word)



words = read_words_from_file("input.txt")

# Создаём таблицу из 30 пустых списков
hash_table = [[] for _ in range(TABLE_SIZE)]

# Добавляем слова
for word in words:
    add_to_hashtable(word, hash_table)

# Записываем результат в файл
with open("output.txt", "w", encoding="utf-8") as out:
    counter = 1
    for bucket in hash_table:
        for word in bucket:
            out.write(f"{counter}: {word}\n")
            counter += 1



