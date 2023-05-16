lucky :: (Integral a) => a -> String
lucky 7 = "LUCKY NUMBER SEVEN!"
lucky x = "Sorry, you're out of luck, pal!"

sayMe :: (Integral a) => a -> String
sayMe 1 = "one"
sayMe 2 = "two"
sayMe 3 = "three"
sayMe 4 = "four"
sayMe 5 = "five"
sayMe x = "not between 1 and 5" -- "Pattern match is redundant" if this line is at top

factorial :: (Integral a) => a -> a
factorial 0 = 1
factorial n = n * factorial (n - 1)

addVectors1 :: (Num a) => (a,a) -> (a,a) -> (a,a)
addVectors1 a b = (fst a + fst b, snd a + snd b)

addVectors2 :: (Num a) => (a,a) -> (a,a) -> (a,a)
addVectors2 (x1, y1) (x2, y2) = (x1+x2, y1+y2)

-- bmiTell :: (RealFloat a) => a -> String
-- bmiTell bmi
--     | bmi <= 18.5 = "You're underweight, you emo, you!"
--     | bmi <= 25.0 = "You're supposedly normal. Pffft, I bet you're ugly!"
--     | bmi <= 30.0 = "You're fat! Lose some weight, fatty!"
--     | otherwise   = "You're a whale, congratulations!"

-- bmiTell :: (RealFloat a) => a -> a -> String
-- bmiTell weight height
--     | weight / height ^ 2 <= 18.5 = "You're underweight, you emo, you!"
--     | weight / height ^ 2 <= 25.0 = "You're supposedly normal. Pffft, I bet you're ugly!"
--     | weight / height ^ 2 <= 30.0 = "You're fat! Lose some weight, fatty!"
--     | otherwise                 = "You're a whale, congratulations!"

bmiTell :: (RealFloat a) => a -> a -> String
bmiTell weight height
    | bmi <= skinny = "You're underweight, you emo, you!"
    | bmi <= normal = "You're supposedly normal. Pffft, I bet you're ugly!"
    | bmi <= fat    = "You're fat! Lose some weight, fatty!"
    | otherwise     = "You're a whale, congratulations!"
    where bmi = weight / height ^ 2
          skinny = 18.5
          normal = 25.0
          fat = 30.0


first1 :: (a,b,c) -> a
first1 (x,_,_) = x

first2 :: (Num a) => (a,b,c) -> a
first2 (x,_,_) = x


length1 :: (Num b) => [a] -> b
length1 [] = 0
length1 (_:xs) = 1 + length1 xs

sum1 :: (Num a) => [a] -> a
sum1 [] = 0
sum1 (x:xs) = x + sum1 xs

-- :t sum
-- sum :: (Foldable t, Num a) => t a -> a


capital :: String -> String
capital "" = error "Empty string, whoops!"
capital all1@(x:xs) = "The first letter of " ++ all1++ " is " ++ [x]

initials :: String -> String -> String
initials firstname lastname = [f] ++ ". " ++ [l] ++ "."
    where (f:_) = firstname
          (l:_) = lastname

cylinder :: (RealFloat a) => a -> a -> a
cylinder r h =
    let sidearea = 2 * pi * r
        toparea = pi * r ^ 2
    in sidearea + 2 * toparea



