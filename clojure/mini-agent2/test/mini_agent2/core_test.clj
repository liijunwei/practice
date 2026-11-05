(ns mini-agent2.core-test
  (:require [clojure.test :refer :all]
            [mini-agent2.core :as agent]))

(deftest answer-text-ignores-reasoning
  (is (= "hello"
         (agent/answer-text
          [{:type "reasoning"
            :content [{:type "reasoning_text" :text "secret"}]}
           {:type "message"
            :content [{:type "output_text" :text "hello"}]}]))))

(deftest function-calls-picks-only-calls
  (let [call {:type "function_call" :name "bash" :arguments "{}" :call_id "c1"}]
    (is (= [call]
           (agent/function-calls
            [{:type "reasoning"} call {:type "message"}])))))

(deftest run-bash-captures-output
  (is (= {:output "hi" :exit-code 0} (agent/run-bash "printf hi"))))

(deftest react-loop-runs-a-tool-then-answers
  (let [respond (fn [history]
                  (if (= 1 (count history))
                    [{:type "function_call"
                      :name "bash"
                      :arguments "{\"command\":\"echo 42\"}"
                      :call_id "c1"}]
                    [{:type "message"
                      :content [{:type "output_text" :text "done: 42"}]}]))
        result (agent/react-loop respond (fn [_] "exit code: 0\n42")
                                  [{:type "message" :role "user" :content "go"}])]
    (is (= "done: 42" (:answer result)))
    (is (= 1 (:rounds result)))
    (is (= 4 (count (:history result))))))

(deftest react-loop-appends-assistant-message
  (let [respond (fn [_]
                  [{:type "message"
                    :content [{:type "output_text" :text "blue"}]}])
        result (agent/react-loop
                respond (fn [_] "")
                [{:type "message" :role "user" :content "favorite color?"}])]
    (is (= [{:type "message" :role "user" :content "favorite color?"}
            {:type "message" :role "assistant"
             :content [{:type "output_text" :text "blue"}]}]
           (:history result)))))

(deftest ralph-loop-verifies-before-completing
  (let [prompts (atom [])
        calls (atom 0)
        result (agent/ralph-loop
                {:task "fix it"
                 :max-iterations 3
                 :iterate (fn [prompt]
                            (swap! prompts conj prompt)
                            {:answer (if (= 1 (swap! calls inc))
                                       "implemented"
                                       "<promise>COMPLETE</promise>")
                             :history []})
                 :verify (fn [_] {:exit-code 0 :output "ok"})})]
    (is (= :complete (:status result)))
    (is (= 2 (:iterations result)))
    (is (= 2 (count @prompts)))
    (is (re-find #"Previous iteration feedback" (second @prompts)))))

(deftest ralph-loop-rejects-completion-with-failing-verification
  (let [calls (atom 0)
        result (agent/ralph-loop
                {:task "fix it"
                 :max-iterations 2
                 :iterate (fn [_]
                            (swap! calls inc)
                            {:answer "<promise>COMPLETE</promise>"})
                 :verify (fn [_] {:exit-code 1 :output "failed"})})]
    (is (= :max-iterations (:status result)))
    (is (= 2 @calls))))

(deftest non-complete-answer-stops-at-limit
  (let [calls (atom 0)
        result (agent/ralph-loop
                {:task "unfinished"
                 :max-iterations 2
                 :iterate (fn [_]
                            (swap! calls inc)
                            {:answer "still working"})})]
    (is (= :max-iterations (:status result)))
    (is (= 2 @calls))))

(deftest parse-ralph-arguments
  (is (= {:max-iterations 4
          :verify-command "lein test"
          :task "make it pass"}
         (agent/parse-ralph-args
          ["--max-iterations" "4" "--verify" "lein test" "make" "it" "pass"]))))

(deftest parse-ralph-rejects-missing-option-values
  (is (thrown? clojure.lang.ExceptionInfo
               (agent/parse-ralph-args ["--max-iterations"])))
  (is (thrown? clojure.lang.ExceptionInfo
               (agent/parse-ralph-args ["--verify"])))
  (is (thrown? clojure.lang.ExceptionInfo
               (agent/parse-ralph-args ["--max-iterations" "0" "task"]))))
