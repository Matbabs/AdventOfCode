import re

d, m, _id, s = re.findall(r"\d", open("input.txt", "r").read()), [], 0, {}
for i in range(len(d)):
    for j in range(int(d[i])):
        m.append(_id if i % 2 == 0 else -1)
    s[_id] = int(d[i])
    if i % 2 == 0:
        _id += 1
_id -= 1
while _id >= 0:
    fp = m.index(_id)
    for i in range(len(m)):
        if (
            m[i] == -1
            and all(b == -1 for b in m[i : i + s[_id]])
            and len(m[i : i + s[_id]]) == s[_id]
            and i <= fp
        ):
            for j in range(s[_id]):
                m[i + j] = _id
                m[fp + j] = -1
            break
    _id -= 1
print(sum(i * m[i] for i in range(len(m)) if m[i] != -1))
