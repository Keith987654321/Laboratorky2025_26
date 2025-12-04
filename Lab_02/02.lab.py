def evaluate_expression(expression):
    try:
        if not check_brackets(expression):  
            return "Ошибка: Некорректное расположение скобок"

        
        expression = expression.replace(" ", "").rstrip("=")
        result = calculate(expression)

        return result

    # Блок обработки ошибок:
    except SyntaxError:  
        return "Ошибка: Неверное математическое выражение"
    except ValueError:
        return "Ошибка: Некорректные данные"
    except ZeroDivisionError: 
        return "Ошибка: Деление на ноль"
    except Exception as e:  
        return f"Ошибка: {str(e)}"


def check_brackets(s):  # Функция проверки корректности скобок
    stack = []  # Стек для хранения скобок
    opening = "({["
    closing = ")}]"

    for char in s: 
        if char in opening:
            stack.append(char)  
        elif char in closing:
            if not stack:  
                return False
            top = stack.pop()  
            if (
                (char == ")" and top != "(")
                or (char == "}" and top != "{")
                or (char == "]" and top != "[")
            ):
                return False
    return len(stack) == 0  


def calculate(expression):  
    tokens = tokenize(expression) 
    return parse_expression(tokens)  


def tokenize(expr):  
    tokens = []  
    current_number = ""  # Временная строка для накопления цифр числа

    for char in expr:  
        if char.isdigit() or char == ".":
            current_number += char
        else:
            if current_number:
                tokens.append(float(current_number))
                current_number = ""  
            tokens.append(char)

    if current_number:
        tokens.append(float(current_number))

    return tokens  


def parse_expression(tokens):  # Обрабатывает сложение и вычитание
    # Начинаем с обработки умножения/деления (более высокий приоритет)
    result = parse_term(tokens)

    # Пока есть операции + или - в начале списка токенов
    while tokens and tokens[0] in ("+", "-"):
        operator = tokens.pop(0)  # Извлекаем оператор из начала списка
        right = parse_term(tokens)  # Обрабатываем правую часть выражения

        if operator == "+":
            result += right
        else:
            result -= right

    return result


def parse_term(tokens):  
    result = parse_factor(tokens)
    
    while tokens and tokens[0] in ("*", "/"):
        operator = tokens.pop(0) 
        right = parse_factor(tokens)  

        if operator == "*":
            result *= right
        else:
            result /= right 

    return result


def parse_factor(tokens):  # Обрабатывает базовые элементы: числа, скобки, унарный минус
    if not tokens: 
        raise SyntaxError()

    token = tokens.pop(0)

    if isinstance(token, float):
        return token
    elif token == "(":  
        result = parse_expression(tokens)
        if not tokens or tokens.pop(0) != ")":
            raise SyntaxError("Незакрытая скобка")
        return result
    elif token == "-":
        return -parse_factor(tokens)
    else:  
        raise SyntaxError(f"Неожиданный токен: {token}")


def main():  
    print("Введите арифметическое выражение (окончание ввода - '='):")
    expression = input()  

    if not expression.endswith("="):
        print("Ошибка: Выражение должно заканчиваться символом '='")
        return

    result = evaluate_expression(expression)
    print(f"Результат: {result}")


if __name__ == "__main__":
    main()  
