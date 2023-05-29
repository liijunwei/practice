-- ghc --make 09-hellow-world.hs -o a.out
-- ./a.out

-- main = putStrLn "hello, world"

-- echo "余则成" | ./a.out

main = do
  putStrLn "hello, what's your name?"
  name <- getLine
  putStrLn ("Hey " ++ name ++ ", you rock!!!")
