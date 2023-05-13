-- http://learnyouahaskell.com/types-and-typeclasses

-- :: is read as "has type of" ***
:t 1
:t "hello" -- a list of characters
:t True
:t [1,2,3]

-- Unlike lists, each tuple length has its own type.
:t ("junwei", "male", "programmer", False, 67.5)
:t (True, 'a')
:t ('a','b','c')

removeNonUppercase :: String -> String
removeNonUppercase st = [ c | c <- st, c `elem` ['A'..'Z']]

-- Unlike Java or Pascal, Haskell has type inference.
-- If we write a number, we don't have to tell Haskell it's a number.
-- It can infer that on its own, so we don't have to explicitly write out the types of our functions and expressions to get things done.

-- However, **understanding the type system is a very important part of learning Haskell.**

-- A type is a kind of label that every expression has.
-- It tells us in which category of things that expression fits.

-- Functions also have types
let tripleMe x = x * 3
:t tripleMe

-- how do we write out the type of a function that takes several parameters?

-- The parameters are separated with -> and there's no special distinction between the parameters and the return type.
-- **The return type is the last item in the declaration** and the parameters are the first three.
-- Later on we'll see why they're all just separated with -> instead of having some more explicit distinction between the return types and the parameters like Int, Int, Int -> Int or something.
addThree :: Int -> Int -> Int -> Int
addThree x y z = x + y + z

-- a can be of any type.
:t head
-- head :: [a] -> a

-- Functions that have type variables are called polymorphic functions. 多态函数
-- The type declaration of head states that it takes a list of any type and returns one element of that type.

-- ### Typeclasses 101

-- A typeclass is a sort of interface that defines some behavior. ***
-- If a type is a part of a typeclass, that means that it supports and implements the behavior the typeclass describes.
-- typeclass 类似于java里的interface，但是更好(better) (不明白...)

:t (==)
-- (==) :: Eq a => a -> a -> Bool
-- Everything before the => symbol is called a **class constraint**.
-- the equality function takes any two values that are of the same type and returns a Bool.
-- The type of those two values must be a member of the Eq class (this was the class constraint).
-- Q: member of 和 instance of 有区别吗？

:t elem
-- elem :: (Foldable t, Eq a) => a -> t a -> Bool

-- Some basic typeclasses:
    -- `Eq` is used for types that support equality testing.
    -- `Ord` is for types that have an ordering.
    -- Members of `Show` can be presented as strings.
    -- `Read` is sort of the opposite typeclass of Show
    -- `Enum` members are sequentially ordered types — they can be enumerated
    -- `Bounded` members have an upper and a lower bound.
    -- `Num` is a numeric typeclass.

:t compare

read "1.4" :: Float
read "True" :: Bool

-- successors and predecesors
succ 'B'
pred 'B'

