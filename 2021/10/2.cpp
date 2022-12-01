#include <iostream>
#include <fstream>
#include <vector>
#include <map>
#include <stack>
#include <algorithm>
#define all(x) x.begin(), x.end()
using namespace std;

int main(void) {
    ifstream file; string line;
    file.open("input.txt");
    vector<char> open = {'(','[', '{', '<'}, close = {')',']', '}', '>'};
    map<char, int> pts = {{close[0], 1},{close[1], 2},{close[2], 3},{close[3], 4}};
    vector<long> scores;
    while(file >> line){
        stack<char> toClose;
        bool isCorrupted = false;
        for(auto c : line)
            if(find(all(open), c) != open.end()) toClose.push(c);
            else if(toClose.size()){
                char toCls = toClose.top(); toClose.pop();
                int actualCls = find(all(close), c) - close.begin();
                if(toCls != open[actualCls]){
                    isCorrupted = true;   
                    break;
                }
            }
        if(!isCorrupted){
            long sum = 0;
            while(toClose.size()){
                char toCls = toClose.top(); toClose.pop();
                int actualOpn = find(all(open), toCls) - open.begin();
                sum = sum * 5 + pts[close[actualOpn]];
            }
            scores.push_back(sum);
        }
    }
    sort(all(scores));
    cout << scores[scores.size() / 2] << endl;
    return 0;
}