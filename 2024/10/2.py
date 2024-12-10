from collections import deque

D = [(1, 0), (0, 1), (-1, 0), (0, -1)]


def bfs(sx, sy, m):
    tv = deque()
    tv.append((sx, sy, []))
    pth = []
    while len(tv) > 0:
        cx, cy, p = tv.popleft()
        e = False
        for dx, dy in D:
            n = (cx + dx, cy + dy)
            nx, ny = n
            if (
                nx >= 0
                and nx < len(m[0])
                and ny >= 0
                and ny < len(m)
                and m[cy][cx] == m[ny][nx] - 1
            ):
                if n not in p:
                    tv.append((nx, ny, p + [n]))
                    e = True
        if not e:
            lx, ly = p[-1]
            if m[ly][lx] == 9:
                pth.append(p)
    return pth


m, ss = [list(map(int, l)) for l in open("input.txt", "r").read().split("\n")], set()
for y in range(len(m)):
    for x in range(len(m[0])):
        if m[y][x] == 0:
            ss.add((x, y))
print(sum(len(bfs(sx, sy, m)) for sx, sy in ss))
