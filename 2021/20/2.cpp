#include <iostream>
#include <fstream>
#include <vector>
#include <algorithm>
using namespace std;

int main(void) {
    const int MARGIN = 2, TIMES = 50;
    const string DARK = ".", LIGHT = "#";
    const char DARK_BIT = '0', LIGHT_BIT = '1';
    ifstream file; string line;
    file.open("input_test.txt");
    string enhanced_algo = "", spacer = "";
    vector<string> tiles, new_tiles_extended;
    int lits = 0;
    while(getline(file, line)){
        if(!enhanced_algo.size()) enhanced_algo = line;
        else if (line.size()) tiles.push_back(line);
    };
    for(int k = 0; k < TIMES; k++){
        vector<string> tiles_extended; spacer = "";
        for(int i = 0; i < tiles[0].size(); i++) spacer += DARK;
        for(int i = 0; i < MARGIN; i++) spacer = DARK + spacer;
        for(int i = 0; i < MARGIN; i++) spacer = spacer + DARK;
        for(int i = 0; i < MARGIN; i++) tiles_extended.push_back(spacer);
        for(auto row : tiles){
            for(int i = 0; i < MARGIN; i++) row = DARK + row;
            for(int i = 0; i < MARGIN; i++) row = row + DARK;
            tiles_extended.push_back(row);
        }
        for(int i = 0; i < MARGIN; i++) tiles_extended.push_back(spacer);
        new_tiles_extended = tiles_extended;
        lits = 0;
        for(int i = 1; i < tiles_extended.size() - 1; i++){
            for(int j = 1; j < tiles_extended[0].size() - 1; j++){
                string number = "";
                for(int r = -1; r <= 1; r++)
                    for(int s = -1; s <= 1; s++)
                        number += tiles_extended[i + r][j + s] == LIGHT[0] ? LIGHT_BIT : DARK_BIT;
                int decimal_value = stoull(number, 0, 2);
                new_tiles_extended[i][j] = enhanced_algo[decimal_value];
                lits += enhanced_algo[decimal_value] == LIGHT[0] ? 1 : 0;
            }
        }
        tiles = new_tiles_extended;
    }
    cout << lits << endl;
    return 0;
}