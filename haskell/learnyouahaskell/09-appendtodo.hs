-- runhaskell 09-appendtodo.hs && echo ---- && cat todo.txt

-- import System.IO

-- main = do
--   todoItem <- getLine
--   appendFile "todo.txt" $ todoItem ++ "\n"


import System.IO
import System.Directory
import Data.List

main = do
  handle <- openFile "todo.txt" ReadMode
  (tempName, tempHandle) <- openTempFile "." "temp"
  contents <- hGetContents handle
  let todoTasks = lines contents
      numberedTask = zipWith (\n line -> show n ++ " - " ++ line) [0..] todoTasks
  putStrLn "these are your TODO items:"
  putStr $ unlines numberedTask
  putStrLn "which one do you wantto delete?"
  numberString <- getLine
  let number = read numberString
      newTodoItems = delete (todoTasks !! number) todoTasks
  hPutStr tempHandle $ unlines newTodoItems
  hClose handle
  hClose tempHandle
  removeFile "todo.txt"
  renameFile tempName "todo.txt"

-- it's better practice to use openTempFile so you know you're probably not overwriting anything.
