from functools import lru_cache


@lru_cache(maxsize=None)
def dp(r):
    sr = str(r)
    lsr = len(sr)
    if r == 0:
        return [1]
    elif lsr % 2 == 0:
        return [int(sr[: lsr // 2]), int(sr[lsr // 2 :])]
    else:
        return [r * 2024]


def s(dd):
    nt = {}
    for r, c in dd.items():
        for nr in dp(r):
            nt[nr] = nt.get(nr, 0) + c
    return nt


d, dd = list(map(int, open("input.txt", "r").read().split())), {}
for r in d:
    dd[r] = dd.get(r, 0) + 1
for i in range(75):
    dd = s(dd)
print(sum(dd.values()))
