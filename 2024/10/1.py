from collections import deque

D = [(1, 0), (0, 1), (-1, 0), (0, -1)]


def bfs(sx, sy, m):
    tv = deque()
    tv.append((sx, sy))
    vst, r = set(), 0
    vst.add((sx, sy))
    while len(tv) > 0:
        cx, cy = tv.popleft()
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
                if n not in vst:
                    tv.append(n)
                    vst.add(n)
        if m[cy][cx] == 9:
            r += 1
    return r


m, ss = [list(map(int, l)) for l in open("input.txt", "r").read().split("\n")], set()
for y in range(len(m)):
    for x in range(len(m[0])):
        if m[y][x] == 0:
            ss.add((x, y))
print(sum(bfs(sx, sy, m) for sx, sy in ss))
