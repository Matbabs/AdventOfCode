#include <iostream>
#include <fstream>
#include <vector>  
using namespace std;

vector<int> _split(string s, string delimiter){
    vector<int> arr;
    size_t pos = 0;
    string token;
    while ((pos = s.find(delimiter)) != string::npos) {
        token = s.substr(0, pos);
        arr.push_back(stoi(token));
        s.erase(0, pos + delimiter.length());
    }
    arr.push_back(stoi(s));
    return arr;
}   

int main(void) {
    const int SIZE = 5;
    ifstream file; string line, in;
    vector<int> numbers;
    vector<vector<vector<pair<int, bool>>>> grids;
    vector<vector<pair<int, bool>>> grid(SIZE, vector<pair<int, bool>>(SIZE));
    int l = 0, i = 0, res_number, res_grid, res_sum = 0;
    file.open("input.txt");
    while(getline(file, line)){
        if(l == 0) numbers = _split(line, ",");
        if(l !=0 && line.length() != 0){
            vector<int> numb = _split(line, " ");
            for(int j = 0; j < numb.size(); j++)
                grid[i][j] = {numb[j], false};
            i++;
            if(i == SIZE){
                grids.push_back(grid);
                i = 0;
            }
        }
        l++;
    }
    while(numbers.size()){
        bool isWinning = false;
        int n = numbers[0], nb_grid = 0;
        for(auto& grid : grids){
            // apply number in grids
            for(int i = 0; i < SIZE; i++)
                for(int j = 0; j < SIZE; j++)
                    if(n == grid[i][j].first)
                        grid[i][j].second = true;
            // check win for row or columns
            auto checkWin = [&](bool mode){
                for(int i = 0; i < SIZE; i++){
                    isWinning = true;
                    for(int j = 0; j < SIZE; j++)
                        if(mode && !grid[i][j].second || !mode && !grid[j][i].second){
                            isWinning = false; break;
                        }
                    if(isWinning) break;
                }
                return isWinning;
            };
            if(checkWin(true)) break;
            if(checkWin(false)) break;
            if(nb_grid < grids.size() - 1) nb_grid++;
        }
        if(isWinning){
            numbers.clear();
            res_number = n;
            res_grid = nb_grid;
        } else
            numbers.erase(numbers.begin());
    }
    for(int i = 0; i < SIZE; i++)
        for(int j = 0; j < SIZE; j++)
            if(!grids[res_grid][i][j].second)
                res_sum += grids[res_grid][i][j].first;
    cout << res_sum * res_number << endl;
    return 0;
}