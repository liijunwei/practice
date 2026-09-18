(ns complete-example
  (:require [agent.agent :as agent]
            [agent.evals :as evals]
            [agent.memory :as memory]
            [agent.telemetry :as telemetry]
            [evals.golden-datasets :as golden]))

(defn- section [title]
  (println "\n" (apply str (repeat 50 "=")))
  (println title)
  (println (apply str (repeat 50 "="))))

(defn lesson-01-basic-chat []
  (section "LESSON 01: Basic LLM Chat")
  (let [a (agent/new-agent)]
    (println "Response:" (agent/simple-generate a "Explain what an AI agent is?"))))

(defn lesson-02-with-role []
  (section "LESSON 02: With System Prompt")
  (let [a (agent/new-agent)]
    (println "Response:" (agent/generate-with-role a "Explain what an AI agent is?"))))

(defn lesson-03-structured []
  (section "LESSON 03: Structured Output")
  (let [a (agent/new-agent)
        schema "{\n  \"topic\": string,\n  \"difficulty\": \"beginner\" | \"intermediate\" | \"advanced\"\n}"]
    (println "Structured result:" (agent/generate-structured a "Explain quantum computing" schema))))

(defn lesson-04-decisions []
  (section "LESSON 04: Decision Making")
  (let [a (agent/new-agent)]
    (println "Decision:"
             (agent/decide a "Can you summarize this article for me?"
                           ["answer_question" "summarize_text" "translate"]))))

(defn lesson-05-tools []
  (section "LESSON 05: Tool Calling")
  (let [a (agent/new-agent)
        tool-call (agent/request-tool a "What is 42 * 7?")]
    (println "Tool request:" tool-call)
    (when tool-call
      (println "Tool result:" (agent/execute-tool-call a tool-call)))))

(defn lesson-06-agent-loop []
  (section "LESSON 06: Agent Loop")
  (let [a (agent/new-agent)
        results (agent/run-loop a "Help me understand loops" 3)]
    (println "\nNote: Repetition in early iterations is expected.")
    (doseq [[i result] (map-indexed vector results)]
      (println "Iteration" (inc i) ":")
      (println "  Action:" (get result "action" "unknown"))
      (println "  Reason:" (get result "reason" "No reason provided")))))

(defn lesson-07-memory []
  (section "LESSON 07: Memory")
  (let [a (agent/new-agent)
        response1 (agent/run-with-memory a "My name is Alice")
        response2 (agent/run-with-memory a "What's my name?")]
    (println "Response 1:" (get response1 "reply" response1))
    (when-let [saved (get response1 "save_to_memory")]
      (println "  → Saved to memory:" saved))
    (println "Response 2:" (get response2 "reply" response2))
    (when-let [saved (get response2 "save_to_memory")]
      (println "  → Saved to memory:" saved))
    (println "\nMemory contents:" (memory/get-all (:memory a)))))

(defn lesson-08-planning []
  (section "LESSON 08: Planning")
  (let [a (agent/new-agent)
        plan (agent/create-plan a "Write a blog post about AI agents")]
    (println "Plan:" plan)
    (when plan
      (println "Execution results:" (agent/execute-plan a plan)))))

(defn lesson-09-atomic-actions []
  (section "LESSON 09: Atomic Actions")
  (let [a (agent/new-agent)
        step "Write an explanation of AI agents"
        atomic-action (agent/create-atomic-action a step)]
    (println "Step:" step)
    (println "Atomic action:" atomic-action)
    (let [plan (agent/create-plan a "Create a tutorial about Python")]
      (when (and plan (seq (get plan "steps")))
        (let [first-step (first (get plan "steps"))]
          (println "\nPlan step:" first-step)
          (println "Atomic action from plan step:" (agent/create-atomic-action a first-step)))))))

(defn lesson-10-aot []
  (section "LESSON 10: Atom of Thought")
  (let [a (agent/new-agent)
        graph (agent/create-aot-plan a "Research and write article")]
    (println "AoT graph:" graph)
    (when graph
      (println "Execution results:" (agent/execute-aot-plan a graph)))))

