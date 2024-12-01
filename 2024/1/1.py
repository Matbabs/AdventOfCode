d = open("input.txt", "r").read().split("\n")
l, r, s = [], [], 0
for e in d:
    v = list(map(int, e.split()))
    l.append(v[0])
    r.append(v[-1])
sl, sr = sorted(l), sorted(r)
for i in range(len(sl)):
    s += abs(sl[i] - sr[i])
print(s)
