import re

d, m, _id = re.findall(r"\d", open("input.txt", "r").read()), [], 0
for i in range(len(d)):
    for j in range(int(d[i])):
        m.append(_id if i % 2 == 0 else -1)
    if i % 2 == 0:
        _id += 1
p = m.index(-1)
while p != -1:
    for i in range(len(m)):
        if m[len(m) - 1 - i] != -1:
            v = len(m) - 1 - i
            break
    m = m[:p] + [m[v]] + m[p + 1 : v]
    try:
        p = m.index(-1)
    except:
        break
print(sum(i * m[i] for i in range(len(m))))
