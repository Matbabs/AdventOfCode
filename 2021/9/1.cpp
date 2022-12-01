#include <iostream>
#include <fstream>
#include <vector>
using namespace std;

int main(void) {
    ifstream file; string line;
    file.open("input.txt");
    vector<vector<int>> tiles;
    vector<pair<int, int>> directions = {{0, 1}, {1, 0}, {0, -1}, {-1, 0}};
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
            res += isLowPoint ? 1 + tiles[i][j] : 0;
        }
    cout << res << endl;
    return 0;
}