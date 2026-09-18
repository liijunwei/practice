(ns evals.golden-datasets
  (:require [clojure.data.json :as json]
            [clojure.string :as string]))

(defn schema-string
  [fields example]
  (str (json/write-str fields)
       "\n\nExample: "
       (json/write-str example)))

(def structured-output-golden
  [{:input "Explain quantum computing in one sentence"
    :schema (schema-string
             {"topic" "the topic name as a string"
              "difficulty" "beginner | intermediate | advanced"}
             {"topic" "machine learning" "difficulty" "intermediate"})
    :must-have-fields ["topic" "difficulty"]}
   {:input "What is Python in one sentence?"
    :schema (schema-string
             {"topic" "the topic name as a string"
              "difficulty" "beginner | intermediate | advanced"}
             {"topic" "web development" "difficulty" "beginner"})
    :must-have-fields ["topic" "difficulty"]}
   {:input "What is the significance of 42?"
    :schema (schema-string
             {"answer" "your answer as a string"}
             {"answer" "It is the meaning of life"})
    :must-have-fields ["answer"]}
   {:input "What does hello world mean in programming?"
    :schema (schema-string
             {"explanation" "your explanation as a string"}
             {"explanation" "It is a simple test program"})
    :must-have-fields ["explanation"]}])

(def tool-call-golden
  [{:input "What is 42 * 7?" :expected-tool "calculator" :expected-args {"operation" "multiply"}}
   {:input "Calculate 100 + 50" :expected-tool "calculator" :expected-args {"operation" "add"}}
   {:input "What is 100 / 5?" :expected-tool "calculator" :expected-args {"operation" "divide"}}
   {:input "What's 50 minus 25?" :expected-tool "calculator" :expected-args {"operation" "subtract"}}
   {:input "If I have 15 apples and buy 27 more, how many do I have?"
    :expected-tool "calculator" :expected-args {"operation" "add"}}])

(def decision-golden
  [{:input "Can you summarize this article for me?"
    :choices ["answer_question" "summarize_text" "translate"]
    :expected "summarize_text"}
   {:input "Translate 'hello' to Spanish"
    :choices ["answer_question" "summarize_text" "translate"]
    :expected "translate"}
   {:input "What is the capital of France?"
    :choices ["answer_question" "summarize_text" "translate"]
    :expected "answer_question"}
   {:input "What is 5 + 5?"
    :choices ["answer_question" "calculate" "search"]
    :expected "calculate"}])

(def memory-golden
  [{:store-input "My name is Alice"
    :query-input "What's my name?"
    :expected-in-response "Alice"}
   {:store-input "I prefer dark mode"
    :query-input "What's my preference for display mode?"
    :expected-in-response "dark"}
   {:store-input "I live in New York"
    :query-input "Where do I live?"
    :expected-in-response "New York"}])

(def edge-cases-golden
  {"empty_input"
   {:structured
    {:input "Respond with a greeting"
     :schema (schema-string
              {"response" "your response"}
              {"response" "Hello!"})
     :must-have-fields ["response"]}}
   "very_long_input"
   {:structured
    {:input (str "Summarize: " (string/join " " (repeat 20 "very")) " complex topic")
     :schema (schema-string
              {"summary" "brief summary"}
              {"summary" "A complex topic"})
     :must-have-fields ["summary"]}}
   "unicode_input"
   {:structured
    {:input "What does hello mean in Chinese?"
     :schema (schema-string
              {"translation" "the translation"}
              {"translation" "你好"})
     :must-have-fields ["translation"]}}
   "json_in_input"
   {:structured
    {:input "What format is this: key value pairs?"
     :schema (schema-string
              {"parsed" "your answer"}
              {"parsed" "dictionary"})
     :must-have-fields ["parsed"]}}})
