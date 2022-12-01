#include <iostream>
#include <fstream>
using namespace std;

int main(void) {
    int actual, last = 0, inc = 0;
    ifstream file; string line;
    file.open("input.txt");
    while(getline(file, line)){
        actual = stoi(line);
        if(last && actual > last) inc++;
        last = actual;
    }
    cout << inc << endl;
    return 0;
}