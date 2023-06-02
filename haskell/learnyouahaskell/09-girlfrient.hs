-- runhaskell 09-girlfrient.hs

import System.IO
import Data.Char

-- main = do
--   handle <- openFile "girlfriend.txt" ReadMode
--   contents <- hGetContents handle
--   putStr contents
--   hClose handle

-- :t ReadMode
-- :t openFile
-- :info FilePath

-- main = do
--   withFile "girlfriend.txt" ReadMode (\handle -> do
--     contents <- hGetContents handle
--     putStr contents) -- this right parenthese can not go to next line...

-- runhaskell 09-girlfrient.hs && cat girlfriendcaps.txt
main = do
  contents <- readFile "girlfriend.txt"
  writeFile "girlfriendcaps.txt" $ map toUpper contents
