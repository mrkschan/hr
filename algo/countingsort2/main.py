from __future__ import print_function
import collections
import itertools


def func(nums):
    counter = collections.Counter()

    for i in nums:
        counter[i] += 1

    return counter


size = input()
numbers = raw_input()
numbers = [int(i) for i in numbers.strip().split(' ')]

counts = func(numbers)
output = [[i] * counts[i] for i in sorted(counts.keys())]
output = itertools.chain(*output)
output = ' '.join(str(i) for i in output)
print(output)
