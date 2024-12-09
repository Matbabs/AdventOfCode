import re

d, m, _id = re.findall(r"\d", open("input.txt", "r").read()), [], 0
for i in range(len(d)):
    for j in range(int(d[i])):
        m.append(_id if i % 2 == 0 else -1)
    if i % 2 == 0:
        _id += 1
_id -= 1
while _id >= 0:
    s = m.count(_id)
    fp = m.index(_id)
    for i in range(len(m)):
        if all(b == -1 for b in m[i : i + s]) and len(m[i : i + s]) == s and i <= fp:
            for j in range(s):
                m[i + j] = _id
                m[fp + j] = -1
            break
    _id -= 1
print(sum(i * m[i] for i in range(len(m)) if m[i] != -1))
