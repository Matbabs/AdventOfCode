#!/usr/bin/env python3
from collections import Counter, namedtuple, defaultdict
import copy
import sys
import time
import itertools

def p(cs):
    ranges = []
    for c in cs.split(","):
        d, r = c.split("=")
        a, b = r.split("..")
        a = int(a)
        b = int(b)
        ranges.append((a,b))
    return ranges


def read(fn):
    for line in open(fn):
        action, cs = line.split()
        action = action == "on"
        yield action, p(cs)

def clamp(r):
    a, b = r
    return a,b
    a = max(a, -50)
    b = min(b, 50)
    return a,b

def overlaps_one(a, b):
    a1, a2 = a
    b1, b2 = b

    if a2 < b1:
        return False
    if a1 > b2:
        return False
    return True

def split(a, b):
    """
    3 cases. fully inside
    to the left, to the right

    xxxxxxxx
       xxx

    xxxxxxxx
    xxx

    xxxxxxxx
         xxx
    treat all the same?
    """
    #a, b = sorted([a,b], key=lambda r: r.range)
    a1, a2 = a.range
    b1, b2 = b.range

    if a1 == b1 and a2 == b2:
        #print("Same!")
        action = a.action if a.prec > b.prec else b.action
        prec = max(a.prec, b.prec)
        return [TreeNode((a1, a2), action=action, prec=prec, children=a.children + b.children)]

    action = a.action if a.prec > b.prec else b.action
    prec = max(a.prec, b.prec)
    m1, m2 = overlapping_part(a.range, b.range)
    #print("Middle", m1, m2)
    m = TreeNode((m1, m2), action=action, prec=prec, children=a.children + b.children)
    new = [m]

    lr = left_part(a.range, b.range)
    if lr:
        l1, l2 = lr
        #print("Left", l1, l2)
        nex = a if overlaps_one((l1,l2), a.range) else b
        l = TreeNode((l1, l2), action=nex.action, prec=nex.prec, children=nex.children)
        new.append(l)

    rr = right_part(a.range, b.range)
    if rr:
        r1, r2 = rr
        #print("Right", r1, r2)
        nex = a if overlaps_one((r1,r2), a.range) else b
        r = TreeNode((r1, r2), action=nex.action, prec=nex.prec, children=nex.children)
        new.append(r)

    new.sort(key=lambda r: r.range)
    #print("inp", [a,b])
    #print("new", new)

    return new

def overlapping_part(a, b):
    a, b = sorted([a,b])
    return max(a[0], b[0]), min(a[1], b[1])

def left_part(a, b):
    a, b = sorted([a,b])
    m1, m2 = overlapping_part(a,b)
    res = a[0], m1-1
    if res[1] < res[0]:
        return None
    return res

def right_part(a, b):
    a, b = sorted([a,b])
    m1, m2 = overlapping_part(a,b)
    res = m2+1, max(a[1],b[1])
    if res[0] > res[1]:
        return None
    return res

def expand_once(ranges):
    new_ranges = []
    for r,c in itertools.combinations(ranges, 2):
        if overlaps_one(c.range, r.range):
            #print("overlaps", c, r)
            nr = split(c, r)
            #print("new", nr)
            #print()
            others = [x for x in ranges if x not in (c,r)]
            new_ranges.extend(nr)
            new_ranges.extend(others)
            return new_ranges, True
    return ranges, False

def expand(ranges):
    chg = True
    while True:
        new_ranges, chg = expand_once(ranges)
        if not chg:
            return new_ranges
        ranges = new_ranges

class TreeNode:
    def __init__(self, r, prec,action, children=None):
        self.range = r
        self.prec = prec
        self.action = action
        if children is None:
            children = []
        self.children = children

    def __repr__(self):
        return f"r={self.range} ch={len(self.children)} action={self.action} p={self.prec}"

    def size(self):
        return 1 + self.range[1] - self.range[0]
    
    def r(self):
        a, b = self.range
        return range(a, b+1)

def expand_part1(cs):
    xr, yr, zr = cs
    xa,xb = clamp(xr)
    ya,yb = clamp(yr)
    za,zb = clamp(zr)

    for x in range(xa, xb+1):
        for y in range(ya, yb+1):
            for z in range(za, zb+1):
                yield x,y,z

def part_1(fn):
    d = {}
    rules = list(read(fn))
    for action, cs in rules:
        for x,y,z in expand_part1(cs):
            if action:
                d[x,y,z] = True
            elif (x,y,z) in d:
                del d[x,y,z]
    return set(d)

def recurse(rules):
    total = 0
    its = 0
    for x in expand(rules):
        #print("  X", x)
        for y in expand(x.children):
            #print("    Y", y)
            #for z in y.children:
                #print("      Z(pre)", z)
            for z in expand(y.children):
                #print("      Z", z)
                if not z.action: continue
                total += x.size() * y.size() * z.size()
                its += 1
    print("loop iterations", its)
    return total

def make_tree(rules):
    t = []
    prec = 0
    for action, (xr,yr, zr) in rules:
        prec += 1
        #print(action, xr, yr, zr)
        nx = TreeNode(xr, prec=prec, action=action)
        ny = TreeNode(yr, prec=prec, action=action)
        nz = TreeNode(zr, prec=prec, action=action)
        nx.children=[ny]
        ny.children=[nz]
        #nx.children=[ny]
        t.append(nx)

    #t = expand(t)
    return recurse(t)

def run(fn):
    total = 0
    rules = list(read(fn))
    t = make_tree(rules)
    return t

def go(fn):
    s = time.time()
    print("ans", run(fn))
    e = time.time()
    print("took", e-s)

if __name__ == "__main__":
    go("input.txt")
