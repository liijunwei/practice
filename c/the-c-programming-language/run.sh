# echo $HOME/practice/c/the-c-programming-language/run.sh >> $HOME/.zshrc
function run(){
  local filename=$1;
  gcc $filename && ./a.out
}