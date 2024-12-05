from functools import cmp_to_key


def safe(u):
    for i in range(len(u)):
        if any(u[i] in rls.get(p) for p in u[i + 1 :]):
            return 0
    return u[len(u) // 2]


def fix(u):
    return sorted(
        u,
        key=cmp_to_key(lambda p1, p2: 1 if (rls.get(p2) and p1 in rls.get(p2)) else -1),
    )[len(u) // 2]


d, rls, upd = open("input.txt", "r").read().split("\n"), {}, []
for l in d:
    if "|" in l:
        r = list(map(int, l.split("|")))
        rls.setdefault(r[0], []).append(r[1])
    if "," in l:
        upd.append(list(map(int, l.split(","))))
print(sum(fix(u) for u in upd if not safe(u)))
