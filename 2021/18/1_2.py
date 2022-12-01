from math import ceil

def explode(row, i=0, nest=0, success=False):
    while i < len(row):
        nest = nest + 1 if row[i] == '[' else (nest - 1 if row[i] == ']' else nest)
        if nest == 5:
            left, right = [int(x.strip()) for x in row[i + 1 : i + row[i:].index(']')].split(',')]
            row = row[:i] + '0' + row[i + row[i:].index(']') + 1:]
            left_number = next((e + 1 for e,x in enumerate(row[:i][::-1]) if x.isdigit()), None)
            right_number = next((e + 1 for e,x in enumerate(row[i + 1:]) if x.isdigit()), None)
            if right_number:
                enum = right_number
                while row[i + enum].isdigit():
                    enum += 1
                row = row[:i + right_number] + str(int(row[i + right_number : i + enum]) + right) + row[enum + i:]
            if left_number:
                enum, temp = left_number, []
                while row[i - enum].isdigit():
                    temp.append(i - enum)
                    enum += 1
                row = row[:min(temp)] + str(int(row[min(temp) : max(temp) + 1]) + left) + row[max(temp) + 1:]
            i, nest = -1, 0
            success = True
        i += 1
    return success, row

def split_number(row, success = False):
    digit = next((int(y) for y in ''.join([x for x in row if x.isdigit() or x == ',']).split(',') if int(y) > 9), None)
    if digit:
        row = row[:row.index(str(digit))] + '[{},{}]'.format(digit // 2, ceil(digit / 2)) + row[row.index(str(digit)) + len(str(digit)):]
        success = True
    return success, row

def solve(row, success = True):
    while success:
        while success:
            success, row = explode(row)
        success, row = split_number(row)
    return row

def magnitude(row):
    while '[' in row:
        i = 0
        while i < len(row):
            if row[i].isdigit():
                if not '[' in row[i:] or row[i:].index('[') > row[i:].index(']'):
                    left, right = [int(x) for x in row[i : row.index(']')].split(',')]
                    row = row[:i-1] + str(left * 3 + right * 2) + row[row.index(']') + 1:]
                    break
            i += 1
    return int(row)

with open("input.txt", 'r') as file:
    data = file.read().splitlines()
    for e,x in enumerate(data):
        data[e] = solve(x)
    end = data[0]
    for x in data[1:]:
        end = solve('[{},{}]'.format(end, x))
    print(magnitude(end))
    print(max(magnitude(solve('[{},{}]'.format(x, y))) for x in data for y in data if x != y))