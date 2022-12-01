#include <iostream>
#include <fstream>
#include <vector>
using namespace std;

int main(void) {
    ifstream file; string line;
    file.open("input.txt");
    vector<vector<char>> tiles;
    int steps = 0;
    while(getline(file, line)){
        vector<char> row(line.size());
        for(int j = 0; j < line.size(); j++)
            row[j] = line[j];
        tiles.push_back(row);
    }
    bool isMoving = true;
    while(isMoving){
        isMoving = false;
        vector<vector<char>> east, south;
        east = tiles;
        for(int i = 0; i < tiles.size(); i++)
            for(int j = 0; j < tiles[i].size(); j++){
                int j_new_pos = (j + 1) % tiles[0].size();
                if(tiles[i][j] == '>' && tiles[i][j_new_pos] == '.'){
                    east[i][j] = '.';
                    east[i][j_new_pos] = '>';
                    isMoving = true;
                }
            }
        south = east;
        for(int i = 0; i < east.size(); i++)
            for(int j = 0; j < east[i].size(); j++){
                int i_new_pos = (i + 1) % east.size();
                if(east[i][j] == 'v' && east[i_new_pos][j] == '.'){
                    south[i][j] = '.';
                    south[i_new_pos][j] = 'v';
                    isMoving = true;
                }
            }
        tiles = south;
        steps++;
    }
    cout << steps << endl;
    return 0;
}