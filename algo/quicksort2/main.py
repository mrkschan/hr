from __future__ import print_function
import itertools


def dump(nums):
    print(' '.join(str(i) for i in nums))


def qsort(nums):
    if len(nums) < 2:
        return nums

    pivot = nums[0]
    lpart = []
    rpart = []
    for i in xrange(1, len(nums)):
        n = nums[i]
        if n > pivot:
            rpart.append(n)
        else:
            lpart.append(n)

    lpart = qsort(lpart)
    rpart = qsort(rpart)

    chained = list(itertools.chain(lpart, [pivot], rpart))
    dump(chained)

    return chained


size = input()
numbers = raw_input()
numbers = [int(i) for i in numbers.strip().split(' ')]

qsort(numbers)
