d = open("input.txt", "r").read().split("\n")
l, r = [], []
for e in d:
    v = list(map(int, e.split()))
    l.append(v[0])
    r.append(v[-1])
sl, sr = sorted(l), sorted(r)
print(sum(abs(sl[i] - sr[i]) for i in range(len(sl))))
