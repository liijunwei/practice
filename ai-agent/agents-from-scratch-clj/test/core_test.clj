(ns core-test
  (:require [clojure.test :refer [deftest is]]
            [shared.utils :as utils]
            [agent.state :as state]
            [agent.memory :as memory]
            [agent.tools :as tools]
            [agent.planner :as planner]))

(deftest json-extraction
  (is (= {"a" 1} (utils/extract-json-from-text "prefix {\"a\": 1} suffix")))
  (is (= {"a" 1} (utils/extract-json-from-text "```json\n{\"a\":1}\n```")))
  (is (nil? (utils/safe-json-parse "not json"))))

(deftest state-and-memory
  (let [s (state/new-state)
        m (memory/new-memory)]
    (state/increment-step! s)
    (state/mark-done! s)
    (is (= {:steps 1 :done true :current-plan nil :last-action nil} (state/state-map s)))
    (memory/add! m "alpha")
    (memory/add! m "alpha")
    (memory/add! m "beta")
    (is (= ["alpha" "beta"] (memory/get-all m)))
    (is (= ["beta"] (memory/search m "bet")))))

(deftest calculator-tool
  (is (= 6.0 (tools/calculator 2.0 4.0 "add")))
  (is (= 12.0 (tools/execute-tool "calculator" {"a" 3.0 "b" 4.0 "operation" "multiply"}))))

(deftest graph-execution
  (let [graph {"nodes" [{"id" "2" "action" "write" "depends_on" ["1"]}
                         {"id" "1" "action" "research" "depends_on" []}]}
        results (planner/execute-graph graph #(str "executed-" %))]
    (is (= #{"1" "2"} (set (map #(get % "node_id") results))))
    (is (= 2 (count (filter #(get % "success") results))))))

(defn -main [& _]
  (let [{:keys [fail error]} (clojure.test/run-tests 'core-test)]
    (System/exit (if (zero? (+ fail error)) 0 1))))
