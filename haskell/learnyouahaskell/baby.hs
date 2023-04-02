doubleMe x = x + x
tripleMe x = x * 3
-- doubleUs x y = x*2 + y*2
doubleUs x y = doubleMe x + doubleMe y

-- doubleSmallNumber x = if x > 100
--                         then x
--                         else x*2

doubleSmallNumber' x = (if x > 100 then x else x*2) + 1
