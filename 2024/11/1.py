def s(t):
    nt = []
    for r in t:
        sr = str(r)
        lsr = len(sr)
        if r == 0:
            nt.append(1)
        elif lsr % 2 == 0:
            nt.extend([int(sr[: lsr // 2]), int(sr[lsr // 2 :])])
        else:
            nt.append(r * 2024)
    return nt


d = list(map(int, open("input.txt", "r").read().split()))
for i in range(25):
    d = s(d)
print(len(d))
