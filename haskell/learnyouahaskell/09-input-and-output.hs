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


