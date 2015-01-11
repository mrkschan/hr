from __future__ import print_function


def func(s):
    s = list(s)

    head, tail = 0, len(s) - 1
    mutated = 0

    def mutate(c):
        return chr(ord(c) - 1)

    while head <= tail:
        if s[head] > s[tail]:
            s[head] = mutate(s[head])
            mutated += 1
        elif s[head] < s[tail]:
            s[tail] = mutate(s[tail])
            mutated += 1
        else:
            head += 1
            tail -= 1

    return mutated


size = input()
for i in xrange(size):
    s = raw_input()
    c = func(s)

    print(c)
