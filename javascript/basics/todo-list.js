const todoList = [];

function addTodo() {
  const inputElememt = document.querySelector('.js-name-input');
  const name = inputElememt.value;
  todoList.push(name);

  console.log(todoList);

  inputElememt.value = ''
}
