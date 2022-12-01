#include <iostream>
#include <fstream>
using namespace std;

int main(void) {
    int actual, x = 0, d = 0, aim = 0, value;
    ifstream file; string line, command;
    file.open("input.txt");
    while(getline(file, line)){
        command = line.substr(0, line.find(" "));
        value = stoi(line.substr(line.find(" "), line.size()));
        if(command == "up") aim -= value;
        if(command == "down") aim += value;
        if(command == "forward") {
            x += value;
            d += aim * value;
        }
    }
    cout << x * d << endl;
    return 0;
}