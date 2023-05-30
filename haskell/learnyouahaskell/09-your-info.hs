-- ghc --make 09-your-info.hs -o a.out
-- echo -e "Jake\nBright" | ./a.out

import Data.Char

-- use `<-` when you want to bind results of I/O actions to names
-- use `let` bindings to bind pure expressions to names.
main = do
  putStrLn "what's your first name?"
  firstName <- getLine
  putStrLn "what's your last name?"
  lastName <- getLine
  let bigFirstName = map toUpper firstName
      bigLastName = map toUpper lastName
  putStrLn $ "Hey " ++ bigFirstName ++ " " ++ bigLastName ++ ", how are you?"
