#include <iostream>
#include <fstream>
#include <vector>
using namespace std;

int main(void) {
    const int MAP_SIZE = 1000;
    ifstream file; 
    string start, _, end;
    int x1, y1, x2, y2, sum = 0;
    file.open("input.txt");
    vector<vector<int>> tiles(MAP_SIZE, vector<int>(MAP_SIZE, 0));
    while(file >> start >> _ >> end){
        x1 = stoi(start.substr(0, start.find(',')));
        y1 = stoi(start.substr(start.find(',') + 1, start.size()));
        x2 = stoi(end.substr(0, end.find(',')));
        y2 = stoi(end.substr(end.find(',') + 1, end.size()));
        if(x1 == x2) for(int i = min(y1, y2); i <= max(y1, y2); i++) tiles[i][x1]++;
        if(y1 == y2) for(int i = min(x1, x2); i <= max(x1, x2); i++) tiles[y1][i]++;
        if(x1 != x2 && y1 != y2){
            tiles[y1][x1]++;
            while(x1 != x2 && y1 != y2){
                if(y1 != y2) (y2 - y1) > 0 ? y1++ : y1--;
                if(x1 != x2) (x2 - x1) > 0 ? x1++ : x1--;
                tiles[y1][x1]++;
            }
        }
    }
    for(auto line : tiles) for(auto c : line) if(c > 1) sum++;
    cout << sum << endl;
    return 0;
}