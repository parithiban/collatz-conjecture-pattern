from datetime import datetime

print(datetime.now())

chain = 0
def collatz(n):
    list1 = [n]
    # chain = 0
    if n == 1 :
        return 1                 
    elif n % 2 == 0 :
        collatz(n//2)
        # chain += 1
    else:
        collatz(n*3+1)
        # chain += 1
    # return chain
 

print(collatz(25))
# for n in range(3,100000000):
#     collatz(n)

print(datetime.now())