(defn lesson-11-evals []
  (section "LESSON 11: Evals")
  (let [a (agent/new-agent)
        results (evals/run-all
                 a
                 :structured-cases (take 2 golden/structured-output-golden)
                 :tool-cases (take 2 golden/tool-call-golden)
                 :decision-cases (take 2 golden/decision-golden)
                 :memory-cases (take 1 golden/memory-golden))]
    (println "\nRunning eval suites...")
    (evals/print-eval-report results)))

(defn lesson-12-telemetry []
  (section "LESSON 12: Telemetry")
  (let [a (agent/new-agent)
        tel (telemetry/new-telemetry "agent_telemetry.jsonl")]
    (telemetry/clear! tel)
    (let [trace-id (telemetry/start-trace! tel)]
      (println "Trace ID:" trace-id)
      (let [start (System/currentTimeMillis)
            result1 (agent/generate-structured
                     a "What is Python?"
                     "{\"answer\": string, \"difficulty\": \"beginner\" | \"intermediate\" | \"advanced\"}")
            duration1 (- (System/currentTimeMillis) start)]
        (telemetry/log-llm-call! tel 150 (count (str result1)) duration1
                                 :success (some? result1)
                                 :error (when-not result1 "Failed to parse JSON"))
        (println "\n1. Structured output call...")
        (println "   Result:" result1)
        (println "   Duration:" duration1 "ms"))
      (let [start (System/currentTimeMillis)
            tool-call (agent/request-tool a "What is 15 * 8?")
            duration2 (- (System/currentTimeMillis) start)]
        (telemetry/log-llm-call! tel 200 (count (str tool-call)) duration2
                                 :success (some? tool-call))
        (when tool-call
          (telemetry/log-tool-call! tel (get tool-call "tool" "unknown")
                                    (get tool-call "arguments" {})
                                    :result (agent/execute-tool-call a tool-call)
                                    :duration-ms 1.0)
          (println "\n2. Tool call...")
          (println "   Tool:" tool-call)))
      (let [start (System/currentTimeMillis)
            result3 (agent/run-with-memory a "My favorite color is blue")
            duration3 (- (System/currentTimeMillis) start)]
        (telemetry/log-llm-call! tel 300 (count (str result3)) duration3
                                 :success (some? result3))
        (telemetry/log-memory-operation! tel "add" "favorite color is blue")
        (println "\n3. Memory operation...")
        (println "   Result:" result3))
      (telemetry/print-summary tel)
      (println "\nRecent spans:")
      (doseq [span (telemetry/get-recent-spans tel 5)]
        (println "  [" (get span "event_type" "unknown") "] duration="
                 (get span "duration_ms" "N/A") "ms"))
      (println "\nTelemetry logged to: agent_telemetry.jsonl"))))

(defn -main [& _]
  (println "\n" (apply str (repeat 50 "#")))
  (println "# AI Agent Examples - All Lessons")
  (println (apply str (repeat 50 "#")))
  (try
    (lesson-01-basic-chat)
    (lesson-02-with-role)
    (lesson-03-structured)
    (lesson-04-decisions)
    (lesson-05-tools)
    (lesson-06-agent-loop)
    (lesson-07-memory)
    (lesson-08-planning)
    (lesson-09-atomic-actions)
    (lesson-10-aot)
    (lesson-11-evals)
    (lesson-12-telemetry)
    (println "\n" (apply str (repeat 50 "=")))
    (println "All examples completed!")
    (println (apply str (repeat 50 "=")))
    (catch Exception e
      (println "\n❌ Error:" (.getMessage e))
      (println "\nMake sure you have:")
      (println "1. Exported DEEPSEEK_API_KEY")
      (println "2. Network access to api.deepseek.com")
      (println "3. Optional: exported DEEPSEEK_MODEL")))
  (shutdown-agents))
