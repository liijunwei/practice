# echo $HOME/practice/c/the-c-programming-language/run.sh >> $HOME/.zshrc
function run(){
  local filename=$1
  gcc -g $filename
  local exitcode_compile=$?

  ./a.out
  local exitcode_execute=$?

  echo
  echo "-----------------"
  echo "compile result: $exitcode_compile"
  echo "execute result: $exitcode_execute"
}
