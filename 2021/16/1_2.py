import unittest, functools

# Read puzzle input and return as usable data structure
def parse_input(file):
    hex_vals = []
    with open(file, 'r') as input:
        for line in input:
            hex_vals.extend([char for char in line.rstrip()])
    return hex_vals

# Find value of a literal value subpacket
def get_lit_val(packet):
    lit_val = ""
    while packet[0] == "1":
        lit_val += "".join(packet[1:5])
        packet = packet[5:]
    lit_val += "".join(packet[1:5])
    packet = packet[5:]
    return int(lit_val, 2), packet

# Recursive function to decode packets and subpackets
def pack_decode(packet):
    version, typeID, packet = int(packet[:3], 2), int(packet[3:6], 2), packet[6:]
    value = 0

    if typeID == 4:
        value, packet = get_lit_val(packet)
    else:
        lengthID, packet = int(packet[0], 2), packet[1:]
        literals = []
        if lengthID == 0:
            # If length type is 0, keep converting subpackets for length given
            length, packet = int(packet[:15], 2), packet[15:]
            sub_packet, packet = packet[:length], packet[length:]
            while sub_packet:
                sub_version, sub_value, sub_packet = pack_decode(sub_packet)
                version += sub_version
                literals.append(sub_value)
        else:
            # If length type is 1, convert x number of subpackets
            pack_num, packet = int(packet[:11], 2), packet[11:]
            for _ in range(pack_num):
                sub_version, sub_value, packet = pack_decode(packet)
                version += sub_version
                literals.append(sub_value)
        if typeID == 0:
            value = sum(literals)
        elif typeID == 1:
            value = functools.reduce(lambda x, y: x * y, literals)
        elif typeID == 2:
            value = min(literals)
        elif typeID == 3:
            value = max(literals)
        elif typeID == 5:
            value = 1 if literals[0] > literals[1] else 0
        elif typeID == 6:
            value = 1 if literals[0] < literals[1] else 0
        elif typeID == 7:
            value = 1 if literals[0] == literals[1] else 0

    return version, value, packet

def packet_decoder(file, result):
    hex_vals = parse_input(file)

    # Convert hex value to decimal, then to binary
    packet = f'{int("".join(hex_vals), 16):0>{4 * len(hex_vals)}b}'
    total_ver, total_val, packet = pack_decode(packet)

    # result = 0 for part 1
    # result = 1 for part 2
    if result == 0:
        return total_ver
    elif result == 1:
        return total_val

# Practicing unit testing on given test input and expected results
class Day16Tests(unittest.TestCase):
    def test_packet_decoder_part1(self):
        self.assertEqual(packet_decoder('Advent of Code/test.txt'), 20)
    
    def test_packet_decoder_part2(self):
        self.assertEqual(packet_decoder('Advent of Code/test.txt'), 1)

if __name__ == "__main__":
    #unittest.main(verbosity=2)
    
    # Part 1 solution
    #print(packet_decoder('Advent of Code/test.txt', 0))
    print(packet_decoder('input.txt', 0))
    
    # Part 2 solution
    #print(packet_decoder('Advent of Code/test.txt', 1))
    print(packet_decoder('input.txt', 1))