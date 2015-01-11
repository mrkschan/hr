from __future__ import print_function


def dump(nums):
    print(' '.join(str(i) for i in nums))


def reorder(nums, tail):
    n = nums[tail]

    for i in xrange(tail, 0, -1):
        if nums[i - 1] > n:
            nums[i] = nums[i - 1]
        else:
            nums[i] = n
            return

    nums[0] = n


def func(nums, size):
    for i in xrange(1, size):
        if nums[i] < nums[i-1]:
            reorder(nums, i)
        dump(nums)


size = input()
inputs = raw_input()

nums = [int(i) for i in inputs.strip().split(' ')]
func(nums, size)
