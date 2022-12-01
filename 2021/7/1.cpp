#include <iostream>
#include <fstream>
#include <vector>
#include <algorithm>
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
    ifstream file; string line;
    file.open("input.txt"); getline(file, line);
    vector<int> crabs = _split(line, ",");
    int res = 1e9, sum;
    for(int i = 0; i < crabs.size(); i++){
        sum = 0;
        for(auto cp : crabs)
            sum += abs(i - cp);
        res = min(res, sum);
    }   
    cout << res << endl;
    return 0;
}