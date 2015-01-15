from __future__ import print_function
import collections


def func(nums):
    counter = collections.Counter()

    for i in nums:
        counter[i] += 1

    return counter


size = input()
numbers = [raw_input() for i in xrange(size)]
numbers = [int(i.split(' ')[0]) for i in numbers]

counts = func(numbers)
levels = []
occurrence = 0
for i in xrange(100):
    occurrence += counts[i]
    levels.append(occurrence)

print(' '.join(str(i) for i in levels))
