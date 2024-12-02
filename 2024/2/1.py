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


d = open("input.txt", "r").read().split("\n")
print(sum(safe(list(map(int, l.split()))) for l in d))
