(ns shared.utils
  (:require [clojure.data.json :as json]
            [clojure.string :as string]))

(defn safe-json-parse
  [text]
  (try
    (json/read-str text)
    (catch Exception _ nil)))

(defn- strip-code-fence
  [text]
  (let [text (string/trim text)]
    (cond
      (string/starts-with? text "```json") (string/trim (subs text 7))
      (string/starts-with? text "```") (string/trim (subs text 3))
      :else text)))

(defn- strip-trailing-fence
  [text]
  (if (string/ends-with? text "```")
    (string/trim (subs text 0 (- (count text) 3)))
    text))

(defn- strip-common-prefixes
  [text]
  (reduce (fn [s prefix]
            (if (string/starts-with? s prefix)
              (string/trim (subs s (count prefix)))
              s))
          text
          ["JSON:" "Response:" "Answer:" "Here's the JSON:" "The JSON is:"]))

(defn- char-count
  [s ch]
  (count (filter #(= ch %) s)))

(defn- try-fix-unclosed-string
  [json-text]
  (let [last-quote (string/last-index-of json-text "\"")]
    (when last-quote
      (let [before (subs json-text 0 last-quote)]
        (when (even? (char-count before \"))
          (safe-json-parse
           (str (subs json-text 0 (inc last-quote))
                "\""
                (subs json-text (inc last-quote))
                "}")))))))

(defn- parse-between
  [text open close]
  (let [start (string/index-of text open)
        end (string/last-index-of text close)]
    (when (and start end (> end start))
      (let [json-text (subs text start (inc end))]
        (or (safe-json-parse json-text)
            ;; The model sometimes leaves the final string unclosed.
            (when (odd? (char-count json-text \"))
              (try-fix-unclosed-string json-text)))))))

(defn extract-json-from-text
  [text]
  (when (seq text)
    (let [text (-> text strip-code-fence strip-trailing-fence strip-common-prefixes string/trim)]
      (or (safe-json-parse text)
          (parse-between text "{" "}")
          (parse-between text "[" "]")
          (some (fn [line]
                  (let [line (string/trim line)]
                    (when (or (string/starts-with? line "{")
                              (string/starts-with? line "["))
                      (safe-json-parse line))))
                (string/split-lines text))))))

(defn format-messages
  [messages]
  (string/join
   "\n"
   (map (fn [msg]
          (str "[" (string/upper-case (get msg "role" "unknown")) "]\n"
               (get msg "content" "") "\n"))
        messages)))
