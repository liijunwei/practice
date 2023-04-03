-- http://learnyouahaskell.com/types-and-typeclasses

-- :: is read as "has type of"
:t 1
:t "hello" -- a list of characters
:t True
:t [1,2,3]

-- Unlike lists, each tuple length has its own type.
:t ("junwei", "male", "programmer", False, 67.5)
:t (True, 'a')
:t ('a','b','c')


-- Unlike Java or Pascal, Haskell has type inference.
-- If we write a number, we don't have to tell Haskell it's a number.
-- It can infer that on its own, so we don't have to explicitly write out the types of our functions and expressions to get things done.

-- However, understanding the type system is a very important part of learning Haskell.

-- A type is a kind of label that every expression has.
-- It tells us in which category of things that expression fits.

-- Functions also have types
let tripleMe x = x * 3
:t tripleMe
