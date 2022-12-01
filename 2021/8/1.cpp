#include <iostream>
#include <fstream>
#include <vector>
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
    vector<string> inputs, values;
    int sum = 0;
    int easy_digits[4] = {2, 3, 4, 7};
    while(getline(file, line)){
        inputs = _split(line, "|");
        values = _split(inputs[1], " ");
        for(auto v : values)
            for(auto d : easy_digits)
                if(v.size() == d)
                    sum++;
    }
    cout << sum << endl;
    return 0;
}