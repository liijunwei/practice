```bash
$ mit-scheme

1 ]=> (+ 1 2)

;Value: 3

4 error> (+)

;Value: 0

4 error> (+ 1)

;Value: 1

1 ]=> (+ 1 2 3)

;Value: 6


```


# 3. Four Basic Arithmetic Operations

```
+

-

*

/
```

```
1 ]=> (* (+ 2 3) (- 5 3))

;Value: 10

1 ]=> (exact->inexact (/ 29 3 7))

;Value: 1.380952380952381
```

# 4. Other Arithmetic Operations


```
(quotient 7 3)
(modulo 7 3)
(sqrt 8)
```

## 4.2 Trigonometric Functions

sin, cos, tan, asin, acos, and atan

```
(atan 1)
(atan 1 0)
(sin 3.14)
```

## 4.3. Exponential and Logarithm

Exponential and logarithm are calculated by exp, and log, respectively. The value of b to the power of a can be calculated by (expt a b).

```

```
