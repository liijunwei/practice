# echo $HOME/practice/c/the-c-programming-language/run.sh >> $HOME/.zshrc
function run(){
  local filename=$1
  gcc -g $filename
  echo "ExitCode(Compile): $?"
  ./a.out
  echo "ExitCode(Execute): $?"
  echo "========================"
  echo
}
