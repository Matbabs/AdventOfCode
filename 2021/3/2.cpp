#include <iostream>
#include <fstream>
#include <vector>
using namespace std;

int main(void) {
    ifstream file; string line, oxygen = "", co2 = "";
    vector<string> diags;
    file.open("input.txt");
    while(getline(file, line))
        diags.push_back(line);
    auto rating = [](vector<string> diags, bool mode){
        vector<string> _diags = diags;
        int bits[12][2] = {0};
        int i = 0;
        while(i < 12 && _diags.size() > 1){
            vector<string> new_diags;
            for(auto d : _diags)
                if(d[i] == '0') bits[i][0]++; else bits[i][1]++;
            char occ;
            if(mode) occ = bits[i][0] > bits[i][1] ? '0' : '1';
            if(!mode) occ = bits[i][1] < bits[i][0] ? '1' : '0';
            for(auto d : _diags)
                if(occ == d[i])
                    new_diags.push_back(d);
            _diags = new_diags;
            i++;
        }
        return _diags[0];
    };
    oxygen = rating(diags, true);
    co2 = rating(diags, false);
    cout << stol(oxygen, 0, 2) * stol(co2, 0, 2) << endl;
    return 0;
}