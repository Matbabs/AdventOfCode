#include <iostream>
#include <fstream>
#include <vector>
#define sbs substr
#define sti stoi
#define in "input.txt"
#define op open
#define sz size
#define fd find
using namespace std;

//                               _    _     __          ___  __
//                              | |  | |   /\ \        / / |/ /
//                              | |__| |  /  \ \  /\  / /| ' / 
//                              |  __  | / /\ \ \/  \/ / |  <  
//                              | |  | |/ ____ \  /\  /  | . \ 
//                              |_|  |_/_/    \_\/  \/   |_|\_\ 

int main(void){const int                                            MPS=1000;ifstream fl;
    string st,_,ed;int x1,                                      y1,x2,y2,sum=0;fl.op(in) 
        ;vector<vector<int>>                                  ts(MPS,vector<int>( 
            MPS,0));while (fl>>                           st>>_>>ed){x1=sti( 
                st.sbs(0,st.fd(',')));                y1=sti(st.sbs(st.fd(',')+   
                    1,st.sz()));x2= sti(              ed.sbs(0,ed.fd(','))); 
                        y2=sti(ed.sbs(ed.fd(       ',')+ 1,ed.sz()));   
                            if(x1==x2)for(int i=min(y1,y2); i <= max(y1,y2); 
                                i++)ts[i][x1]++;if(y1==y2)for(int i=min(x1  
                                    ,x2);i <= max(x1,x2);i++)ts[y1][i]++; 
                                    if(/*#O*/x1!=x2&&y1!=y2){ts[y1][x1]++; 
                                while(x1!=x2&&y1!=y2){if(y1!=y2)(y2-y1)>0?y1++:y1--;if(x1!=x2)(x2-x1) 
                                        >0?x1++:x1--;ts[y1][x1]++;}}}for(auto ln : ts)for(auto c : ln)if(c > 1) 
                                                                sum++;cout<<sum<<endl;return 0;} 
//                 ___                                                      ___
//                 \  \                                                    /  /
//                  `. `.                                                ,' ,' 
//                    \__\_____ ______ ______ ______ ______ ______ _____/__/   
//                      |______|______|______|______|______|______|______|     

