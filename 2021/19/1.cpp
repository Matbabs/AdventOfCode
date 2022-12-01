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

void cout_scanner(vector<vector<int>> scanner){
    for(auto beacons : scanner){
        for(auto coords : beacons)
            cout << coords << ", ";
        cout << endl;
    }
}

void rotation_scanner(vector<vector<int>> scanner){
    cout_scanner(scanner);
}

int main(void) {
    ifstream file; string line;
    file.open("input_test.txt");
    vector<vector<int>> scanner;
    vector<vector<vector<int>>> scanners;
    while(getline(file, line)){
        if(line.size() && line[0] != '-') {
            scanner.push_back(_split(line, ","));
        } else if(!line.size())
            scanners.push_back(scanner);
    }

    rotation_scanner(scanners[0]);

    return 0;
}