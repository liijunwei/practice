(ns fwpd.core
  (:require [clojure.string :as str]))

(def filename "suspects.csv")

;(print (slurp filename))

(mapv #(clojure.string/split % #"\n") (slurp filename))

(mapv #(clojure.string/split % #",") (clojure.string/split (slurp filename) #"\n"))

(defn row->map
  [[name glitter-index]]
  {:name name
   :glitter-index (Integer/parseInt glitter-index)})

(defn row->map-1
  [row]
  (let [name (nth row 0 nil)
        glitter-index (nth row 1 nil)]
    {:name name :glitter-index glitter-index}))

(map row->map (mapv #(clojure.string/split % #",") (clojure.string/split (slurp filename) #"\n")))

(->> (slurp filename)
     str/split-lines
     (map #(str/split % #","))
     (map row->map-1))

