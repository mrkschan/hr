from __future__ import print_function


def parition(nums, head, tail):
    pivot = head
    for i in xrange(head, tail):
        if nums[i] < nums[tail]:
            nums[pivot], nums[i] = nums[i], nums[pivot]
            pivot += 1

    nums[pivot], nums[tail] = nums[tail], nums[pivot]
    return pivot


size = input()
nums = raw_input()
nums = [int(i) for i in nums.strip().split(' ')]

mid = size // 2

pivot = size
head, tail = 0, size - 1
while pivot != mid:
    pivot = parition(nums, head, tail)
    if pivot > mid:
        tail = pivot - 1
    elif pivot < mid:
        head = pivot + 1

print(nums[mid])
