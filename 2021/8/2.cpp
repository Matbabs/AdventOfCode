#include <iostream>
#include <fstream>
#include <vector>
#include <algorithm>
#include <map>
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

bool checkLettersOfD(string v, int n, map<int, vector<char>> cspd){
    bool isValid = true;
    for(auto l : cspd[n]) if(find(v.begin(), v.end(), l) == v.end()){isValid = false; break;}
    return isValid;
}

bool checkPartLettersOfD(string v, int n, int c, map<int, vector<char>> cspd) {
    int isPart = 0;
    for(auto l : cspd[n]) if(find(v.begin(), v.end(), l) != v.end()) isPart++;
    return isPart == c;
}

int main(void) {
    ifstream file; string line;
    file.open("input.txt");
    vector<string> inputs, values, signals;
    map<int, int> size_to_val = {{2, 1}, {3, 7}, {4, 4}, {7, 8}};
    long res = 0;
    while(getline(file, line)){
        map<int, vector<char>> cspd;
        inputs = _split(line, "|");
        signals = _split(inputs[0], " ");
        values = _split(inputs[1], " ");
        for(auto s : signals)
            for(auto p : size_to_val)
                if(s.size() == p.first)
                    for(auto c : s) cspd[size_to_val[p.first]].push_back(c);
        string number = "";
        for(auto v : values){
            if(v.size() > 0){
                for(auto p : size_to_val)
                    if(v.size() == p.first)
                        number += to_string(size_to_val[p.first]);
                if(v.size() == 5){
                    if(checkPartLettersOfD(v, 4, 2, cspd)) number += '2';
                    else if(checkLettersOfD(v, 7, cspd)) number += '3';
                    else if(checkPartLettersOfD(v, 4, 3, cspd)) number += '5'; 
                }
                else if(v.size() == 6){
                    if(checkLettersOfD(v, 4, cspd)) number += '9';
                    else if(checkLettersOfD(v, 1, cspd)) number += '0';
                    else if(checkPartLettersOfD(v, 4, 3, cspd)) number += '6';
                }
            }
        }
        res += stoi(number);
    }
    cout << res << endl;
    return 0;
}