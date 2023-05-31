-- when a function changes state, we say that the function has side-effects
-- he only thing a function can do in Haskell is give us back some result based on the parameters we gave it.
-- If a function is called two times with the same parameters, it has to return the same result.

-- 不存在绝对没有副作用的程序，因为真实的世界里，副作用就是无处不在的
-- In order to tell us what it calculated, it has to change the state of an output device

-- haskell里有一种机制可以将函数pure的部分和non-pure的部分区分开来
-- Do not despair, all is not lost. It turns out that Haskell actually has a really clever system for dealing with functions that have side-effects that neatly separates the part of our program that is pure and the part of our program that is impure, which does all the dirty work like talking to the keyboard and the screen.

-- 09-hellow-world.hs

-- :t putStrLn
-- putStrLn takes a string and returns an I/O action that has a result type of ()
-- An I/O action is something that, when performed, will carry out an action with a side-effect (that's usually either reading from the input or printing stuff to the screen) and will also contain some kind of return value inside it.
-- Printing a string to the terminal doesn't really have any kind of meaningful return value, so a dummy value of () is used.

-- :t ()

-- Q: when will an I/O action be performed?
-- A: An I/O action will be performed when we give it a name of main and then run our program.

-- :t getLine
-- :info IO

-- Q: how haskell manages to neatly separate the pure and impure part of our code
--
-- you can think of an IO action as a box with little feet that will go out into real world and do something there and maybe bring back some data
-- once it's feched that data for you, the only way to open the box and get the data inside it is to use `<-` construct
-- and when we're taking data out of an IO action, we can only take it out when we're inside another IO action
--
-- getLine is in a sense impure because its result value is not guaranteed to be the same when performed twice
-- (which means it has state)

-- tainted 受到污染
--
-- we temporarily un-taint the data inside an I/O action when we bind it to a name

-- invalid statement, because String ++ IO action is invalid
--
-- nameTag = "Hello, my name is " ++ getLine
--
-- we need to bind the return value of `getLine` to a variable and then use it

-- 看完这个例子后，我理解了IOaction的return value的含义
-- foo <- putStrLn "Hello, what's your name?"
-- foo
-- :t foo
-- :t putStrLn
-- :t getLine

-- in a do block, the last action cannot be bound to a name like the first two were

-- I/O actions will only be performed when they are given a name of main or when they're inside a bigger I/O action that we composed with a do block.

-- haskell 里的每个if都必须要有一个对应的else分支，因为每个表达式都要有值
-- Remember that in Haskell, every if must have a corresponding else because every expression has to have some sort of value.


-- TODO not understand the meaning of `sequence` and  `mapM`



