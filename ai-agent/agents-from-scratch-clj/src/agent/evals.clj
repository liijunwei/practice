(ns agent.evals
  (:require [agent.agent :as agent]
            [agent.memory :as memory]
            [clojure.string :as string]))

(defn eval-result
  [passed input & [expected actual error]]
  (cond-> {:passed passed :input input}
    (some? expected) (assoc :expected expected)
    (some? actual) (assoc :actual actual)
    (some? error) (assoc :error error)))

(defn suite-result [name results]
  (let [passed (count (filter :passed results))
        failed (- (count results) passed)
        total (count results)]
    {:name name
     :passed passed
     :failed failed
     :total total
     :pass-rate (if (pos? total) (/ passed total) 0.0)
     :results results}))

(defn- suite-summary [suite]
  (let [status (if (zero? (:failed suite)) "✓ PASSED" "✗ FAILED")]
    (str (:name suite) ": " status " (" (:passed suite) "/" (:total suite) ")")))

(defn test-structured-output [agent-instance cases]
  (let [results
        (mapv
         (fn [case]
           (let [input (:input case)
                 schema (:schema case)
                 required-fields (get case :must-have-fields [])]
             (try
               (let [result (agent/generate-structured agent-instance input schema)]
                 (cond
                   (nil? result)
                   (eval-result false input nil nil "Failed to parse JSON after retries")

                   (seq (remove #(contains? result %) required-fields))
                   (let [missing (remove #(contains? result %) required-fields)]
                     (eval-result false input required-fields missing "Schema contract violated"))

                   :else
                   (eval-result true input nil result)))
               (catch Exception e
                 (eval-result false input nil nil (.getMessage e))))))
         cases)]
    (suite-result "Structured Output" results)))

(defn test-tool-calls [agent-instance cases]
  (let [results
        (mapv
         (fn [case]
           (let [input (:input case)
                 expected-tool (:expected-tool case)
                 expected-args (:expected-args case)]
             (try
               (let [tool-call (agent/request-tool agent-instance input)]
                 (cond
                   (nil? tool-call)
                   (eval-result false input expected-tool nil "No tool call generated")

                   (not= expected-tool (get tool-call "tool"))
                   (eval-result false input expected-tool (get tool-call "tool") "Wrong tool selected")

                   (and expected-args
                        (some (fn [[k v]]
                                (not= v (get-in tool-call ["arguments" k])))
                              expected-args))
                   (eval-result false input expected-args (get tool-call "arguments") "Wrong argument")

                   :else
                   (eval-result true input expected-tool tool-call)))
               (catch Exception e
                 (eval-result false input expected-tool nil (.getMessage e))))))
         cases)]
    (suite-result "Tool Calls" results)))

(defn test-decisions [agent-instance cases]
  (let [results
        (mapv
         (fn [case]
           (let [input (:input case)
                 choices (:choices case)
                 expected (:expected case)]
             (try
               (let [decision (agent/decide agent-instance input choices)]
                 (cond
                   (nil? decision)
                   (eval-result false input expected nil "No decision made")

                   (not= expected decision)
                   (eval-result false input expected decision "Wrong decision")

                   :else
                   (eval-result true input expected decision)))
               (catch Exception e
                 (eval-result false input expected nil (.getMessage e))))))
         cases)]
    (suite-result "Decisions" results)))

(defn test-memory-cycle [agent-instance cases]
  (let [results
        (mapv
         (fn [case]
           (let [store-input (:store-input case)
                 query-input (:query-input case)
                 expected-substring (get case :expected-in-response "")]
             (try
               (memory/clear! (:memory agent-instance))
               (let [store-response (agent/run-with-memory agent-instance store-input)
                     query-response (agent/run-with-memory agent-instance query-input)]
                 (cond
                   (nil? store-response)
                   (eval-result false store-input nil nil "Failed to store to memory")

                   (nil? query-response)
                   (eval-result false query-input nil nil "Failed to query memory")

                   (string/includes? (string/lower-case (get query-response "reply" ""))
                                     (string/lower-case expected-substring))
                   (eval-result true (str store-input " → " query-input) expected-substring (get query-response "reply"))

                   :else
                   (eval-result false (str store-input " → " query-input)
                                expected-substring (get query-response "reply")
                                "Expected content not in response")))
               (catch Exception e
                 (eval-result false store-input nil nil (.getMessage e))))))
         cases)]
    (suite-result "Memory Cycle" results)))

(defn run-all
  [agent-instance & {:keys [structured-cases tool-cases decision-cases memory-cases]}]
  (cond-> []
    (seq structured-cases) (conj (test-structured-output agent-instance structured-cases))
    (seq tool-cases) (conj (test-tool-calls agent-instance tool-cases))
    (seq decision-cases) (conj (test-decisions agent-instance decision-cases))
    (seq memory-cases) (conj (test-memory-cycle agent-instance memory-cases))))

(defn print-eval-report [suites]
  (println "\n" (apply str (repeat 50 "=")))
  (println "EVAL REPORT")
  (println (apply str (repeat 50 "=")))
  (let [total-passed (reduce + (map :passed suites))
        total-failed (reduce + (map :failed suites))]
    (doseq [suite suites]
      (println "\n" (suite-summary suite))
      (doseq [result (remove :passed (:results suite))]
        (println "  ✗ Input:" (subs (str (:input result)) 0 (min 50 (count (str (:input result))))) "...")
        (when-let [expected (:expected result)] (println "    Expected:" expected))
        (when-let [actual (:actual result)] (println "    Actual:" actual))
        (when-let [error (:error result)] (println "    Error:" error))))
    (println "\n" (apply str (repeat 50 "-")))
    (println "Overall:"
             (if (zero? total-failed)
               "✓ ALL PASSED"
               (str "✗ " total-failed " FAILED"))
             (str "(" total-passed "/" (+ total-passed total-failed) ")"))
    (println (apply str (repeat 50 "=")))))
