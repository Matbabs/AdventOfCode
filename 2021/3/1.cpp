#include <iostream>
#include <fstream>
using namespace std;

int main(void) {
    ifstream file; string line, gamma = "", epsilon = "";
    int bits[12][2];
    file.open("input.txt");
    while(getline(file, line))
        for(int i = 0; i < line.size(); i++)
            if(line[i] == '0') bits[i][0]++; else bits[i][1]++; 
    for(int i = 0; i < 12; i++){
        gamma += bits[i][0] > bits[i][1] ? '0' : '1';
        epsilon += bits[i][1] < bits[i][0] ? '1' : '0';
    }
    cout << stol(gamma, 0, 2) * stol(epsilon, 0, 2) << endl;
    return 0;
}