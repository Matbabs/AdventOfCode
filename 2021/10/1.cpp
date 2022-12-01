#include <iostream>
#include <fstream>
#include <vector>
#include <map>
#include <stack>
#include <algorithm>
using namespace std;

int main(void) {
    ifstream file; string line;
    file.open("input.txt");
    vector<char> open = {'(','[', '{', '<'}, close = {')',']', '}', '>'};
    map<char, int> pts = {{close[0], 3},{close[1], 57},{close[2], 1197},{close[3], 25137}};
    int res = 0;
    while(file >> line){
        stack<char> toClose;
        for(auto c : line)
            if(find(open.begin(), open.end(), c) != open.end()) toClose.push(c);
            else if(toClose.size()){
                char toCls = toClose.top(); toClose.pop();
                int actualCls = find(close.begin(), close.end(), c) - close.begin();
                if(toCls != open[actualCls]) res += pts[c];
            }
    }
    cout << res << endl;
    return 0;
}