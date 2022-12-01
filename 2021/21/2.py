from functools import cache
from collections import Counter
from itertools import product

throws = Counter(map(sum, product(*[(1, 2, 3)] * 3)))

@cache
def wins(p1, p2, s1=0, s2=0):
  return s2 > 20 or sum(1j * count *
    wins(p2, state := (p1 + throw - 1) % 10 + 1, s2, s1 + state).conjugate()
    for throw, count in throws.items()
  )

c = wins(*eval("int(input().split()[-1])," * 2))
print(int(max(c.real, c.imag)))