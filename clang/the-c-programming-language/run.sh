# echo $HOME/practice/clang/the-c-programming-language/run.sh >> $HOME/.zshrc
function run(){
  git add .

  local filename=$1
  gcc -g $filename
  local exitcode_compile=$?
  shift

  local params=$@
  ./a.out $params
  local exitcode_execute=$?

  echo
  echo "-----------------"
  echo "compile result: $exitcode_compile" # TODO add filename info
  echo "execute result: $exitcode_execute" # TODO add filename info
}
