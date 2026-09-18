(ns agent.agent
  (:require [shared.llm :as llm]
            [shared.utils :as utils]
            [agent.state :as state]
            [agent.memory :as memory]
            [agent.tools :as tools]
            [agent.planner :as planner]
            [clojure.string :as string]))

(def system-prompt-text
  (str "You are a calm, precise, and helpful AI assistant. "
       "You explain concepts simply and avoid unnecessary jargon. "
       "You are honest about what you know and don't know."))

(defn new-agent
  ([] (new-agent nil))
  ([model-path]
   {:llm (llm/new-llm model-path)
    :system-prompt system-prompt-text
    :state (state/new-state)
    :memory (memory/new-memory)}))

(defn- retry-json
  [agent prompt valid?]
  (loop [attempt 0]
    (when (< attempt 3)
      (let [parsed (utils/extract-json-from-text
                    (llm/generate (:llm agent) prompt 0.0 nil))]
        (if (valid? parsed)
          parsed
          (recur (inc attempt)))))))

;; Lesson 01
(defn simple-generate [agent user-input]
  (llm/generate (:llm agent) user-input))

;; Lesson 02
(defn generate-with-role [agent user-input]
  (let [prompt (str (:system-prompt agent) "\n\nUser: " user-input "\nAssistant:")
        response (or (llm/generate (:llm agent) prompt) "")]
    (-> response
        (string/replace "<SYSTEM>" "")
        (string/replace "</SYSTEM>" "")
        (string/replace "<USER>" "")
        (string/replace "</USER>" "")
        string/trim)))

;; Lesson 03
(defn generate-structured [agent user-input schema]
  (let [prompt (str (:system-prompt agent)
                    "\n\nCRITICAL INSTRUCTIONS:\n"
                    "1. Respond with ONLY valid JSON\n"
                    "2. No explanations, no markdown, no extra text before or after the JSON\n"
                    "3. Start your response with { and end with }\n\n"
                    "Schema you must follow:\n" schema "\n\n"
                    "User request: " user-input "\n\nResponse (JSON only):")]
    (retry-json agent prompt some?)))

