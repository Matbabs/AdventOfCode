#include <iostream>
#include <fstream>
#include <vector>
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

int main(void) {
    const int STEPS = 10;
    ifstream file; string line;
    file.open("input_test.txt");
    vector<pair<char, char>> char_pairs;
    map<pair<char, char>, char> instructions;
    map<char, int> occurences;
    bool isFormula = true;
    while(getline(file, line)){
        if(line.length())
            if(isFormula)
                for(int i = 0; i < line.length() - 1; i++){
                    char_pairs.push_back({line[i], line[i + 1]});
                    occurences[line[i]]++; occurences[line[i + 1]]++;
                }
            else {
                auto inst = _split(line, " -> ");
                instructions[{inst[0][0], inst[0][1]}] = inst[1][0];
            }
        else isFormula = false;
    }
    int steps = 1;
    while(steps <= STEPS){
        vector<pair<char, char>> new_char_pairs;
        for(auto p : char_pairs){
            if(instructions.find(p) != instructions.end()){
                new_char_pairs.push_back({p.first, instructions[p]});
                new_char_pairs.push_back({instructions[p], p.second});
                occurences[instructions[p]]++;
            }
        }
        char_pairs = new_char_pairs;
        steps++;
    }
    int min_c = 1e9, max_c = 0;
    for(auto o : occurences){
        min_c = min(min_c, o.second);
        max_c = max(max_c, o.second);
    }
    cout << max_c - min_c << endl;
    return 0;
}