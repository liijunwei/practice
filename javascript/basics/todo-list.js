const todoList = ['task1', 'task2'];
renderTodoList()

// 1. save the data
// 2. generate the HTML
// 3. make it interactive
function addTodo() {
  const inputElememt = document.querySelector('.js-name-input');
  const name = inputElememt.value;
  if (name !== '' && name.trim() !== '') {
    todoList.push(name);
  }
  console.log(todoList);
  inputElememt.value = ''
  renderTodoList()
}

function handleAddKeydown(event) {
  if (event.key === 'Enter') {
    addTodo()
  }
}

function renderTodoList(){
  let todoHTML = ''
  todoList.forEach(item => {
    const html = `<p>${item}</p>`
    todoHTML += html
  })

  document.querySelector('.js-todo-list').innerHTML = todoHTML
}
