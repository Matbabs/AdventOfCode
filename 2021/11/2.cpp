#include <iostream>
#include <fstream>
#include <vector>
#include <queue>
using namespace std;

const int SIZE = 10, FLASH = 9;
vector<pair<int, int>> directions = {{1, 0}, {1, 1}, {0, 1}, {-1, 0}, {-1, -1}, {0, -1}, {-1, 1}, {1 ,-1}};

void flash(pair<int, int> node, vector<vector<int>>& tiles){
    if(tiles[node.first][node.second] > FLASH){
        queue<pair<int, int>> toFlash;
        tiles[node.first][node.second] = 0;
        for(auto d : directions){
            int y = node.first + d.first, x = node.second + d.second;
            if(y >= 0 && y < tiles.size() && x >= 0 && x < tiles[0].size()){
                if(tiles[y][x] != 0) tiles[y][x]++;
                if(tiles[y][x] > FLASH) toFlash.push({y ,x});
            }
        }
        while(toFlash.size()){
            auto node = toFlash.front(); toFlash.pop();
            flash({node.first, node.second}, tiles);
        }
    }
}

int main(void) {
    ifstream file; string line;
    file.open("input.txt");
    vector<vector<int>> tiles(SIZE, vector<int>(SIZE));
    int steps = 1;
    bool areAllFlashing = false;
    for(int i = 0; i < SIZE; i++){
        file >> line;
        for(int j = 0; j < SIZE; j++)
            tiles[i][j] = line[j] - '0';
    }
    while(!areAllFlashing){
        queue<pair<int, int>> toFlash;
        for(int i = 0; i < SIZE; i++)
            for(int j = 0; j < SIZE; j++){
                tiles[i][j]++;
                if(tiles[i][j] > FLASH) toFlash.push({i, j});
            }
        while(toFlash.size()){
            auto node = toFlash.front(); toFlash.pop();
            flash({node.first, node.second}, tiles);
        }
        areAllFlashing = true;
        for(int i = 0; i < SIZE; i++)
            for(int j = 0; j < SIZE; j++)
                if(tiles[i][j] != 0){ areAllFlashing = false; break;}
        if(!areAllFlashing) steps++;
    }
    cout << steps << endl;
    return 0;
}