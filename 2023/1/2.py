m = {
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
}
lns = open("input.txt", "r").read().split("\n")
r = 0
for l in lns:
    n = ""
    for i in range(len(l)):
        if l[i].isdigit():
            n += l[i]
        for wn in m.items():
            if l[i : i + len(wn[0])] == wn[0]:
                n += str(wn[1])
    r += int(n[0] + n[-1])
print(r)
