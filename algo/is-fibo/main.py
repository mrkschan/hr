from __future__ import print_function


# If either (5x^2 + 4) or (5x^2 - 4) is a perfect square, x is a fibo.
# Ref - http://is.gd/LAHyan
def func(q):
    def is_square(n):
        # Ref - http://stackoverflow.com/a/2489519/433662
        if n < 0:
            return False

        x = n // 2
        seen = {x}
        while x * x != n:
            x = (x + (n // x)) // 2
            if x in seen:
                return False
            seen.add(x)
        return True

    q1, q2 = 5 * (q ** 2) + 4, 5 * (q ** 2) - 4
    return is_square(q1) or is_square(q2)


size = input()
for i in xrange(size):
    n = input()
    if func(n):
        print('IsFibo')
    else:
        print('IsNotFibo')
