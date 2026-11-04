(ns mini-agent1.core-test
  (:require [clojure.test :refer :all]
            [mini-agent1.core :refer :all]))

(deftest answer-text-ignores-reasoning
  (is (= "hello" (answer-text [{:type "reasoning" :content [{:type "reasoning_text" :text "secret"}]}
                               {:type "message" :content [{:type "output_text" :text "hello"}]}]))))

(deftest function-calls-picks-only-calls
  (let [call {:type "function_call" :name "bash" :arguments "{}" :call_id "c1"}]
    (is (= [call] (function-calls [{:type "reasoning"} call {:type "message"}])))))

(deftest run-bash-captures-output
  (is (= {:output "hi" :exit-code 0} (run-bash "printf hi"))))

(deftest react-loop-runs-a-tool-then-answers
  (let [respond (fn [history]
                  (if (= 1 (count history))
                    [{:type "function_call" :name "bash" :arguments "{\"command\":\"echo 42\"}" :call_id "c1"}]
                    [{:type "message" :content [{:type "output_text" :text "done: 42"}]}]))
        execute (fn [_call] "exit code: 0\n42")]
    (is (= {:answer "done: 42" :rounds 1} (react-loop respond execute "go")))))
