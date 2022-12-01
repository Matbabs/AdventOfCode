#include <iostream>
#include <fstream>
#include <vector>
#include <set>
using namespace std;

vector<string> _split(string s, string delimiter){
    vector<string> arr;
    size_t pos = 0;
    string token;
    while((pos = s.find(delimiter)) != string::npos) {
        token = s.substr(0, pos);
        arr.push_back(token);
        s.erase(0, pos + delimiter.length());
    }
    arr.push_back(s);
    return arr;
}

int main(void) {
    const int DIM = 50;
    ifstream file; string line;
    file.open("input.txt");
    set<vector<int>> cubes_on;
    vector<pair<string, vector<pair<int, int>>>> steps;
    while(getline(file, line)){
        auto c = _split(line, " ")[0];
        auto x = _split(_split(line, "x=")[1], "..");
        auto y = _split(_split(line, "y=")[1], "..");
        auto z = _split(_split(line, "z=")[1], "..");
        steps.push_back({c, {{stoi(x[0]), stoi(x[1])}, {stoi(y[0]), stoi(y[1])}, {stoi(z[0]), stoi(z[1])}}});
    }
    for(auto s : steps)
        if(s.second[0].first >= -DIM && s.second[0].first < DIM && s.second[0].second >= -DIM && s.second[0].second)
            for(int x = s.second[0].first; x <= s.second[0].second; x++)
                for(int y = s.second[1].first; y <= s.second[1].second; y++)
                    for(int z = s.second[2].first; z <= s.second[2].second; z++)
                       if(s.first == "on") cubes_on.insert({x, y, z});
                       else cubes_on.erase({x, y, z});
    cout << cubes_on.size() << endl;
    return 0;
}