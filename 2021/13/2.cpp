#include <iostream>
#include <fstream>
#include <vector>
#include <set>
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

int main(void) {
    ifstream file; string line;
    file.open("input.txt");
    bool isDots = true;
    vector<string> in;
    set<pair<int, int>> dots;
    vector<pair<string, int>> instructions;
    while(getline(file, line)){
        if(line.size()){
            if(isDots){
                in = _split(line, ",");
                dots.insert({stoi(in[0]), stoi(in[1])});
            } else {
                in = _split(line, " ");
                in = _split(in[in.size() - 1], "=");
                instructions.push_back({in[0], stoi(in[1])});
            }
        } else isDots = false;
    }
    for(auto i : instructions){
        set<pair<int, int>> new_dots;
        for(auto d : dots)
            new_dots.insert({
                i.first == "x" && d.first > i.second ? i.second * 2 - d.first : d.first,
                i.first == "y" && d.second > i.second ? i.second * 2 - d.second : d.second
            });
        dots = new_dots;
    }
    int max_x = 0, max_y = 0;
    for(auto d : dots){
        max_x = max(max_x, d.first);
        max_y = max(max_y, d.second);
    }
    for(int i = 0; i <= max_y; i++){
        for(int j = 0; j <= max_x; j++)
            cout << (dots.find({j, i}) != dots.end() ? "#" : " ");
        cout << endl;
    }
    return 0;
}