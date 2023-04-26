-- 好难懂

applyTwice :: (a -> a) -> a -> a
applyTwice f x = f (f x)

-- applyTwice (+2) 1
-- applyTwice (++ " HAHA") "HEY"
-- applyTwice ("HAHA " ++) "HEY"
-- applyTwice (multThree 2 2) 9 -- 4 * (4 * 9)
-- applyTwice (3:) [1]

zipWith' :: (a -> b -> c) -> [a] -> [b] -> [c]
zipWith' _ [] _ = []
zipWith' _ _ [] = []
zipWith' f (x:xs) (y:ys) = f x y : zipWith' f xs ys

-- zipWith' (+) [4,2,5,6] [2,6,2,3]
-- zipWith' max [6,3,2,1] [7,3,1,5]
-- zipWith' (++) ["foo ", "bar ", "baz "] ["fighters", "hoppers", "aldrin"]
-- zipWith' (*) (replicate 5 2) [1..]
-- this one is hard to understand...
-- zipWith' (zipWith' (*)) [[1,2,3],[3,5,6],[2,3,4]] [[3,2,2],[3,4,5],[5,4,3]]

-- signature is hard to understand...
-- flip' :: (a -> b -> c) -> (b -> a -> c)
-- flip' f = g
--     where g x y = f y x

flip' :: (a -> b -> c) -> b -> a -> c
flip' f y x = f x y

-- zip [1,2,3,4,5] "hello"
-- flip' zip [1,2,3,4,5] "hello"
-- zipWith (flip' div) [2,2..] [10,8,6,4,2]

-- Maps and filters(legendary)
-- :t map    -- takes a mapping function and a list, return a new list
-- :t filter -- takes a predicate function and a list, return a new list

-- map (+3) [1,5,3,1,6]
-- filter even [1..10]

quicksort :: (Ord a) => [a] -> [a]
quicksort [] = []
quicksort (x:xs) =
    let smallerSorted = quicksort (filter (<=x) xs)
        biggerSorted = quicksort (filter (>x) xs)
    in  smallerSorted ++ [x] ++ biggerSorted
