const todoList = []
todoList.push({name: 'task1', dueDate: '2024-11-30'})
todoList.push({name: 'task2', dueDate: '2024-12-1'})
todoList.push({name: 'task2', dueDate: '2024-12-1'})

setRandomDefaultDueDate()
renderTodoList()

// 1. save the data
// 2. generate the HTML
// 3. make it interactive
function addTodo() {
  const nameElememt = document.querySelector('.js-name-input');
  const name = nameElememt.value;

  const dateElememt = document.querySelector('.js-due-date-input');
  const date = dateElememt.value;

  if (name !== '' && name.trim() !== '') {
    todoList.push({name: name, dueDate: date});
  }

  console.log(todoList);
  nameElememt.value = ''
  renderTodoList()
}

function handleAddKeydown(event) {
  if (event.key === 'Enter') {
    addTodo()
  }
}

function renderTodoList() {
  let todoHTML = ''
  todoList.forEach((item, index) => {
    const {name, dueDate} = item;

    const html = `
      <div class="grid-item">${name}</div>
      <div class="grid-item">${dueDate}</div>
      <button class="delete-todo-button" onclick="handleDelete(${index})">Delete</button>
    `
    todoHTML += html
  })

  document.querySelector('.js-todo-list').innerHTML = todoHTML
}

function handleDelete(index) {
  todoList.splice(index, 1);
  renderTodoList();
}

function setRandomDefaultDueDate() {
  const today = new Date();
  const tomorrow = new Date(today);
  tomorrow.setDate(today.getDate() + getRandomNumber(1,3));
  const defaultDueDate = formatDate(tomorrow);

  document.querySelector('.js-due-date-input').value = defaultDueDate
}

// Format the date as "yyyy-MM-dd"
function formatDate(date) {
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0'); // Adding 1 since months are zero-based
  const day = String(date.getDate()).padStart(2, '0');
  const formattedDate = `${year}-${month}-${day}`;

  return formattedDate
}

function getRandomNumber(min, max) {
  return Math.floor(Math.random() * (max - min + 1)) + min;
}
