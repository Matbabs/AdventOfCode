mkdir $1
cp _template/1.go $1
cd $1
wget --load-cookies=../cookie.txt https://adventofcode.com/2016/day/$1/input
mv input input.txt