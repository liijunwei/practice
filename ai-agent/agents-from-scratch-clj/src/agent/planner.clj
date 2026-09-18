(ns agent.planner
  (:require [shared.llm :as llm]
            [shared.utils :as utils]))

(defn- retry-json
  [llm prompt valid?]
  (loop [attempt 0]
    (when (< attempt 3)
      (let [parsed (utils/extract-json-from-text
                    (llm/generate llm prompt 0.0 nil))]
        (if (valid? parsed)
          parsed
          (recur (inc attempt)))))))

(defn create-plan
  [llm goal]
  (let [prompt (str "Create a step-by-step plan to achieve the goal. Respond with ONLY valid JSON.\n\n"
                    "CRITICAL INSTRUCTIONS:\n"
                    "1. Respond with ONLY valid JSON\n"
                    "2. No explanations, no markdown, no other text\n"
                    "3. Start your response with { and end with }\n\n"
                    "Required JSON format:\n"
                    "{\"steps\": [\"step1\", \"step2\", \"step3\"]}\n\n"
                    "Goal: " goal "\n\nResponse (JSON only):")]
    (retry-json llm prompt #(and (map? %) (vector? (get % "steps"))))))

(defn create-atomic-action
  [llm step]
  (let [prompt (str "Convert this step into an atomic action. Respond with ONLY valid JSON.\n\n"
                    "CRITICAL INSTRUCTIONS:\n"
                    "1. Respond with ONLY valid JSON\n"
                    "2. No explanations, no markdown, no other text\n"
                    "3. Start your response with { and end with }\n\n"
                    "Required JSON format:\n"
                    "{\n  \"action\": \"action_name\",\n  \"inputs\": {\"key\": \"value\"}\n}\n\n"
                    "The action should be a simple, atomic operation name.\n"
                    "The inputs should be a dictionary with the parameters needed for this action.\n\n"
                    "Step to convert:\n" step "\n\nResponse (JSON only):")]
    (retry-json llm prompt #(and (map? %) (contains? % "action")))))

(defn create-aot-graph
  [llm goal]
  (let [prompt (str "Create an atomic execution graph for the goal. Each node is a single action. "
                    "Dependencies are node IDs. Respond with ONLY valid JSON.\n\n"
                    "CRITICAL INSTRUCTIONS:\n"
                    "1. Respond with ONLY valid JSON\n"
                    "2. No explanations, no markdown, no other text\n"
                    "3. Start your response with { and end with }\n\n"
                    "Required JSON format:\n"
                    "{\"nodes\": [{\"id\": \"1\", \"action\": \"research\", \"depends_on\": []}, "
                    "{\"id\": \"2\", \"action\": \"write\", \"depends_on\": [\"1\"]}]}\n\n"
                    "Each node must have:\n"
                    "- id: unique string like \"1\", \"2\", \"3\"\n"
                    "- action: what to do (e.g., \"research\", \"write\", \"review\")\n"
                    "- depends_on: list of node IDs that must complete first (empty [] for first step)\n\n"
                    "Goal: " goal "\n\nResponse (JSON only):")]
    (retry-json
     llm
     prompt
     (fn [graph]
       (and (map? graph)
            (vector? (get graph "nodes"))
            (every? (fn [node]
                      (and (map? node)
                           (contains? node "id")
                           (contains? node "action")
                           (contains? node "depends_on")
                           (vector? (get node "depends_on"))))
                    (get graph "nodes")))))))

(defn execute-graph
  [graph executor-fn]
  (let [nodes (get graph "nodes" [])]
    (loop [nodes nodes
           executed #{}
           results []
           iteration 0]
      (if (or (= (count executed) (count nodes))
              (>= iteration (* 2 (count nodes))))
        results
        (if-let [node (first (filter (fn [n]
                                       (let [id (get n "id")]
                                         (and (not (contains? executed id))
                                              (every? executed (get n "depends_on")))))
                                     nodes))]
          (let [id (get node "id")
                action (get node "action")
                outcome (try
                          {"result" (executor-fn action) "success" true}
                          (catch Exception e
                            {"error" (.getMessage e) "success" false}))]
            (recur nodes
                   (conj executed id)
                   (conj results (merge {"node_id" id "action" action} outcome))
                   (inc iteration)))
          results)))))
