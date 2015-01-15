from __future__ import print_function
import collections


def func(nums):
    counter = collections.Counter()

    for i in nums:
        counter[i] += 1

    return counter


size = input()
numbers = raw_input()
numbers = [int(i) for i in numbers.strip().split(' ')]

counts = func(numbers)
output = ' '.join(str(counts[i]) for i in xrange(100))
print(output)
