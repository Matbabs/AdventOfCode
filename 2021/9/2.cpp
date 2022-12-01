#include <iostream>
#include <fstream>
#include <vector>
#include <queue>
#include <set>
#include <numeric>
#include <algorithm>
using namespace std;

vector<pair<int, int>> directions = {{0, 1}, {1, 0}, {0, -1}, {-1, 0}};

int bfs(pair<int, int> start, vector<vector<int>> tiles){
    queue<pair<int, int>> toVisit;
    set<pair<int, int>> mark;
    int size = 0;
    toVisit.push(start);
    mark.insert(start);
    while(toVisit.size()){
        size++;
        auto node = toVisit.front(); toVisit.pop();
        for(auto d : directions){
            int y = node.first + d.first, x = node.second + d.second;
            if(y >= 0 && y < tiles.size() && x >= 0 && x < tiles[0].size() && tiles[y][x] < 9 && mark.find({y, x}) == mark.end()){
                toVisit.push({y, x});
                mark.insert({y, x});
            }
        }
    }
    return size;
}

int main(void) {
    ifstream file; string line;
    file.open("input.txt");
    vector<vector<int>> tiles;
    vector<int> bassins;
    int res = 0;
    while(getline(file, line)){
        vector<int> row;
        for(auto c : line)
            row.push_back(c - '0');
        tiles.push_back(row);
    }
    for(int i = 0; i < tiles.size(); i++)
        for(int j = 0; j < tiles[0].size(); j++){
            bool isLowPoint = true;
            for(auto d : directions){
                int y = i + d.first, x = j + d.second;
                if(y >= 0 && y < tiles.size() && x >= 0 && x < tiles[0].size() && tiles[y][x] <= tiles[i][j]){
                    isLowPoint = false;
                    break;
                }
            }
            if(isLowPoint) bassins.push_back(bfs({i, j}, tiles));
        }
    sort(bassins.begin(), bassins.end());
    cout << accumulate(bassins.end() - 3, bassins.end(), 1, multiplies<int>()) << endl;
    return 0;
}