from __future__ import print_function


def func(s):
    last = s[0]
    discarded = 0

    for i in xrange(1, len(s)):
        this = s[i]
        if last == this:
            discarded += 1
        else:
            last = this

    return discarded


size = input()
for i in xrange(size):
    s = raw_input()
    c = func(s)

    print(c)
