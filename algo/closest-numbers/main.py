from __future__ import print_function


size = input()
nums = raw_input()
nums = [int(i) for i in nums.strip().split(' ')]

nums.sort()

min_diff = 10 ** 7
mins = []
for j in xrange(1, size):
    i = j - 1

    x, y = nums[i], nums[j]
    diff = y - x

    if diff < min_diff:
        min_diff = diff
        mins = [x, y]
    elif diff == min_diff:
        mins.append(x)
        mins.append(y)

print(' '.join(str(i) for i in mins))