;; Lesson 04
(defn decide [agent user-input choices]
  (let [options (string/join "\n" (map #(str "- " %) choices))
        prompt (str (:system-prompt agent)
                    "\n\nYou must choose ONE of the following options. Respond with ONLY valid JSON.\n\n"
                    "CRITICAL INSTRUCTIONS:\n"
                    "1. Respond with ONLY valid JSON\n"
                    "2. No explanations, no markdown, no other text\n"
                    "3. Start your response with { and end with }\n\n"
                    "Available choices:\n" options "\n\n"
                    "Required JSON format:\n"
                    "{\"decision\": \"one_of_the_choices_above\"}\n\n"
                    "User request: " user-input "\n\nResponse (JSON only):")]
    (retry-json
     agent
     prompt
     (fn [parsed]
       (and (contains? parsed "decision")
            (some #(= (get parsed "decision") %) choices))))))

;; Lesson 05
(defn request-tool [agent user-input]
  (let [prompt (str (:system-prompt agent)
                    "\n\nYou are a tool-calling assistant. When asked a math question, "
                    "you must respond with ONLY valid JSON.\n\n"
                    "Available tool: calculator\n"
                    "- Parameters: a (number), b (number), operation (\"add\", \"subtract\", \"multiply\", or \"divide\")\n\n"
                    "CRITICAL INSTRUCTIONS:\n"
                    "1. Respond with ONLY valid JSON\n"
                    "2. No explanations, no markdown, no other text\n"
                    "3. Start your response with { and end with }\n\n"
                    "Example format:\n"
                    "{\"tool\": \"calculator\", \"arguments\": {\"a\": 42, \"b\": 7, \"operation\": \"multiply\"}}\n\n"
                    "User request: " user-input "\n\nResponse (JSON only):")]
    (retry-json
     agent
     prompt
     (fn [parsed]
       (and (contains? parsed "tool") (contains? parsed "arguments"))))))

(defn execute-tool-call [agent tool-call]
  (tools/execute-tool (get tool-call "tool") (get tool-call "arguments")))

;; Lesson 06
(defn agent-step [agent user-input]
  (let [state-map (state/state-map (:state agent))
        prompt (str (:system-prompt agent)
                    "\n\nYou are an agent. You must decide the next action and respond with ONLY valid JSON.\n\n"
                    "Current state: steps=" (get state-map :steps 0)
                    ", done=" (get state-map :done false) "\n\n"
                    "Available actions: analyze, research, summarize, answer, done\n\n"
                    "CRITICAL INSTRUCTIONS:\n"
                    "1. Respond with ONLY valid JSON\n"
                    "2. No explanations, no markdown, no other text\n"
                    "3. Start your response with { and end with }\n\n"
                    "Required JSON format:\n"
                    "{\"action\": \"action_name\", \"reason\": \"explanation\"}\n\n"
                    "User input: " user-input "\n\nResponse (JSON only):")]
    (retry-json
     agent
     prompt
     (fn [parsed]
       (when (contains? parsed "action")
         (when-not (contains? parsed "reason")
           (assoc parsed "reason" (str "Taking action: " (get parsed "action"))))
         (state/increment-step! (:state agent))
         true)))))

(defn run-loop
  ([agent user-input] (run-loop agent user-input 5))
  ([agent user-input max-steps]
   (state/reset-state! (:state agent))
   (loop [results []]
     (let [{:keys [done steps]} (state/state-map (:state agent))]
       (if (or done (>= steps max-steps))
         results
         (if-let [action (agent-step agent user-input)]
           (do
             (when (= (get action "action") "done")
               (state/mark-done! (:state agent)))
             (recur (conj results action)))
           results))))))

;; Lesson 07
(defn run-with-memory [agent user-input]
  (let [memory-context (memory/get-all (:memory agent))
        memory-str (if (seq memory-context)
                     (str "You remember the following:\n"
                          (string/join "\n" (map #(str "- " %) memory-context)))
                     "You have no memories yet.")
        prompt (str (:system-prompt agent)
                    "\n\nYou are an agent with memory. You must respond with ONLY valid JSON.\n\n"
                    memory-str "\n\n"
                    "CRITICAL INSTRUCTIONS:\n"
                    "1. Respond with ONLY valid JSON\n"
                    "2. No explanations, no markdown, no other text\n"
                    "3. Start your response with { and end with }\n"
                    "4. If the user tells you information (like their name), save it to memory\n"
                    "5. If the user asks about something you remember, USE YOUR MEMORY to answer\n\n"
                    "Required JSON format:\n"
                    "{\"reply\": \"your response text\", \"save_to_memory\": \"fact to remember\" or null}\n\n"
                    "Examples:\n"
                    "- User says \"My name is Alice\" → {\"reply\": \"Nice to meet you, Alice!\", "
                    "\"save_to_memory\": \"User's name is Alice\"}\n"
                    "- User asks \"What's my name?\" and you remember \"User's name is Alice\" → "
                    "{\"reply\": \"Your name is Alice\", \"save_to_memory\": null}\n\n"
                    "User input: " user-input "\n\nResponse (JSON only):")]
    (retry-json
     agent
     prompt
     (fn [parsed]
       (when (contains? parsed "reply")
         (when-let [fact (get parsed "save_to_memory")]
           (memory/add! (:memory agent) fact))
         (state/increment-step! (:state agent))
         true)))))

;; Lesson 08
(defn create-plan [agent goal]
  (let [plan (planner/create-plan (:llm agent) goal)]
    (when plan
      (swap! (:state agent) assoc :current-plan plan))
    plan))

(defn execute-plan [agent plan]
  (if-not (and plan (contains? plan "steps"))
    []
    (mapv (fn [step]
            (state/increment-step! (:state agent))
            {"step" step "executed" true})
          (get plan "steps"))))

;; Lesson 09
(defn create-atomic-action [agent step]
  (planner/create-atomic-action (:llm agent) step))

;; Lesson 10
(defn create-aot-plan [agent goal]
  (planner/create-aot-graph (:llm agent) goal))

(defn execute-aot-plan [agent graph]
  (planner/execute-graph graph (fn [action] (str "Executed: " action))))

(defn run [agent user-input]
  (let [result (run-with-memory agent user-input)]
    (if (and result (contains? result "reply"))
      (get result "reply")
      (generate-with-role agent user-input))))
