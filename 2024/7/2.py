import re


def s(v):
    r = []
    for p in v:
        if len(r) == 0:
            r.append(p)
        else:
            nr = []
            for a in r:
                nr.append(a + p)
                nr.append(a * p)
                nr.append(int(f"{a}{p}"))
            r = nr
    return r


d = [
    list(map(int, re.findall(r"\d+", e)))
    for e in open("input.txt", "r").read().split("\n")
]
print(sum(e[0] if e[0] in s(e[1:]) else 0 for e in d))
