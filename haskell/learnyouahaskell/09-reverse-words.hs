-- ghc --make 09-reverse-words.hs -o a.out

reverseWords :: String -> String
reverseWords = unwords . map reverse . words

main = do
  line <- getLine
  if null line
    then return ()
    else do
      putStrLn $ reverseWords line
      main -- this loop looks strange to me...
