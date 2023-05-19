-- A Haskell module is a collection of related functions, types and typeclasses.

-- A Haskell program is a collection of modules where the main module loads up the other modules and then uses the functions defined in them to do something.

-- The syntax for importing modules in a Haskell script is import <module name>.

-- module for
--    manipulating lists
--    concurrent programming
--    dealing with complex numbers
--    ...


import Data.List
-- selectively import
import Data.List (nub, sort)
import Data.List hiding (nub)
import Data.Function
import Data.Char
import qualified Data.Map as Map

numUniques :: (Eq a) => [a] -> Int
numUniques = length . nub

-- :m + Data.List
-- :m + Data.List Data.Map Data.Set

-- https://downloads.haskell.org/ghc/latest/docs/libraries/
-- A great way to pick up new Haskell knowledge is to just click through the standard library reference and explore the modules and their functions.
-- You can also view the Haskell source code for each module. Reading the source code of some modules is a really good way to learn Haskell and get a solid feel for it.

-- intersperse 穿插
-- intersperse '.' "MONKEY"
-- intersperse 0 [1,2,3,4,5,6]

-- 3x2 + 5x + 9
-- 10x3 + 9
-- 8x3 + 5x2 + x - 1
-- map sum $ transpose [[0,3,5,9],[10,0,0,9],[8,5,1,-1]]

-- and $ map (>4) [5,6,7,8]
-- nub $ map (==4) [4,4,4,3,4]

-- any (==4) [2,3,5,6,1,4]
-- all (>4) [6,9,10]

-- any (`elem` ['A'..'Z']) "HEYGUYSwhatsup"
-- all (`elem` ['A'..'Z']) "HEYGUYSwhatsup"

-- let values = [-4.3, -2.4, -1.2, 0.4, 2.3, 5.9, 10.5, 29.1, 5.3, -2.4, -14.5, 2.9, 2.3]
-- groupBy (\x y -> (x > 0) == (y > 0)) values
-- groupBy (\x y -> (x > 0) && (y > 0) || (x <= 0) && (y <= 0)) values

-- pretty neat
-- :t on
-- groupBy ((==) `on` (> 0)) values
-- groupBy (on (==) (> 0)) values

-- xs = [[5,4,5,4,4],[1,2,3],[3,5,4,3],[],[2],[2,2]]
-- sortBy (\x y -> length x `compare` length y) xs
-- sortBy (compare `on` length) xs

-- all isAlphaNum "bobby283"
-- isControl 'a'
-- isControl '\n'

-- hashmap
phoneBook = [("betty","555-2938"),("bonnie","452-2928"),("patsy","493-2928"),("lucille","205-2928"),("wendy","939-8282"),("penny","853-2492")]
findKey key xs = snd . head . Data.List.filter (\(k,_) -> key == k) $ xs
-- findKey "betty" phoneBook
-- findKey "betty1" phoneBook

findKey1 :: (Eq k) => k -> [(k,v)] -> Maybe v
findKey1 key [] = Nothing
findKey1 key ((k,v):xs) = if key == k
    then Just v
    else findKey1 key xs
-- findKey1 "betty" phoneBook
-- findKey1 "betty1" phoneBook
-- Q: any way to rewrite using guard clause ?

findKey2 :: (Eq k) => k -> [(k,v)] -> Maybe v
findKey2 key = Data.List.foldr (\(k,v) acc -> if key == k then Just v else acc) Nothing
-- 遍历时没有修改acc, 所以遍历结束以后，结果为 Just v 或者 Nothing
-- findKey2 "betty" phoneBook
-- findKey2 "betty1" phoneBook

-- import qualified Data.Map as Map

-- Map.fromList [("betty","555-2938"),("bonnie","452-2928"),("lucille","205-2928")]
-- Map.fromList [(1,2),(3,4),(3,2),(5,5)]

-- You should always use Data.Map for key-value associations unless you have keys that aren't part of the Ord typeclass.

-- Map.empty
-- Map.insert 3 100 Map.empty
-- Map.insert 5 600 (Map.insert 4 200 ( Map.insert 3 100  Map.empty))
-- Map.insert 5 600 . Map.insert 4 200 . Map.insert 3 100 $ Map.empty
-- Map.null Map.empty
-- Map.null $ Map.fromList [(2,3),(5,5)]

phoneBookMap = Map.fromList phoneBook
-- Map.member "betty" phoneBookMap
-- Map.lookup "betty" phoneBookMap
-- Map.lookup "bonnie" phoneBookMap

-- :t Map.map
-- :t Map.filter

-- toList phoneBookMap

anotherPhoneBook =
    [("betty","555-2938")
    ,("betty","342-2492")
    ,("bonnie","452-2928")
    ,("patsy","493-2928")
    ,("patsy","943-2929")
    ,("patsy","827-9162")
    ,("lucille","205-2928")
    ,("wendy","939-8282")
    ,("penny","853-2492")
    ,("penny","555-2111")
    ]

phonebookToMap :: (Ord k) => [(k,String)] -> Map.Map k String
phonebookToMap xs = Map.fromListWith (\number1 number2 -> number1 ++ ", " ++ number2) xs

-- Map.lookup "patsy" $ phonebookToMap anotherPhoneBook
-- Map.lookup "wendy" $ phonebookToMap anotherPhoneBook
-- Map.lookup "betty" $ phonebookToMap anotherPhoneBook
