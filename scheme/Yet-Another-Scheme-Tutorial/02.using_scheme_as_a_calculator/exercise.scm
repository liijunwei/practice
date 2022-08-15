; Exercise 1

; Calculate followings using Scheme interpreter.

; (1+39) * (53-45)
(* (+ 1 39) (- 53 45))

; (1020 / 39)  (45 * 2)
(+ (/ 1020 39) (* 45 2))

; Sum of 39, 48, 72, 23, and 91
(+ 39 48 72 23 91)

; Average of 39, 48, 72, 23, and 91 as a floating point.
(exact->inexact (/ (+ 39 48 72 23 91) 5))
; 54.6


; Exercise 2

; circle ratio, π
; sin(pi/2) = 1
(* 2 (asin 1))

; exp(2/3)
(exp (/ 2 3))
(exp (exact->inexact (/ 2 3)))
(exp 2/3)

; 3 to the power of 4
; The value of b to the power of a can be calculated by (expt a b)
(expt 3 4)

; logarithm of 1000
(log 1000)

