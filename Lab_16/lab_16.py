import math

class Node:
    def __init__(self, data=None):
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

def insertNode(n, data):  # Так же, как и в 15 лабе
    if n is None:
        n = Node()
        n.Data = data
        return n

    if n.Data > data:
        n.Left = insertNode(n.Left, data)
    else:
        n.Right = insertNode(n.Right, data)
    return n

def treeTraversal(n):
    stack = []  # Создаём стек
    stack.append(n)
    while len(stack) > 0:
        v = stack[-1]           # Вытаскиваем эл-т из стека
        stack = stack[:-1]      # Удаляем этот эл-т из стека

        if v is not None:
            print(f"{v.Data}", end=" ")   # Печатаем данные из ноды
            stack.append(v.Right)         # Добавляем правую ноду
            stack.append(v.Left)          # И левую


input_str = "8 (3 (1, 6 (4,7)), 10 (, 14(13,)))"
arr = readInput(input_str)

BinaryTree = None
for num in arr:
    BinaryTree = insertNode(BinaryTree, num)

treeTraversal(BinaryTree)
