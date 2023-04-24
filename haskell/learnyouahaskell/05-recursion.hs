-- Recursion is actually a way of defining functions in which the function is applied inside its own definition.

-- example: fibonacci sequence

-- example: factorial

-- edge condition

-- Recursion is important to Haskell because unlike imperative languages, you do computations in Haskell by declaring what something is instead of declaring how you get it.
-- **That's why there are no while loops or for loops in Haskell** and instead we many times have to use recursion to declare what something is.

-- example: maximum
-- example: replicate
-- example: take
-- example: reverse
-- example: repeat
-- example: zip
-- example: elem
-- example: quicksort

maximum' :: (Ord a) => [a] -> a
maximum' [] = error "maximum of empty list"
maximum' [x] = x
maximum' (x:xs)
    | x > maxTail = x
    | otherwise = maxTail
    where maxTail = maximum' xs

-- maximum' [2,5,1]

replicate' :: (Num i, Ord i) => i -> a -> [a]
replicate' n x
    | n <= 0    = []
    | otherwise = x:replicate' (n-1) x

take' :: (Num i, Ord i) => i -> [a] -> [a]
take' n _
    | n <= 0 = []
take' _ [] = []
take' n (x:xs) = x : take' (n-1) xs

reverse' :: [a] -> [a]
reverse' [] = []
reverse' (x:xs) = reverse' xs ++ [x]

repeat' :: a -> [a]
repeat' x = x:(repeat' x) -- never ending

zip' ::[a] -> [b] -> [(a,b)]
zip' _ [] = []
zip' [] _ = []
zip' (x:xs) (y:ys) = (x,y):zip' xs ys


elem' :: (Eq a) => a -> [a] -> Bool
elem' a [] = False
elem' a (x:xs)
    | a == x = True
    | otherwise = a `elem'` xs

-- 精妙...
quicksort :: (Ord a) => [a] -> [a]
quicksort [] = []
quicksort (x:xs) = -- x is so-called pivot
    let smallerSorted = quicksort [a | a <- xs, a <= x]
        biggerSorted  = quicksort [a | a <- xs, a > x]
    in smallerSorted ++ [x] ++ biggerSorted

-- Thinking recursively
