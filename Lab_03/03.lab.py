x = int(input())

result = []

for i in range(1, x + 1):  
    n = i  
    
    while n % 3 == 0:
        n //= 3

    while n % 5 == 0:
        n //= 5

    while n % 7 == 0:
        n //= 7

    if n == 1:  
        result += [
            i
        ]  
print(result)
