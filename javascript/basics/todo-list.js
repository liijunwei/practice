const todoList = [
  {name: 'task1', dueDate: '2024-11-30'},
  {name: 'task2', dueDate: '2024-12-1'},
];

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
    const html = `
    <p class="css-item">
      ${item.name} ${item.dueDate}
      <button class="js-delete-button" onclick="handleDelete(${index})">Delete</button>
    </p>`
    todoHTML += html
  })

  document.querySelector('.js-todo-list').innerHTML = todoHTML
}

function handleDelete(index) {
  todoList.splice(index, 1);
  renderTodoList();
}

getDefaultDueDate()
function getDefaultDueDate() {
  const today = new Date();

  const tomorrow = new Date(today);
  tomorrow.setDate(today.getDate() + 1);

  console.log(formatDate(tomorrow))

  return formatDate(tomorrow);
}

// Format the date as "yyyy-MM-dd"
function formatDate(date) {
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0'); // Adding 1 since months are zero-based
  const day = String(date.getDate()).padStart(2, '0');
  const formattedDate = `${year}-${month}-${day}`;

  return formattedDate
}
