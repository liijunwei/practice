const todoList = [];

function addTodo() {
  const inputElememt = document.querySelector('.js-name-input');
  const name = inputElememt.value;
  if (name !== '' || name.trim() !== '') {
    todoList.push(name);
  }
  console.log(todoList);
  inputElememt.value = ''
}

function handleAddKeydown(event) {
  if (event.key === 'Enter') {
    addTodo()
  }
}
