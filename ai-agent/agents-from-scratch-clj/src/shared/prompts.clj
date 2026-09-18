(ns shared.prompts
  (:require [clojure.string :as string]))

(defn base-prompt [user-input] user-input)

(defn system-prompt [role user-input]
  (str "<SYSTEM>\n" role "\n</SYSTEM>\n\n<USER>\n" user-input "\n</USER>"))

(defn json-contract [schema content]
  (str "Return ONLY valid JSON.\nNo explanations. No markdown. No extra text.\n\n"
       "Schema:\n" schema "\n\nContent:\n" content))

(defn decision-prompt [choices user-input]
  (let [options (string/join "\n" (map #(str "- " %) choices))]
    (str "You must choose ONE of the following options.\nReturn ONLY valid JSON.\n\n"
         "Available choices:\n" options "\n\nSchema:\n{ \"decision\": string }\n\n"
         "Input:\n" user-input)))

(defn tool-call-prompt [tools user-input]
  (str "You may request ONE tool call.\n\nAvailable tools:\n" tools "\n\n"
       "Return ONLY valid JSON.\n\nSchema:\n{\n  \"tool\": string,\n  \"arguments\": object\n}\n\n"
       "User request:\n" user-input))

(defn agent-step-prompt [state user-input]
  (str "You are an agent.\n\nCurrent state:\n" state "\n\n"
       "Decide the next action.\n\nReturn ONLY valid JSON.\n\nSchema:\n{\n  \"action\": string,\n  \"reason\": string\n}\n\n"
       "User input:\n" user-input))

(defn memory-prompt [state memory user-input]
  (str "You are an agent with memory.\n\nCurrent state:\n" state "\n\nRelevant memory:\n"
       memory "\n\nDecide what to do next.\n\nReturn ONLY valid JSON.\n\n"
       "Schema:\n{\n  \"action\": string,\n  \"save_to_memory\": string | null\n}\n\n"
       "User input:\n" user-input))

(defn planning-prompt [goal]
  (str "Create a step-by-step plan to achieve the goal.\n\nReturn ONLY valid JSON.\n\n"
       "Schema:\n{\n  \"steps\": [string]\n}\n\nGoal:\n" goal))

(defn atomic-action-prompt [step]
  (str "Convert this step into an atomic action.\n\nReturn ONLY valid JSON.\n\n"
       "Schema:\n{\n  \"action\": string,\n  \"inputs\": object\n}\n\nStep:\n" step))

(defn aot-prompt [goal]
  (str "Create an atomic execution graph for the goal.\n\nReturn ONLY valid JSON.\n\n"
       "Schema:\n{\n  \"nodes\": [\n    {\n      \"id\": string,\n      \"action\": string,\n      \"depends_on\": [string]\n    }\n  ]\n}\n\n"
       "Goal:\n" goal))
