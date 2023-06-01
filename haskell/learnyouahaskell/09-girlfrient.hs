-- runhaskell 09-girlfrient.hs

import System.IO

-- main = do
--   handle <- openFile "girlfriend.txt" ReadMode
--   contents <- hGetContents handle
--   putStr contents
--   hClose handle

-- :t ReadMode
-- :t openFile
-- :info FilePath

main = do
  withFile "girlfriend.txt" ReadMode (\handle -> do
    contents <- hGetContents handle
    putStr contents) -- this right parenthese can not go to next line...
