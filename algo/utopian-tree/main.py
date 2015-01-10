from __future__ import print_function


def func(a):
    h = 1
    if a < 1:
        return h

    for g in xrange(1, a + 1):
        if (g % 2) == 0:
            h += 1
        else:
            h *= 2

    return h


size = input()
for i in xrange(size):
    g = input()

    print(func(g))
