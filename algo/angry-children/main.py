from __future__ import print_function


def func(numbers, range_size):
    m = 10**9
    offset = range_size - 1

    for i in xrange(offset, len(numbers)):
        diff = numbers[i] - numbers[i - offset]
        m = min(m, diff)

    return m


size = input()
range_size = input()

numbers = [input() for i in xrange(size)]
numbers.sort()

m = func(numbers, range_size)
print(m)
