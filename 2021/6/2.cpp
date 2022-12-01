#include <iostream>
#include <fstream>
#include <vector>
#include <numeric>
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
    const int DAYS = 256, CYCLE = 7, NEW_BORN = 9;
    ifstream file; string line;
    file.open("input.txt");
    getline(file, line);
    vector<int> lanternFishs = _split(line, ",");
    vector<long> cycles(NEW_BORN + 1, 0);
    long res = 0;
    for(auto x : lanternFishs) cycles[x]++;
    for(int i = 0; i < DAYS; i++){
        cycles[CYCLE] += cycles[0];
        cycles[NEW_BORN] += cycles[0];
        for(int c = 0; c <= NEW_BORN ; c++)
            cycles[c] = cycles[c + 1];
    }
    for(auto lf : cycles) res += lf;
    cout << res << endl;
    return 0;
}