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

(->> (slurp filename)
     str/split-lines
     (map #(str/split % #","))
     (map row->map))

(def vamp-keys [:name :glitter-index])
(defn str->int
  [str]
  (Integer. str))
(def conversions {:name identity
                  :glitter-index str->int})

(defn convert
  [vamp-key value]
  ((get conversions vamp-key) value))


(defn mapify [rows]
  (map (fn [unmapped-row]
         (let [[k1 k2] vamp-keys
               [v1 v2] unmapped-row]
           (-> {}
               (assoc k1 (convert k1 v1))
               (assoc k2 (convert k2 v2)))))
       rows))

(defn parse
  "Convert a CSV into rows of columns"
  [string]
  (map #(clojure.string/split % #",")
       (clojure.string/split string #"\n")))

(defn mapify
  "Return a seq of maps like {:name \"Edward Cullen\" :glitter-index 10}"
  [rows]
  (map (fn [unmapped-row]
         (reduce (fn [row-map [vamp-key value]]
                   (assoc row-map vamp-key (convert vamp-key value)))
                 {}
                 (map vector vamp-keys unmapped-row)))
       rows))

(mapify (parse (slurp filename)))