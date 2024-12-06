def s(px, py, p, d):
    v = set()
    while True:
        v.add(f"{px},{py}")
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
    return v


d, p, px, py, v = open("input.txt", "r").read().split("\n"), "", 0, 0, set()
for y in range(len(d)):
    for x in range(len(d[0])):
        if d[y][x] in ["^", ">", "v", "<"]:
            px, py, p = x, y, d[y][x]
print(len(s(px, py, p, d)))
