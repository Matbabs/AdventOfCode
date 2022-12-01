#include <iostream>
#include <fstream>
using namespace std;

int main(void) {
    ifstream file; string line;
    file.open("input.txt");
    while(getline(file, line)){
        cout << line << endl;
    }
    return 0;
}