-- runhaskell 09-demo.hs

-- main = do
--     a <- return "hell"
--     b <- return "yeah!"
--     putStrLn $ a ++ " " ++ b

-- main = do
--   let a = "hell"
--       b = "yeah"
--   putStrLn $ a ++ " " ++ b

-- main = do
--   putStr "Hey, "
--   putStr "I'm "
--   putStrLn "Andy!"

-- main = do   putChar 't'
--             putChar 'e'
--             putChar 'h'

-- main = do
--     c <- getChar
--     if c /= ' '
--         then do
--             putChar c
--             main
--         else return ()

-- import Control.Monad
-- main = do
--     c <- getChar
--     when (c /= ' ') $ do -- when is a function
--         putChar c
--         main

-- echo -e "1\n2\n3" | runhaskell 09-demo.hs
-- main = do
--     a <- getLine
--     b <- getLine
--     c <- getLine
--     print [a,b,c]

-- :t sequence
-- main = do
--     rs <- sequence [getLine, getLine, getLine]
--     print rs


-- import Control.Monad
-- import Data.Char

-- main = forever $ do
--     putStr "Give me some input: "
--     l <- getLine
--     putStrLn $ map toUpper l

-- echo -e "red\nyellow\nblue\ngrey" | runhaskell 09-demo.hs
-- import Control.Monad
-- main = do
--     colors <- forM [1,2,3,4] (\a -> do
--         putStrLn $ "Which color do you associate with the number " ++ show a ++ "?"
--         color <- getLine
--         return color)
--     putStrLn "The colors that you associate with 1, 2, 3 and 4 are: "
--     mapM putStrLn colors

-- -- cat haiku.txt | runhaskell 09-demo.hs
-- import Control.Monad
-- import Data.Char

-- main = forever $ do
--   -- putStr "Give me some input: "
--   l <- getLine
--   putStrLn $ map toUpper l


-- cat haiku.txt | runhaskell 09-demo.hs
-- import Data.Char

-- main = do
--   contents <- getContents
--   putStr $ map toUpper contents


-- cat shortlines.txt | runhaskell 09-demo.hs
-- main = interact shortLinesOnly
-- main = do
--   contents <- getContents
--   putStr $ shortLinesOnly contents

-- shortLinesOnly :: String -> String
-- shortLinesOnly input =
--   let allLines = lines input
--       shortLines = filter (\line -> length line < 10) allLines
--       result = unlines shortLines
--   in result

-- :t interact
-- interact :: (String -> String) -> IO ()

-- cat shortlines.txt | runhaskell 09-demo.hs
-- 妙哇 composition
-- main = interact $ unlines . filter ((<10) . length) . lines

respondPalindromes :: String -> String
respondPalindromes contents = unlines $ map (\xs -> isPalindrome xs) (lines contents)
    where isPalindrome xs = if xs == reverse xs then "palindrome" else "not a palindrome"

-- putStrLn $ respondPalindromes "elephant\nABCBA\nwhatever"

-- runhaskell 09-demo.hs
main = interact respondPalindromes

