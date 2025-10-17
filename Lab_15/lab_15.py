import math

class Node:
    def __init__(self, data):
        self.Data = data
        self.Left = None
        self.Right = None

def readInput(s):  # Считываем ввод и возвращаем массив чисел для создания бинарного дерева
    result = []
    number = 0
    numberLen = 0
    for ch in s:
        if ch in [' ', '(', ')', ',']:
            if numberLen != 0:
                result.append(number)
                number = 0
                numberLen = 0
            continue
        if '0' <= ch <= '9':
            number = int(ch) + number * int(math.pow(10, numberLen))
            numberLen += 1

    if numberLen != 0:
        result.append(number)
    return result

def insertNode(n, data):  # Рекурсивно добавляем новую ноду в бинарное дерево
    if n is None:
        n = Node(data)
        n.Data = data
        return n

    if n.Data > data:
        n.Left = insertNode(n.Left, data)
    else:
        n.Right = insertNode(n.Right, data)
    return n

def preorderTraversal(n):  # Распечатывание бинарного дерева "прямым" обходом
    if n is None:  # Проверяем, не является ли эта нода пустой
        return

    print(f"{n.Data}", end=" ")   # Сначала печатаем содержимое ноды
    preorderTraversal(n.Left)     # Потом для левой ветки
    preorderTraversal(n.Right)    # И для правой

def inorderTraversal(n):  # Распечатывание дерева "центральным" обходом
    if n is None:
        return

    inorderTraversal(n.Left)      # Печатаем левую ветку
    print(f"{n.Data}", end=" ")   # Печатаем содержимое этой ноды уже после распечатывания левой ветки
    inorderTraversal(n.Right)     # Печатаем правую

def postorderTraversal(n):  # Распечатываем дерево "концевым" обходом
    if n is None:
        return

    postorderTraversal(n.Left)
    postorderTraversal(n.Right)
    print(f"{n.Data}", end=" ")


input_str = "8 (3 (1, 6 (4,7)), 10 (, 14(13,)))" # Входные данные

arr = readInput(input_str) # Преобразовываем input в массив чисел
BinaryTree = None
for num in arr:
    BinaryTree = insertNode(BinaryTree, num) # Каждый эл-т массива добавляем в бинарное дерево

preorderTraversal(BinaryTree)
print()
inorderTraversal(BinaryTree)
print()
postorderTraversal(BinaryTree)
print()


