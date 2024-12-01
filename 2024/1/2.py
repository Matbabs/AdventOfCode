d = open("input.txt", "r").read().split("\n")
l, r = [], []
for e in d:
    v = list(map(int, e.split()))
    l.append(v[0])
    r.append(v[-1])
print(sum(r.count(e) * e for e in l))
