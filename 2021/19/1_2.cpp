#include <iostream>
#include <string>
#include <set>
#include <tuple>
#include <vector>
#include <array>
#include <fstream>
using namespace std;

using ScannerView = set<tuple<int, int, int>>;
using DeltaView = vector<array<int, 3>>;

auto loadInput()
{
    ifstream file; string line;
    file.open("input.txt");
    vector<ScannerView> scanners;
    ScannerView view;
    while (getline(file, line))
    {
        if (line.size() < 2)
        {
            scanners.push_back(view);
            view.clear();
        }
        else if (line[1] != '-')
        {
            int x, y, z;
            sscanf(line.c_str(), "%d,%d,%d\n", &x, &y, &z);
            view.insert({x, y, z});
        }
    }
    scanners.push_back(view);
    return scanners;
}

array<int, 3> setDeltas(const ScannerView &view, const array<int, 3> &point, DeltaView &delta, int rx, int ry, int rz, int sx, int sy, int sz)
{
    delta.clear();
    for (const auto &[x, y, z] : view)
    {
        array<int, 3> coord = {x, y, z};
        delta.push_back({sx * (coord[rx] - point[rx]), sy * (coord[ry] - point[ry]), sz * (coord[rz] - point[rz])});
    }
    return {-sx * point[rx], -sy * point[ry], -sz * point[rz]};
}

bool tryAndMerge(ScannerView &map, const ScannerView &view, vector<array<int, 3>> &positions)
{
    if (map.empty())
    {
        for (const auto &p : view)
            map.insert(p);
        positions.push_back({0, 0, 0});
        return true;
    }

    //  [facing, up, right, sign_f, sign_u, sign_r]
    vector<tuple<int, int, int, int, int, int>> axis = {
        {0, 1, 2, 1, 1, 1},    //  x  y  z
        {0, 1, 2, 1, -1, -1},  //  x -y -z
        {0, 2, 1, 1, 1, -1},   //  x  z -y
        {0, 2, 1, 1, -1, 1},   //  x -z  y
        {0, 1, 2, -1, 1, -1},  // -x  y -z
        {0, 1, 2, -1, -1, 1},  // -x -y  z
        {0, 2, 1, -1, 1, 1},   // -x  z  y
        {0, 2, 1, -1, -1, -1}, // -x -z -y
    };

    vector<tuple<int, int, int, int, int, int>> rot;
    for (int a = 0; a < 3; a++)
        for (const auto &[dx, dy, dz, sx, sy, sz] : axis)
            rot.push_back({(a + dx) % 3, (a + dy) % 3, (a + dz) % 3, sx, sy, sz});

    DeltaView delta;
    for (const auto &[px, py, pz] : view)
    {
        for (const auto &[rx, ry, rz, sx, sy, sz] : rot)
        {
            auto sp = setDeltas(view, {px, py, pz}, delta, rx, ry, rz, sx, sy, sz);
            for (const auto &[x, y, z] : map)
            {
                int matches = 0;
                for (const auto &d : delta)
                    matches += (map.count({d[0] + x, d[1] + y, d[2] + z}));

                if (matches >= 12)
                { //Merge
                    for (const auto &d : delta)
                        map.insert({d[0] + x, d[1] + y, d[2] + z});
                    positions.push_back({sp[0] + x, sp[1] + y, sp[2] + z});
                    return true;
                }
            }
        }
    }
    return false;
}

int main()
{
    auto scanners = loadInput();

    cout << "ok " << endl;

    ScannerView map;

    set<int> remaining;
    for (int i = 0; i < scanners.size(); i++)
        remaining.insert(i);

    vector<array<int, 3>> sp;

    while (!remaining.empty())
        for (int i = 0; i < scanners.size(); i++)
            if (remaining.count(i) && tryAndMerge(map, scanners[i], sp))
                remaining.erase(i);

    int mm = 0;
    for (int i = 0; i < sp.size(); i++)
        for (int j = 0; j < sp.size(); j++)
            mm = max(abs(sp[i][0] - sp[j][0]) + abs(sp[i][1] - sp[j][1]) + abs(sp[i][2] - sp[j][2]), mm);

    cout << map.size() << endl; // Part 1
    cout << mm << endl;         // Part 2
    return 0;
}