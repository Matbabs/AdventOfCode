def s(px, py, p, d):
    v = set()
    while True:
        if f"{px},{py},{p}" in v:
            return v, True
        v.add(f"{px},{py},{p}")
        nx, ny, np = px, py, p
        match p:
            case "^":
                ny, np = py - 1, ">"
            case ">":
                nx, np = px + 1, "v"
            case "v":
                ny, np = py + 1, "<"
            case "<":
                nx, np = px - 1, "^"
        if nx >= 0 and nx < len(d[0]) and ny >= 0 and ny < len(d):
            if d[ny][nx] != "#":
                px, py = nx, ny
            else:
                p = np
        else:
            break
    return v, False


def b(nx, ny, np):
    bx, by, l = nx, ny, False
    match np:
        case "^":
            by = ny - 1
        case ">":
            bx = nx + 1
        case "v":
            by = ny + 1
        case "<":
            bx = nx - 1
    if bx >= 0 and bx < len(d[0]) and by >= 0 and by < len(d) and d[by][bx] == ".":
        d[by] = d[by][:bx] + "#" + d[by][bx + 1 :]
        _, l = s(px, py, p, d)
        d[by] = d[by][:bx] + "." + d[by][bx + 1 :]
    return bx, by, l


d, p, px, py, bv = open("input.txt", "r").read().split("\n"), "", 0, 0, set()
for y in range(len(d)):
    for x in range(len(d[0])):
        if d[y][x] in ["^", ">", "v", "<"]:
            px, py, p = x, y, d[y][x]
v, _ = s(px, py, p, d)
for n in v:
    sn = n.split(",")
    nx, ny, np = int(sn[0]), int(sn[1]), sn[2]
    bx, by, l = b(nx, ny, np)
    if l:
        bv.add(f"{bx},{by}")
print(len(bv))
