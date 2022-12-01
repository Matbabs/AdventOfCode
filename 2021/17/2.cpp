#include <iostream>
#include <fstream>
#include <vector>
#include <set>
#include <algorithm>
using namespace std;

vector<string> _split(string s, string delimiter){
    vector<string> arr;
    size_t pos = 0;
    string token;
    while ((pos = s.find(delimiter)) != string::npos) {
        token = s.substr(0, pos);
        arr.push_back(token);
        s.erase(0, pos + delimiter.length());
    }
    arr.push_back(s);
    return arr;
}

bool isReachable(int vx, int vy, int *xt, int *yt){
    int x = 0, y = 0;
    while(x < xt[1] && y > yt[0]){
        x += vx; y += vy;
        if(vx != 0) vx += vx > 0 ? -1 : 1;
        vy--;
        if(xt[0] <= x && x <= xt[1] && yt[0] <= y && y <= yt[1])
            return true;
    }
    return false;
}

int main(void) {
    const int BRUTE_FORCE = 1000;
    set<pair<int, int>> possibles;
    ifstream file; string line;
    vector<string> line_x, line_y;
    file.open("input.txt");
    int xt[2], yt[2];
    while(getline(file, line)){
        line_x = _split(_split(_split(line, "x=")[1], ",")[0], "..");
        line_y = _split(_split(line, "y=")[1], "..");
        for(int i = 0; i < 2; i++){
            xt[i] = stoi(line_x[i]);
            yt[i] = stoi(line_y[i]);
        }
    }
    for(int i = -BRUTE_FORCE; i < BRUTE_FORCE; i++)
        for(int j = 0; j < BRUTE_FORCE; j++)
            if(isReachable(j, i, xt, yt))
                possibles.insert({j, i});
    cout << possibles.size() << endl;
    return 0;
}