from __future__ import print_function


def dump(nums):
    print(' '.join(str(i) for i in nums))

def func(nums, size):
    n = nums[-1]

    for i in xrange(size - 1, 0, -1):
        if nums[i - 1] > n:
            nums[i] = nums[i - 1]
            dump(nums)
        else:
            nums[i] = n
            dump(nums)

            return

    nums[0] = n
    dump(nums)


size = input()
inputs = raw_input()

nums = [int(i) for i in inputs.strip().split(' ')]
func(nums, size)
