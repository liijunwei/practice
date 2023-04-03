doubleMe x = x + x
tripleMe x = x * 3
-- doubleUs x y = x*2 + y*2
doubleUs x y = doubleMe x + doubleMe y

-- doubleSmallNumber x = if x > 100
--                         then x
--                         else x*2

doubleSmallNumber' x = (if x > 100 then x else x*2) + 1

(succ 9) + (max 5 4) + 1

-- List
let lostNumbers = [4,8,15,16,23,42]
lostNumbers = [4,8,15,16,23,42]

[1,2,3,4] ++ [9,10,11,12]
"abc" ++ "def"

['w','o'] ++ ['o','t']

'A':" SMALL CAT"
5:[1,2,3,4,5]
-- [1,2,3] is actually just syntactic sugar for 1:2:3:[]

[1..20]
['a'..'z']
[0.1, 0.3 .. 1]
take 10 (cycle [1,2,3])

[x*2 | x <- [1..10]] -- something like pipeline

-- Tuples(元组)

-- Another key difference is that they don't have to be homogenous. Unlike a list, a tuple can contain a combination of several types.
-- tuple里的数据不必是 同质(homogenous) 的

[[1,2],[8,11],[4,5]]
[[1,2],[8,11,5],[4,5]]

[(1,2),(8,11),(4,5)]
-- error raised
-- It's telling us that we tried to use a pair and a triple in the same list, which is not supposed to happen.
[(1,2),(8,11,5),(4,5)]
[(1,2),("One",2)]

("Christopher", "Walken", 55)

-- Use tuples when you know in advance how many components some piece of data should have.
-- Tuples are much more rigid because each different size of tuple is its own type, so you can't write a general function to append an element to a tuple


-- ruby has same api
zip [1,2,3,4,5] [5,5,5,5,5]
zip [1 .. 5] ["one", "two", "three", "four", "five"]
zip [5,3,2,6,2,7,2,5,4,6,6] ["im","a","turtle"]

let triangles = [ (a,b,c) | c <- [1..10], b <- [1..10], a <- [1..10] ]
-- add a condition that they all have to be right triangles
[ (a,b,c) | c <- [1..10], b <- [1..c], a <- [1..b], a^2 + b^2 == c^2]
[ (a,b,c) | c <- [1..10], b <- [1..c], a <- [1..b], a^2 + b^2 == c^2, a+b+c == 24]

-- 描述要做什么，不描述怎么做


