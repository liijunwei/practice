module Geometry
( sphereVolume
, sphereArea
, cubeVolume
, cubeArea
, cuboidArea
, cuboidVolume
) where

-- 球的体积
sphereVolume :: Float -> Float
sphereVolume radius = (4.0 / 3.0) * pi * (radius ^ 3)

-- 球的表面积
sphereArea :: Float -> Float
sphereArea radius = 4 * pi * (radius ^2)

-- 立(正)方体的体积
cubeVolume :: Float -> Float
cubeVolume side = cuboidVolume side side side

-- 立(正)方体的表面积
cubeArea :: Float -> Float
cubeArea side = cuboidArea side side side

-- 长方体的表面积
cuboidArea :: Float -> Float -> Float -> Float
-- cuboidArea a b c = 2 * (rectangleArea a b + rectangleArea a c + rectangleArea b c)
cuboidArea a b c = 2 * sum [rectangleArea a b, rectangleArea a c, rectangleArea b c]

-- 长方体的体积
cuboidVolume :: Float -> Float -> Float -> Float
cuboidVolume a b c = rectangleArea a b * c

-- 矩形的表面积
rectangleArea :: Float -> Float -> Float
rectangleArea a b = a * b
