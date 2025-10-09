"""
Идея решения:
Складываем все открывающие скобки в голову односвязного списка,
а когда встречаем закрывающую скобку, то проверяем, подходят ли они друг дружке.
"""

class Node:
    # Односвязный список
    def __init__(self, value=None, next_node=None):
        self.value = value
        self.next = next_node


class Stack:
    def __init__(self):
        self.head = Node()  

    def add_to_head(self, val: str):
        # Добавляем элемент в голову списка
        node = Node(val)
        node.next = self.head.next
        self.head.next = node

    def get(self) -> str:
        # Извлекаем элемент из головы списка
        if self.head.next is None:
            return "stoy, mne ne priyatno!"  

        result = self.head.next.value
        self.head.next = self.head.next.next
        return result

    def is_empty(self) -> bool:
        # Проверяем, пуст ли стек
        return self.head.next is None


def is_valid(s: str) -> bool:
    stack = Stack()

    for char in s:
        if char in "({[":
            stack.add_to_head(char)
            continue

        top = stack.get()

        if top == "stoy, mne ne priyatno!":
            return False

        if char == ")" and (top + char) != "()":
            return False
        elif char == "}" and (top + char) != "{}":
            return False
        elif char == "]" and (top + char) != "[]":
            return False

    return stack.is_empty()



print(is_valid("()"))        
print(is_valid("[{()}]"))    
print(is_valid("()[]{}"))    
print(is_valid("({)"))       
print(is_valid("({)}"))      
print(is_valid(""))         
