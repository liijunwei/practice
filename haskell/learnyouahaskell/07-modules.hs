-- A Haskell module is a collection of related functions, types and typeclasses.

-- A Haskell program is a collection of modules where the main module loads up the other modules and then uses the functions defined in them to do something.

-- The syntax for importing modules in a Haskell script is import <module name>.

import Data.List

numUniques :: (Eq a) => [a] -> Int
numUniques = length . nub

-- :m + Data.List
-- :m + Data.List Data.Map Data.Set

-- selectively import
import Data.List (nub, sort)


