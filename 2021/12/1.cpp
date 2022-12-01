#include <iostream>
#include <fstream>
#include <vector>
#include <stack>
#include <set>
#include <map>
#include <algorithm>
#define all(x) x.begin(), x.end()
using namespace std;

int main(void) {
    const string START = "start", END = "end";
    ifstream file; string line;
    file.open("input.txt");
    map<string, vector<string>> graph;
    while(getline(file, line)){
        string s = line.substr(0, line.find('-'));
        string e = line.substr(line.find('-') + 1, line.size());
        graph[s].push_back(e);
        graph[e].push_back(s);
    }
    vector<vector<string>> paths(1, vector<string>(1, START)), completePaths;
    vector<string> path, new_path;
    string lastCave;

    while(paths.size()){
        path = paths[paths.size() - 1]; paths.pop_back();
        lastCave = path[path.size() - 1];
        for(auto cave : graph[lastCave]){
            new_path = {all(path)};
            new_path.push_back(cave);
            if(cave == END)
                completePaths.push_back(new_path);
            else if(all_of(all(cave), &::isupper) || find(all(path), cave) == path.end())
                paths.push_back(new_path);
        }
    }
    cout << completePaths.size() << endl;
    return 0;
}