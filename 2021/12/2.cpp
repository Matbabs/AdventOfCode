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
    vector<pair<vector<string>, bool>> paths(1, {vector<string>(1, START), true}), completePaths;
    pair<vector<string>, bool> path, new_path;
    string lastCave;

    while(paths.size()){
        path = paths[paths.size() - 1]; paths.pop_back();
        lastCave = path.first[path.first.size() - 1];
        for(auto cave : graph[lastCave])
            if(cave == END){
                new_path = {{all(path.first)}, path.second};
                new_path.first.push_back(cave);
                completePaths.push_back(new_path);
            }   
            else if(cave == START) continue;
            else if(all_of(all(cave), &::isupper) || find(all(path.first), cave) == path.first.end()){
                new_path = {{all(path.first)}, path.second};
                new_path.first.push_back(cave);
                paths.push_back(new_path);
            }
            else if(path.second){
                new_path = {{all(path.first)}, false};
                new_path.first.push_back(cave);
                paths.push_back(new_path);
            }
    }
    cout << completePaths.size() << endl;
    return 0;
}