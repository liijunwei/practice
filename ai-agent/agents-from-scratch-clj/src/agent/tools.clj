(ns agent.tools)

(defn calculator
  ([a b] (calculator a b "add"))
  ([a b operation]
   (case operation
     "add" (+ a b)
     "subtract" (- a b)
     "multiply" (* a b)
     "divide" (if (zero? b) Double/POSITIVE_INFINITY (/ a b))
     (throw (ex-info (str "Unknown operation: " operation) {:operation operation})))))

(defn get-tool-schema []
  {"calculator"
   {"description" "Perform basic arithmetic operations"
    "parameters"
    {"a" {"type" "number" "description" "First number"}
     "b" {"type" "number" "description" "Second number"}
     "operation" {"type" "string"
                  "enum" ["add" "subtract" "multiply" "divide"]
                  "description" "The operation to perform"}}
    "required" ["a" "b"]}})

(def tool-fns
  {"calculator" calculator})

(defn execute-tool
  [tool-name arguments]
  (if-let [tool-fn (get tool-fns tool-name)]
    (tool-fn (get arguments "a")
             (get arguments "b")
             (get arguments "operation" "add"))
    (throw (ex-info (str "Unknown tool: " tool-name) {:tool tool-name}))))
