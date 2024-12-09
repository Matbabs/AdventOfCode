d, s = open("input.txt", "r").read().split("\n"), set()
v = lambda x, y: x >= 0 and x < len(d[0]) and y >= 0 and y < len(d)
p = lambda x, y: s.add(f"{x},{y}")
for y in range(len(d)):
    for x in range(len(d[0])):
        if d[y][x] != ".":
            a = d[y][x]
            for y2 in range(len(d)):
                for x2 in range(len(d[0])):
                    if x != x2 and y != y2 and a == d[y2][x2]:
                        p(x, y), p(x2, y2)
                        dx, dy = x - x2, y - y2
                        nx1, ny1 = x + dx, y + dy
                        nx2, ny2 = x2 - dx, y2 - dy
                        while v(nx1, ny1):
                            p(nx1, ny1)
                            nx1, ny1 = nx1 + dx, ny1 + dy
                        while v(nx2, ny2):
                            p(nx2, ny2)
                            nx2, ny2 = nx2 - dx, ny2 - dy
print(len(s))
