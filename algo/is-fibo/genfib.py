def fib(n):
    if n <= 1:
        return n

    f1 = fib.cache.get(n-1, fib(n - 1))
    f2 = fib.cache.get(n-2, fib(n - 2))

    ff = fib.cache[n] = f1 + f2
    return ff
fib.cache = {}

fib(40)

for i in sorted(fib.cache.keys()):
    print fib.cache[i]
