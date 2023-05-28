module Shapes
( Point(..)
, Shape(..)
, surface
, nudge
) where

import qualified Data.Map as Map
import Data.Time

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

data List a = Empty | MyCon { listHead :: a, listTail :: List a} deriving (Show, Read, Eq, Ord)

-- Empty
-- 3 `Cons` (4 `Cons` (5 `Cons` Empty))

data Tree a = EmptyTree | Node a (Tree a) (Tree a) deriving (Show, Read, Eq)

singleton :: a -> Tree a
singleton x = Node x EmptyTree EmptyTree

treeInsert :: (Ord a) => a -> Tree a -> Tree a
treeInsert x EmptyTree = singleton x
treeInsert x (Node a left right)
  | x == a = Node x left right
  | x < a = Node a (treeInsert x left) right
  | x > a = Node a left (treeInsert x right)

treeElem :: (Ord a) => a -> Tree a -> Bool
treeElem x EmptyTree = False
treeElem x (Node a left right)
  | x == a = True
  | x < a = treeElem x left
  | x > a = treeElem x right


-- let nums = [8,6,4,1,7,3,5]
-- let numsTree = foldr treeInsert EmptyTree nums
-- numsTree
-- 8 `treeElem` numsTree
-- 888 `treeElem` numsTree


data User = User {
  user_id :: Int,
  name :: String
} deriving (Show, Eq)

data Status = Pending | Processing | Failed | Succeed deriving (Show, Eq, Enum)

data Transaction = Transaction {
  tx_id :: Int,
  from_user_id :: Int,
  to_user_id :: Int,
  amount :: Float,
  status :: Status,
  created_at :: UTCTime,
  updated_at :: UTCTime
} deriving (Show, Eq)

user1 :: User
user1 = User {user_id = 1, name = "Ben"}

user2 :: User
user2 = User {user_id = 2, name = "Rob"}

pending_tx :: Transaction
pending_tx = Transaction {tx_id = 1, from_user_id = (user_id user1), to_user_id = (user_id user2), amount = 1, status = Pending, created_at = (read "2023-05-26 00:00:00 UTC" :: UTCTime), updated_at = (read "2023-05-26 00:00:00 UTC" :: UTCTime)}

-- Q: is this the right way to update tx status?
-- probably not
markProcessing1 :: Transaction -> Transaction
markProcessing1 tx = Transaction {
  tx_id = tx_id tx,
  from_user_id = from_user_id tx,
  to_user_id = to_user_id tx,
  amount = amount tx,
  status = Processing,
  created_at = created_at tx,
  updated_at = updated_at tx
}

processing_tx1 :: Transaction
processing_tx1 = markProcessing1 pending_tx

markProcessing2 :: Transaction -> Status -> Transaction
markProcessing2 tx new_status = tx {status = new_status}

processing_tx2 :: Transaction
processing_tx2 = markProcessing2 pending_tx Processing

r1 :: Bool
r1 = processing_tx2 == pending_tx

r2 :: Bool
r2 = processing_tx1 == processing_tx2


-- data Mood = Happy | Sad | Bored

-- data means that we're defining a new data type

-- class Eq a where
--     (==) :: a -> a -> Bool
--     (/=) :: a -> a -> Bool
--     x == y = not (x /= y)
--     x /= y = not (x == y)

-- class Eq a where, this means that we're defining a new typeclass and that's called Eq

-- data type VS typeclass
-- different things

-- mutual recursion

data TrafficLight = Red | Yellow | Green
-- Notice how we didn't derive any class instances for it.
-- That's because we're going to write up some instances by hand,
-- even though we could derive them for types like Eq and Show
-- Here's how we make it an instance of Eq

instance Eq TrafficLight where
  Red == Red = True
  Green == Green = True
  Yellow == Yellow = True
  _ == _ = False


-- We did it by using the instance keyword
-- So class is for defining new typeclasses
--    instance is for making our types instances of typeclasses.
-- 使用class关键字定义新的typeclass
-- 使用instance关键字把我们自定义的类型变为typeclass的实例

-- 前面用到 "mutual recursion" 的作用可能是为了使之后的定义变得简单
-- 比如 将 TrafficLight 变为Eq的实例时，我们只需要定义 == 方法，可以忽略 /= 方法，因为这是一个 "minimal complete definition for the typeclass"
-- 如果不是使用"mutual recursion"， 那么Eq的定义会是下面这样
    -- here is Eq typeclass definition
    -- class Eq a where
    --     (==) :: a -> a -> Bool
    --     (/=) :: a -> a -> Bool
-- 那么将 TrafficLight 变为Eq的实例时，既要定义 ==, 又要定义/=

instance Show TrafficLight where
  show Red = "Red Light"
  show Green = "Green Light"
  show Yellow = "Yellow Light"

