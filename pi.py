import random
import math

N_TRIALS = 100000
N_BOUND = 10000000

# def gcd(a, b):
#     if a > b:
#         return gcd(b, a)
#     if a == b:
#         return a
#     return gcd(b - a, a)


def gcd(x, y):
    a, b = x, y
    if b > a:
        b, a = x, y

    while True:
        if a == b:
            return a
        x = a - b
        y = b

        if x > y:
            a, b = x, y
        else:
            b, a = x, y


def is_co_prime(a, b):
    if gcd(a, b) == 1:
        return 1
    return 0


def calculate():
    coprimes = 0
    for i in range(N_TRIALS):
        x, y = random.randint(1, N_BOUND), random.randint(1, N_BOUND)
        coprimes += is_co_prime(x, y)
        if coprimes:
            fraction_co_prime = coprimes / (i + 1.0)
            pi_estimate = math.sqrt(6 / fraction_co_prime)
    return pi_estimate


for i in range(10):
    estimate = calculate()
    print i, estimate
