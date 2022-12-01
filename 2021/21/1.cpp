#include <iostream>
#include <fstream>
using namespace std;

const int BOARD = 11, DETER_DICE = 101, SCORE = 1000;

int throw_dice(int &dice, int &rolls){
    dice++;
    dice = (dice % DETER_DICE) + dice / DETER_DICE;
    rolls++;
    return dice;
}

void move_pawn(int &pawn, int &score, int moves){
    pawn += moves;
    while(pawn >= BOARD) pawn = (pawn % BOARD) + pawn / BOARD;
    score += pawn;
}

int main(void) {
    ifstream file; string line;
    const string DELIM = "position: ";
    int p1_pawn = -1, p2_pawn = -1, dice = 0, p1_score = 0, p2_score = 0, rolls = 0;
    bool p_turn = true;
    file.open("input.txt");
    while(getline(file, line)){
        if(p1_pawn < 0) p1_pawn = stoi(line.substr(line.find(DELIM) + 9, 3));
        else if(p2_pawn < 0) p2_pawn = stoi(line.substr(line.find(DELIM) + 9, 3));
    }
    while(p1_score < SCORE && p2_score < SCORE){
        int moves = throw_dice(dice, rolls) + throw_dice(dice, rolls) + throw_dice(dice, rolls);
        if(p_turn) move_pawn(p1_pawn, p1_score, moves);
        else move_pawn(p2_pawn, p2_score, moves);
        p_turn = !p_turn;
    }
    cout << (p_turn ? p1_score : p2_score) * rolls << endl;
    return 0;
}