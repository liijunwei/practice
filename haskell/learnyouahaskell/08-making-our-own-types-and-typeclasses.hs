module Shapes
( Point(..)
, Shape(..)
, surface
, nudge
) where

-- how do we make our own? Well, one way is to use the data keyword to define a type
-- data Bool = False | True

-- data means that we're defining a new data type
-- The part before the = denotes the type, which is Bool
-- The parts after the = are value constructors
-- They specify the different values that this type can have.
-- The | is read as or
-- Above read as: "the Bool type can have a value of True or False"

-- data Shape = Circle Float Float Float | Rectangle Float Float Float Float

-- **Value constructors** are actually functions that ultimately return a value of a data type.
-- Let's take a look at the type signatures for these two value constructors.
    -- :t Circle
    -- Circle :: Float -> Float -> Float -> Shape
    -- :t Rectangle
    -- Rectangle :: Float -> Float -> Float -> Float -> Shape

-- type declaration: the function takes a shape and returns a float
-- surface :: Shape -> Float
-- surface (Circle _ _ r) = pi * r ^ 2
-- surface (Rectangle x1 y1 x2 y2) = (abs $ x2 - x1) * (abs $ y2 - y1)

-- surface (Circle 1 1 2)
-- surface (Rectangle 0 0 1 2)

-- we wouldn't write `Circle -> Float` as **Circle is not a type, Shape is**
-- same as we wouldn't write `True -> Int` as True is not a type, Bool is

-- Circle 10 20 5
-- error because our Shape type hasn't implement show function yet

-- data Shape = Circle Float Float Float | Rectangle Float Float Float Float deriving (Show)

-- value constructors 是函数，所以可以利用柯里化
-- Value constructors are functions, so we can map them and partially apply them and everything.

-- 圆心为 10，20 ，半径未定的圆形
-- mycircle = Circle 10 20
-- surface $ mycircle 2
-- surface $ mycircle 3

-- map (Circle 10 20) [4,5,6,6]

data Point = Point Float Float deriving (Show)
data Shape = Circle Point Float | Rectangle Point Point deriving (Show)

surface :: Shape -> Float
surface (Circle _ r) = pi * r ^ 2
surface (Rectangle (Point x1 y1) (Point x2 y2)) = (abs $ x2 - x1) * (abs $ y2 - y1)

-- surface (Rectangle (Point 0 0) (Point 2 2))
-- surface (Circle (Point 0 0) 24)


-- nudge 推动/移动
nudge :: Shape -> Float -> Float -> Shape
nudge (Circle (Point x y) r) a b = Circle (Point (x+a) (y+b)) r
nudge (Rectangle (Point x1 y1) (Point x2 y2)) a b = Rectangle (Point (x1+a) (y1+b)) (Point (x2+a) (y2+b))

-- nudge (Circle (Point 0 0) 10) 2 2
-- nudge (Rectangle (Point 0 0) (Point 1 1)) 2 2

-- Record syntax
data Person = Person {
  firstName :: String,
  lastName :: String,
  age :: Int,
  height :: Float,
  phoneNumber :: String,
  flavor :: String
} deriving (Show)

-- let guy = Person "Buddy" "Finklestein" 43 184.2 "526-2928" "Chocolate"
-- let guy = Person {flavor = "Chocolate", firstName = "Buddy", lastName = "Finklestein", age = 43, height = 184.2, phoneNumber = "526-2928"}
-- firstName guy
-- flavor guy

-- we don't have to necessarily put the fields in the proper order, as long as we list all of them.
-- But if we don't use record syntax, we have to specify them in order.

data Car = Car {
  company :: String,
  model :: String,
  year :: Int
} deriving (Show)

-- Type parameters
    -- A value constructor can take some values parameters and then produce a new value.
    -- type constructors can take types as parameters to produce new types

-- type constructor
    -- Maybe


-- Just "Haha"
-- Just 10 :: Maybe Double

-- An empty list can act like a list of anything.
-- That's why we can do [1,2,3] ++ [] and ["ha","ha","ha"] ++ []

-- 感觉type parameter好难懂...
-- Using type parameters is very beneficial, but only when using them makes sense.



