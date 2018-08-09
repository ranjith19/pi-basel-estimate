import random
import math
cdef int N_TRIALS = 100000
cdef int N_BOUND = 10000000

cdef int gcd(int x, int y):
    cdef int a = x
    cdef int b = y
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


cdef int is_co_prime(int a, int b):
    if gcd(a, b) == 1:
        return 1
    return 0





cdef float calculate():
    cdef int coprimes = 0
    cdef float pi_estimate 
    cdef int i
    for i in range(N_TRIALS):
        x, y = random.randint(1, N_BOUND), random.randint(1, N_BOUND)
        coprimes += is_co_prime(x, y)
        if coprimes:
            fraction_co_prime = coprimes / (i + 1.0)
            pi_estimate = math.sqrt(6 / fraction_co_prime)
    return pi_estimate

cpdef run():
    for i in range(10):
        estimate = calculate()
        print i, estimate
