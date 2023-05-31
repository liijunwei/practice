-- ghc --make 09-reverse-words.hs -o a.out
-- runhaskell 09-reverse-words.hs

reverseWords :: String -> String
reverseWords = unwords . map reverse . words -- I forgot what does this mean, roughly function composition
                                             -- http://learnyouahaskell.com/higher-order-functions#composition

-- words: split string to array of words
-- unwords: convert array to string

main = do
  line <- getLine
  if null line
    then return ()
    else do
      putStrLn $ reverseWords line
      main -- this loop looks strange to me...
