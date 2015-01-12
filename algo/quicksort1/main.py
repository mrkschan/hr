from __future__ import print_function
import itertools


def func(nums, size):
    pivot = nums[0]
    lpart = []
    rpart = []
    for i in xrange(1, size):
        n = nums[i]
        if n > pivot:
            rpart.append(n)
        else:
            lpart.append(n)

    return lpart, pivot, rpart


size = input()
numbers = raw_input()
numbers = [int(i) for i in numbers.strip().split(' ')]

lpart, pivot, rpart = func(numbers, size)
chained = ' '.join(str(i) for i in itertools.chain(lpart, [pivot], rpart))
print(chained)
