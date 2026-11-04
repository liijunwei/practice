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
        execute (fn [_call] "exit code: 0\n42")
        result (react-loop respond execute [{:type "message" :role "user" :content "go"}])]
    (is (= "done: 42" (:answer result)))
    (is (= 1 (:rounds result)))
    (is (= 4 (count (:history result))))))

(deftest react-loop-appends-assistant-message
  (let [respond (fn [_] [{:type "message" :content [{:type "output_text" :text "blue"}]}])
        result (react-loop respond (fn [_] "")
                           [{:type "message" :role "user" :content "fav color?"}])]
    (is (= "blue" (:answer result)))
    (is (= [{:type "message" :role "user" :content "fav color?"}
            {:type "message" :role "assistant" :content [{:type "output_text" :text "blue"}]}]
           (:history result)))))

(deftest usage-mapping
  (is (= {:input 7 :output 11 :total 18}
         (#'mini-agent1.core/api-usage {:input_tokens 7 :output_tokens 11 :total_tokens 18})))
  (is (nil? (#'mini-agent1.core/api-usage nil))))
