def p(r):
    sr = str(r)
    lsr = len(sr)
    if r == 0:
        return [1]
    elif lsr % 2 == 0:
        return [int(sr[: lsr // 2]), int(sr[lsr // 2 :])]
    else:
        return [r * 2024]


def s(t):
    for r in t:
        yield from p(r)


d = list(map(int, open("input.txt", "r").read().split()))
for i in range(25):
    d = s(d)
print(sum(1 for _ in d))
