from __future__ import print_function
import itertools


def dump(nums):
    print(' '.join(str(i) for i in nums))


def qsort(nums, head, tail):
    if tail <= head:
        return

    pivot = head
    for i in xrange(head, tail):
        if nums[i] < nums[tail]:
            nums[pivot], nums[i] = nums[i], nums[pivot]
            pivot += 1

    nums[pivot], nums[tail] = nums[tail], nums[pivot]
    dump(nums)

    qsort(nums, head, pivot - 1)
    qsort(nums, pivot + 1, tail)


size = input()
numbers = raw_input()
numbers = [int(i) for i in numbers.strip().split(' ')]

qsort(numbers, 0, len(numbers) - 1)
