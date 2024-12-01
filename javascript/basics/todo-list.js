const todoList = []
todoList.push({ name: 'task1', dueDate: '2024-11-30' })
todoList.push({ name: 'task2', dueDate: '2024-12-01' })
todoList.push({ name: 'task2', dueDate: '2024-12-01' })

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
    todoList.push({ name: name, dueDate: date });
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
  todoList.forEach(({ name, dueDate }, index) => {
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
  tomorrow.setDate(today.getDate() + getRandomNumber(1, 3));
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

// playground1()
function playground1() {
  const a1 = [1, 2, 3];
  const a2 = a1.slice(); // js shallow copy
  console.log('a1', a1)
  console.log('a1', a2)
  console.log("====")
  a2.push(4)
  console.log('a1', a1)
  console.log('a1', a2)
  console.log("====")

  const [foo, bar, fizz] = ['boo', 'bar', 'fizz', 'buzz'];
  console.log('foo', foo)
  console.log('bar', bar)
  console.log('fizz', fizz)
}

// playground2()
function playground2() {
  const f1 = function () { // anonymous function
    console.log("func1")
  }
  console.log(f1)
  console.log(typeof f1)
  f1()

  const f2 = () => {
    console.log("func2")
  }
  console.log(f2)
  console.log(typeof f2)
  f2()

  const o1 = {
    f1: f1,
    f2: f2,
  }
  console.log(o1)
  o1.f1()


  const o2 = {
    f1,
    f2,
  }
  console.log(o2)
  o2.f2()

  setTimeout(() => console.log("kapoom"), 1000) // async
  setTimeout(function () { console.log("kapoom") }, 1000) // async
  console.log("fire in the hole")

  setInterval(() => { console.log("ping") }, 3000);
}

// playground3()
function playground3() {
  const f1 = () => {
    console.log("arraw function")
  }

  f1()

  const o1 = {
    m1: () => {
      console.log('m1')
    }
  }

  o1.m1()
}

addEventListenerToDemoButton()
function addEventListenerToDemoButton() {
  const btn = document.querySelector('.demo-button');
  console.assert(btn !== null);

  const listener = () => {
    console.log("click1");
  }

  btn.addEventListener('click', listener); // best practice
  btn.removeEventListener('click', listener);

  btn.addEventListener('click', () => {
    console.log("click2");
  })
}
