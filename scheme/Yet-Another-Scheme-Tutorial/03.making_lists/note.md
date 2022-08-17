+ cons cells
+ cons function
+ car part: Contents of the Address part of the Register
+ cdr part: Contents of the Decrement part of the Register

+ These names also indicate that the reality of the cons cell is a memory space
+ The name cons is an abbreviation of a English term 'construction' for your information

```
(cons 1 2)
;Value 11: (1 . 2)
```

![](./cons2.png)

```
(cons 3 (cons 1 2))
;Value 15: (3 1 . 2)
```
