-- ghc --make 09-hellow-world.hs -o a.out
-- ./a.out

-- main = putStrLn "hello, world"

-- echo "余则成" | ./a.out

-- tellFortune doesn't know anything about IO
tellFortune :: String -> String
tellFortune name = "you'll be rick!!! " ++ name

-- By putting them together with do syntax, we glued them into one I/O action
main = do
  putStrLn "hello, what's your name?"
  name <- getLine -- perform the I/O action getLine and then bind its result value to name
                  -- getLine has a type of IO String, so name will have a type of String
  putStrLn $ "Read this carefully, because this is your future: " ++ tellFortune name

