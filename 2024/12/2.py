from collections import defaultdict, deque

D = [(1, 0), (0, 1), (-1, 0), (0, -1)]


def isd(x, y):
    return x >= 0 and x < len(d[0]) and y >= 0 and y < len(d)


def bfs(sx, sy, m):
    tv = deque()
    tv.append((sx, sy))
    vst = set()
    vst.add((sx, sy))
    while len(tv) > 0:
        cx, cy = tv.popleft()
        for dx, dy in D:
            n = (cx + dx, cy + dy)
            nx, ny = n
            if isd(nx, ny) and m[cy][cx] == m[ny][nx]:
                if n not in vst:
                    tv.append(n)
                    vst.add(n)
    return vst


def wll(t):
    r = 0
    for p in s[t]:
        w = len(p) * 4
        for cx, cy in p:
            for dx, dy in D:
                nx, ny = (cx + dx, cy + dy)
                w -= 0 if isd(nx, ny) and d[ny][nx] == t else 1
        r += len(p) * w
    return r


d, s, vst = open("input.txt", "r").read().split("\n"), defaultdict(list), set()
for y in range(len(d)):
    for x in range(len(d[0])):
        if (x, y) not in vst:
            tv = bfs(x, y, d)
            s[d[y][x]].append(tv)
            for v in tv:
                vst.add(v)
print(sum(wll(t) for t in s))
