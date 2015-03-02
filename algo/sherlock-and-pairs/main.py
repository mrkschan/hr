from collections import Counter

cases = input()

for i in xrange(int(cases)):
    counter = Counter()
    size = input()
    nums = raw_input()
    for i in nums.strip().split(' '):
        counter[i] += 1

    print sum(v * (v-1) for _, v in counter.items())
