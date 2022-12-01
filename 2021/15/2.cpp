#include <iostream>
#include <fstream>
#include <vector>
#include <map>
#include <set>
using namespace std;

int main(void) {
    ifstream file; string line;
    file.open("input.txt");
    int marks = 0;
    vector<vector<int>> tls;
    map<pair<int, int>, int> distances;
    map<pair<int, int>, pair<int, int>> predecesors;
    map<pair<pair<int, int>, pair<int, int>>, int> weights;
    vector<pair<int, int>> directions = {{0, 1}, {1, 0}, {-1, 0}, {0, -1}};
    while(getline(file, line)){
        vector<int> row;
        for(auto c : line) row.push_back(c - '0');
        tls.push_back(row);
    }
    const int RATIO = 5;
    const int SIZE = tls.size(), BIG_SIZE = SIZE * RATIO;
    int tiles[SIZE * RATIO][SIZE * RATIO];
    for(int i = 0; i < BIG_SIZE; i++)
        for(int j = 0; j < BIG_SIZE; j++)
            tiles[i][j] = (((tls[i % SIZE][j % SIZE] + i / SIZE + j / SIZE) - 1) % 9) + 1;
    int mark[BIG_SIZE][BIG_SIZE] = {{0}};
    for(int i = 0; i < BIG_SIZE; i++)
        for(int j = 0; j < BIG_SIZE; j++){
            distances[{i, j}] = 1e9;
            for(auto d : directions){
                int y = i + d.first, x = j + d.second;
                if(x >= 0 && y >=0 && x < BIG_SIZE && y < BIG_SIZE)
                    weights[{{y, x}, {i, j}}] = tiles[y][x];
            }
        }
    distances[{0, 0}] = 0;
    while(distances.size() != marks){
        pair<int, int> node, voisin;
        int min_d = 1e9;
        for(auto dist : distances)
            if(!mark[dist.first.first][dist.first.second] && min_d > dist.second){
                min_d = dist.second;
                node = dist.first;
            }
        mark[node.first][node.second] = 1; marks++;
        for(auto d : directions){
            int y = node.first + d.first, x = node.second + d.second;
            if(x >= 0 && y >=0 && x < BIG_SIZE && y < BIG_SIZE){
                voisin = {y, x};
                int new_dist = distances[node] + weights[{voisin, node}];
                if(distances[voisin] > new_dist){
                    distances[voisin] = new_dist;
                    predecesors[voisin] = node;
                }
            }
        }
    }
    cout << distances[{BIG_SIZE - 1, BIG_SIZE - 1}] << endl;
    return 0;
}