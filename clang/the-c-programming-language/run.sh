# echo $HOME/practice/c/the-c-programming-language/run.sh >> $HOME/.zshrc
function run(){
  git add .

  local filename=$1
  gcc -g $filename
  local exitcode_compile=$?
  local params=$@

  ./a.out ${params/$filename}
  local exitcode_execute=$?

  echo
  echo "-----------------"
  echo "compile result: $exitcode_compile" # TODO add filename info
  echo "execute result: $exitcode_execute" # TODO add filename info
}
