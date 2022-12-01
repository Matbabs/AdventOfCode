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
    const int DAYS = 80, CYCLE = 6, NEW_BORN = 9;
    ifstream file; string line;
    file.open("input_test.txt");
    getline(file, line);
    vector<int> lanternFishs = _split(line, ",");
    for(int i = 0; i < DAYS; i++){
        for(int lf = 0; lf < lanternFishs.size(); lf++){
            lanternFishs[lf]--;
            if(lanternFishs[lf] == -1){
                lanternFishs[lf] = CYCLE;
                lanternFishs.push_back(NEW_BORN);
            }
        }
    }
    cout << lanternFishs.size() << endl;
    return 0;
}