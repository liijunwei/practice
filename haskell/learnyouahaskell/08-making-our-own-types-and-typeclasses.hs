-- how do we make our own? Well, one way is to use the data keyword to define a type
data Bool = False | True

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



