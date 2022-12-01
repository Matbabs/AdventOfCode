#include <iostream>
#include <fstream>
#include <vector>
#include <set>
#define vc vector
#define pr pair
#define st set
#define ifs ifstream
#define str string
#define szt size_t
#define inf "input.txt"
#define br break
#define fr for
#define wh while
#define rt return
#define pb push_back
#define ins insert
#define cr clear
#define ln length
#define sz size
#define gl getline
#define op open
#define er erase
#define fd find
#define bg begin
#define sbs substr
#define sti stoi
#define ss second
#define ff first
using namespace std;

//               .~~~~~~.                                                                         .~~~~~~
//~~~~~~~~~~~~~~'        '~~~~~~~~~~~~~~~~~.                              .~~~~~~~~~~~~~~~~~~~~~~'
//                                          '~~~~~~~~~~~~~~.      .~~~~~~'
//                                                          '~~~~'                                       
//                    0                        0                          0                              
//                  0                                       0                0                           
//                0              0                           0                                           
//                             0                           0                                            
vc<int>       /**/
spt(str s,   /*o*/    /**/                                   /**/ 
str del){vc<int>      /*o*/  arr;szt ps=0;str tok;wh           /**/ 
((ps=s.fd(del))!=    /*oo*/  str::npos){tok=s.sbs(0,ps);         /*o*/ 
arr.pb(sti(tok));     /*oo*/ s.er(0,ps + del.ln());}arr.pb(       /*o*/ 
sti(s));rt arr;}int   /*oo*/  main(void){const int sze=5;ifs      /*oo*/ 
fl;str ln,in;vc<int> nbs;vc<vc<vc<pr<int,bool>>>> gds;vc<vc<pr<int,/*ooo*/bool>>> gd(sze,vc<pr<int,
bool>>(sze));st<int> gsw;int/*OO*/l=0,i=0,rn,rg,rs=0/**/;fl.op(inf);/*ooo*/wh(gl(fl,ln)){if(l==0)
    nbs=spt(ln,",");if(l !=0&&ln.ln()!=0){vc<int> numb=/**/spt(ln," ");/*ooo*/fr(int j=0;j<numb.sz();j++)
    gd[i][j]={numb[j],false};i++;if(i==sze){gds.pb(gd)/**/;i=0;}}l++;}wh/*ooo*/(nbs.sz()){bool iWg=false,
    iLG=false;int n=nbs[0],nbg=0;fr(auto& gd : gds)/**/{fr(int i=0;i</*ooo*/ sze;i++)fr(int j=0;j<sze;j++)
    if(n==gd[i][j].ff)gd[i][j].ss=true;auto cLW=/*o*/[&](bool md){fr(/*ooo*/int i=0;i<sze;i++){iWg=true;fr
  (int j=0;j<sze;j++)if(md&&!gd[i][j].ss||!md&&/*oo*/!gd[j][i].ss){iWg/*ooo*/=false;br;}if(iWg){gsw.ins(nbg);
if(gsw.sz()==gds.sz()){iLG=true;br;}}}rt iLG;};if/*oo*/(cLW(true))br;/*ooo*/if(cLW(false))br;if(nbg<gds.sz()
-1)nbg++;}if(iLG){nbs.cr();rn=n;rg=nbg;} else nbs.er(/*oo*/nbs.bg());}fr/*ooo*/(int i=0;i<sze;i++)fr(int 
j=0;j<sze;j++)if(!gds[rg][i][j].ss)rs+=gds[rg][i][j].ff/*oo*/;cout<</*oooo*/rs*rn<<endl;rt 0;}
    /*ooooo*/      /*ooooo*/          /*oo*/        /*oo*/       /*ooo*/ 
    /*ooooo           oooooo*/        /*ooo*/     /*ooo*/      /*oooo*/ 
     /*ooooo*/          /*oooo*/     /*oooo>      <oooo*/      /*oooo*/ 
        /*oooo*/          /*oooooo   ooooooo^^^^^^oooooo     ooooo*/ 
          /*ooooooo*/      /*ooooooooooooooooooooooooooo  ooooooo*/
              /*oooooooo       ooooooooooooooooooooooooooooo*/
                   /*oooooooooooooooooooo   oooooo   ooooo*/
                        /*oooooooooooooo ### oooo ### ooo*/
                                 /*ooooo ### oooo ### oo*/
                                  /*ooooo   oooooo   oo*/
                                   /*oooooooooooooooo*/
                                     /*oooooooooooo*/
                                        /*oooooo*/




