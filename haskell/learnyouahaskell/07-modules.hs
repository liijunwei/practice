-- A Haskell module is a collection of related functions, types and typeclasses.

-- A Haskell program is a collection of modules where the main module loads up the other modules and then uses the functions defined in them to do something.

-- The syntax for importing modules in a Haskell script is import <module name>.

import Data.List
-- selectively import
import Data.List (nub, sort)
import Data.List hiding (nub)
import qualified Data.Map as M
import Data.Function (on)

numUniques :: (Eq a) => [a] -> Int
numUniques = length . nub

-- :m + Data.List
-- :m + Data.List Data.Map Data.Set

-- https://downloads.haskell.org/ghc/latest/docs/libraries/
-- A great way to pick up new Haskell knowledge is to just click through the standard library reference and explore the modules and their functions.
-- You can also view the Haskell source code for each module. Reading the source code of some modules is a really good way to learn Haskell and get a solid feel for it.

findKey key xs = snd . head . filter (\(k,v) -> key == k) $ xs
phoneBook = [("betty","555-2938"),("bonnie","452-2928"),("patsy","493-2928"),("lucille","205-2928"),("wendy","939-8282"),("penny","853-2492")]

-- intersperse 穿插
intersperse '.' "MONKEY"
intersperse 0 [1,2,3,4,5,6]

-- 3x2 + 5x + 9
-- 10x3 + 9
-- 8x3 + 5x2 + x - 1
map sum $ transpose [[0,3,5,9],[10,0,0,9],[8,5,1,-1]]

and $ map (>4) [5,6,7,8]
nub $ map (==4) [4,4,4,3,4]

any (==4) [2,3,5,6,1,4]
all (>4) [6,9,10]

any (`elem` ['A'..'Z']) "HEYGUYSwhatsup"
all (`elem` ['A'..'Z']) "HEYGUYSwhatsup"

let values = [-4.3, -2.4, -1.2, 0.4, 2.3, 5.9, 10.5, 29.1, 5.3, -2.4, -14.5, 2.9, 2.3]
groupBy (\x y -> (x > 0) == (y > 0)) values
groupBy (\x y -> (x > 0) && (y > 0) || (x <= 0) && (y <= 0)) values

-- unclear
-- :t on
groupBy ((==) `on` (> 0)) values
