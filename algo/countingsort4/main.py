from __future__ import print_function
import collections
import itertools


size = input()
pairs = [raw_input() for i in xrange(size)]
pairs = (i.split(' ') for i in pairs)
pairs = [(int(i), j) for i, j in pairs]

half = len(pairs) / 2
masks = [(i, '-') if e < half else (i, j) for e, (i, j) in enumerate(pairs)]

counts = [0] * 100
for i, j in masks:
    counts[i] += 1

places = [0] * 100
lastplace = 0
for i, c in enumerate(counts):
    places[i] = lastplace
    lastplace += c

outputs = [None] * size
for i, j in masks:
    place = places[i]
    outputs[place] = j
    places[i] += 1

print(' '.join(i for i in outputs if i))
