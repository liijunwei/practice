module Shapes
( Point(..)
, Shape(..)
, surface
, nudge
) where

import qualified Data.Map as Map

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
-- If our type acts as some kind of box, it's good to use them

-- a,b,c are type parameters
data Car1 a b c = Car1 {
  company1 :: a,
  model1 :: b,
  year1 :: c
} deriving (Show)


-- data (Ord k) => Map k v = ...
-- However, it's a very strong convention in Haskell to never add typeclass constraints in data declarations.
-- Why? Well, because we don't benefit a lot, but we end up writing more class constraints, even when we don't need them.

data Vector a = Vector a a a deriving (Show)
vplus :: (Num t) => Vector t -> Vector t -> Vector t
(Vector i j k) `vplus` (Vector l m n) = Vector (i+l) (j+m) (k+n)

-- Vector 1 2 3 `vplus` Vector 1 1 1
-- vplus (Vector 1 2 3) (Vector 1 1 1)

vectMult :: (Num t) => Vector t -> t -> Vector t
(Vector i j k) `vectMult` m = Vector (i*m) (j*m) (k*m)

-- vectMult (Vector 1 2 3) 10

scalarMult :: (Num t) => Vector t -> Vector t -> t
(Vector i j k) `scalarMult` (Vector l m n) = i*l + j*m + k*n

-- scalarMult (Vector 1 2 3) (Vector 10 10 10)

-- These functions can operate on types of Vector Int, Vector Integer, Vector Float...
-- whatever, as long as the a from Vector a is from the Num typeclass.

-- http://learnyouahaskell.com/making-our-own-types-and-typeclasses
-- Once again, it's very **important** to distinguish between the type constructor and the value constructor.
-- When declaring a data type, the part before the = is the type constructor and the constructors after it (possibly separated by |'s) are value constructors.

-- type constructor vs type constraint vs value constructor
    -- 用 value constructor 声明值(value), 接收参数得到不同的值
    -- 用 type  constructor 声明类型(type)，接收参数得到不同限制的类型(type constraint)

-- Typeclasses are more like interfaces.
-- We don't make data from typeclasses.
-- Instead, we first make our data type and then we think about what it can act like.

-- 感觉和duck-typing是一个意思

-- 接下来的章节，我们会学习怎么通过实现类型的方法 手动的把我们的类型变为某个类型的实例
-- In the next section, we'll take a look at how we can manually make our types instances of typeclasses by implementing the functions defined by the typeclasses.

-- 但是在这之前，我们先来了解一下怎么通过haskell的deriving来自动让某个实例变为某个类型的实力

data Person1 = Person1 {
  firstName1 :: String,
  lastName1 :: String,
  age1 :: Int
} deriving (Eq, Show, Read)

-- let guy1 = Person1 {firstName1 = "Buddy", lastName1 = "Finklestein", age1 = 99}
-- let guy2 = Person1 {firstName1 = "Buddy", lastName1 = "Finklestein", age1 = 99}
-- guy1 == guy2
-- "hi " ++ show guy1
-- read "Person1 {firstName1 = \"Buddy\", lastName1 = \"Finklestein\", age1 = 99}" :: Person1

data Day = Monday | Tuesday | Wednesday | Thursday | Friday | Saturday | Sunday
  deriving (Eq, Ord, Show, Read, Bounded, Enum)

-- minBound :: Day
-- maxBound :: Day
-- succ Monday
-- pred Tuesday
-- [Monday .. Sunday]
-- [minBound .. maxBound] :: [Day]

-- 没做什么特殊的事，只是提升了可读性而已
-- Type synonyms

-- type String = [Char]

type PPP = Person1
info :: PPP -> String
info p = show p

type Name = String
type PhoneNumber = String
type PhoneBook = [(Name, PhoneNumber)]

phoneBook :: [(Name, PhoneNumber)]
phoneBook = [("betty","555-2938"),("bonnie","452-2928"),("patsy","493-2928"),("lucille","205-2928"),("wendy","939-8282"),("penny","853-2492")]

inPhoneBook :: Name -> PhoneNumber -> PhoneBook -> Bool
inPhoneBook name pnumber pbook = (name, pnumber) `elem` pbook

-- inPhoneBook "betty" "555-2938" phoneBook
-- inPhoneBook "betty" "555-29381" phoneBook

-- type synonyms 可以用来提升代码的可读性，但是要注意不要滥用
-- the type declaration that took advantage of type synonyms is easier to understand.
-- However, you shouldn't go overboard with them.

-- concrete types
-- **values can only have types that are concrete types**

-- Just like we can partially apply functions to get new functions, we can partially apply type parameters and get new type constructors from them.


-- The Maybe a type is defined in Haskell's standard library as follows: (https://wiki.haskell.org/Maybe)
-- data Maybe a = Nothing | Just a
-- Mayby is a type constructor
-- Mayby a is a type

-- 终于基本看懂了 type constructor
-- 又出现了个 type synonyms... x.x
-- 它的含义很简单，只是为什么看它的用法总觉得很绕呢？

-- 函数可以柯里化，类型别名也可以柯里化
-- type IntMap = Map.Map Int

data IntMap a = KKK a deriving (Show)
-- intMapDemo = IntMap 1 String

-- Maybe a 可以用来表示 出错，或者有值，但是仅仅是“出错”这个信息对于有些场景来说太过于模糊了
-- 我们可以用Either这个type constructor来明确说明可能出问题的情况
-- when we're interested in how some function failed or why, we usually use the result type of Either a b, where a is some sort of type that can tell us something about the possible failure and b is the type of a successful computation
-- Hence, errors use the Left value constructor while results use Right.

data LockerState = Taken | Free deriving (Show, Eq)
type Code = String
type ErrorMessage = String
type LockerNumber = Int
type LockerMap = Map.Map LockerNumber (LockerState, Code)

lockerLookup :: LockerNumber -> LockerMap -> Either ErrorMessage Code
lockerLookup lockernumber map =
  case Map.lookup lockernumber map of
    Nothing -> Left $ "Locker number " ++ show lockernumber ++ " doesn't exist!"
    Just (state, code) -> if state /= Taken then Right code
                          else Left $ "Locker " ++ show lockernumber ++ " is already taken!"

lockers :: LockerMap
lockers = Map.fromList
 [(100,(Taken,"ZD39I")),
  (101,(Free,"JAH3I")),
  (103,(Free,"IQSA9")),
  (105,(Free,"QOTSA")),
  (109,(Taken,"893JJ")),
  (110,(Taken,"99292"))]

-- lockerLookup 100 lockers
-- lockerLookup 101 lockers
-- lockerLookup 102 lockers
-- lockerLookup 103 lockers
-- lockerLookup 105 lockers
-- lockerLookup 109 lockers
-- lockerLookup 110 lockers
