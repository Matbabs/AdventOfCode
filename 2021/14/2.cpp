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
    const int STEPS = 40;
    ifstream file; string line;
    file.open("input.txt");
    map<pair<char, char>, char> instructions;
    map<pair<char, char>, long> pairs_occurences;
    map<char, long> char_occurences; 
    bool isFormula = true;
    while(getline(file, line)){
        if(line.length())
            if(isFormula)
                for(int i = 0; i < line.length() - 1; i++){
                    pairs_occurences[{line[i], line[i + 1]}]++;
                    char_occurences[line[i]]++; char_occurences[line[i + 1]]++;
                }
            else {
                auto inst = _split(line, " -> ");
                instructions[{inst[0][0], inst[0][1]}] = inst[1][0];
            }
        else isFormula = false;
    }
    int steps = 1;
    while(steps <= STEPS){
        map<pair<char, char>, long> new_pairs_occurences;
        for(auto o : pairs_occurences){
            if(instructions.find(o.first) != instructions.end()){
                new_pairs_occurences[{o.first.first, instructions[o.first]}] += o.second;
                new_pairs_occurences[{instructions[o.first], o.first.second}] += o.second;
                char_occurences[instructions[o.first]] += o.second;
            }
        }
        pairs_occurences = new_pairs_occurences;
        steps++;
    }
    long min_c = 1e12, max_c = 0;
    for(auto o : char_occurences){
        min_c = min(min_c, o.second);
        max_c = max(max_c, o.second);
    }
    cout << max_c - min_c << endl;
    return 0;
}