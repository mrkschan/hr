from __future__ import print_function


def func(a):
  return a ^ 0xffffffff


size = input()
for i in xrange(size):
    n = input()

    print(func(n))
