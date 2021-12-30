# echo $HOME/practice/c/the-c-programming-language/run.sh >> $HOME/.zshrc
function run(){
  local filename=$1
  gcc -g $filename
  echo "========================"
  echo "ExitCode(Compile): $?"
  echo "========================"
  echo
  ./a.out
  echo
  echo "========================"
  echo "ExitCode(Execute): $?"
  echo "========================"
  echo
}
