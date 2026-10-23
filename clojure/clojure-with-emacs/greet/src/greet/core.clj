(ns greet.core)

(defn greet [s]
  (str "greeting from " s "!"))

; paredit
(defn foo [])
(defn bar [])

; C-c M-j CIDER jackin

(comment
  (greet "moss"))

; C-u C-c C-c to debug
