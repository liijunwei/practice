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
