#include <iostream>
#include <fstream>
using namespace std;

int main(void) {
    const int ITEMS = 4, ELEMENTS = 3;
    int mask[ITEMS] = {0, -1, -2, -3};
    int sliding_window[ITEMS] = {0};
    int last = 0, inc = 0;
    ifstream file; string line; 
    file.open("input.txt");
    while(getline(file, line)){
        for(int i = 0; i < ITEMS; i++){
            mask[i]++;
            mask[i] %= ITEMS;
            if(mask[i] > 0) sliding_window[i] += stoi(line);
            if(mask[i] == ELEMENTS){
                if(last && sliding_window[i] > last) inc++;
                last = sliding_window[i];
                sliding_window[i] = 0;
            }        
        }
    }
    cout << inc << endl;
    return 0;
}