from __future__ import print_function


def func(l, r):
    if l == r:
        return 0

    effective = l ^ r
    effective_bits = len(bin(effective)) - 2

    return int('1' * effective_bits, 2)


# def func(l, r):
#     m = 0
#     for i in xrange(l, r+1):
#         for j in xrange(l+1, r+1):
#             m = max(m, i ^ j)
# 
#     return m


l, r = input(), input()
m = func(l, r)

print(m)