-- ghci> Red
-- Red Light
-- ghci> Green
-- Green Light
-- ghci> Yellow
-- Yellow Light
-- ghci> Red == Red
-- True
-- ghci> Red == Green
-- False
-- ghci> Red `elem` [Red, Yellow, Green]
-- True

-- comon sense: all the types in functions have to be concrete

-- This is like saying that we want to make all types of the form Maybe something an instance of Eq.
--
-- instance Eq (Maybe m) where
--     Just x == Just y = x == y
--     Nothing == Nothing = True
--     _ == _ = False

-- with issue fixed, (Maybe m) has to be comparable
-- instance (Eq m) => Eq (Maybe m) where
--     Just x == Just y = x == y
--     Nothing == Nothing = True
--     _ == _ = False


-- see what the instances of a typeclass are, just do :info YourTypeClass in GHCI
--
-- :info Eq
-- :info Num
-- :info Show
-- :info Maybe
-- :info Bool


class YesNo a where
  yesno :: a -> Bool

instance YesNo Int where
  yesno 0 = False
  yesno _ = True

instance YesNo [a] where
  yesno [] = False
  yesno _ = True

instance YesNo Bool where
  yesno = id
  -- id is a stdlib function that takes a parameter and returns the same thing

instance YesNo (Maybe a) where
  yesno (Just _) = True
  yesno Nothing = False

instance YesNo (Tree a) where
  yesno EmptyTree = False
  yesno _ = True

-- lol
instance YesNo TrafficLight where
  yesno Red = False
  yesno _ = True

y1 :: Bool
y1 = yesno []

y2 :: Bool
y2 = yesno "haha"

y3 :: Bool
y3 = yesno ""

y4 :: Bool
y4 = yesno $ Just 0

y5 :: Bool
y5 = yesno True

y6 :: Bool
y6 = yesno EmptyTree

y7 :: Bool
y7 = yesno [1,2,3]

y8 :: Bool
y8 = yesno True

-- :t yesno

-- mimic the if statement
yesnoIf :: (YesNo y) => y -> a -> a -> a
yesnoIf yesnoVal yesResult noResult = if yesno yesnoVal then yesResult else noResult

-- yesnoIf [] "YEAH!" "NO!"         -- "NO!"
-- yesnoIf [2,3,4] "YEAH!" "NO!"    -- "YEAH!"
-- yesnoIf True "YEAH!" "NO!"       -- "YEAH!"
-- yesnoIf (Just 500) "YEAH!" "NO!" -- "YEAH!"
-- yesnoIf Nothing "YEAH!" "NO!"    -- "NO!"


-- The Functor typeclass
    -- which is basically for things that can be mapped over.
    -- list type is part of the Functor typeclass.

-- :info Functor

-- class Functor f where
--     fmap :: (a -> b) -> f a -> f b

-- :t fmap
-- fmap :: Functor f => (a -> b) -> f a -> f b

-- :t map
-- map :: (a -> b) -> [a] -> [b]

-- in fact, map is just a fmap that works only on lists


-- instance Functor [] where
--     fmap = map

-- [a] is already a concrete type (of a list with any type inside it)
-- while [] is a type constructor that takes one type and can produce types such as [Int], [String] or even [[String]].

-- **Types that can act like a box can be functors.**

-- **Functor wants a type constructor that takes one type and not a concrete type.**
--
-- Functor 的实例，在实现方法的时候，指定的是 type constructor, not concrete type
    -- instance Functor [] where...    (not [a])
    -- instance Functor Maybe where... (not Maybe m)
--
-- 以 :t fmap 为例来看
-- 如果我们把 f 替换为 Maybe，那么fmap的类型签名为(type signature)
    -- (a -> b) -> Maybe a -> Maybe b
    -- 没问题
-- 但是如果把 f 替换为 Maybe m, 那么fmap的类型签名为(type signature)
    -- (a -> b) -> Maybe m a -> Maybe m b
    -- 说不通，没有任何意义，因为 Maybe 只接收一个参数，并返回一个concrete type


-- fmap (++ " HEY GUYS IM INSIDE THE JUST") (Just "Something serious.")
-- fmap (++ " HEY GUYS IM INSIDE THE JUST") Nothing
-- fmap (*2) (Just 200)
-- fmap (*2) Nothing

-- Another thing that can be mapped over and made an instance of Functor is our Tree a type.
    -- (a -> b) -> Tree a -> Tree b

instance Functor Tree where
  fmap f EmptyTree = EmptyTree
  fmap f (Node x left right) = Node (f x) (fmap f left) (fmap f right)

-- fmap (*2) EmptyTree
-- fmap (*2) (foldr treeInsert EmptyTree [5,7,3,2,1,7,6])

-- :info Either
-- fmap :: Functor f => (a -> b) -> f a -> f b
-- (a -> b) -> Either e a -> Either e b


