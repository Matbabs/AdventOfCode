def safe(n):
    d = -1 if n[0] > n[1] else 1
    for i in range(len(n)):
        if i + 1 < len(n) and (
            abs(n[i] - n[i + 1]) > 3
            or n[i] == n[i + 1]
            or (d == -1 and n[i] < n[i + 1])
            or (d == 1 and n[i] > n[i + 1])
        ):
            return 0
    return 1


d, r = open("input.txt", "r").read().split("\n"), 0
for l in d:
    n = list(map(int, l.split()))
    r += 1 if safe(n) or sum(safe(n[:i] + n[i + 1 :]) for i in range(len(n))) else 0
print(r)
