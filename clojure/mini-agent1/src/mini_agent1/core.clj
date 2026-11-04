(ns mini-agent1.core
  "A minimal ReAct agent on DeepSeek's Responses API, reusing Codex's key/URL."
  (:require [clojure.data.json :as json]
            [clojure.string :as str])
  (:import [java.net.http HttpClient HttpRequest HttpRequest$BodyPublishers HttpResponse$BodyHandlers]
           [java.net URI]
           [java.time Duration]
           [java.util.concurrent Executors ThreadFactory]))

;; ---- config: env vars override Codex's defaults ----

(defn codex-api-key
  "Codex's bearer token from ~/.codex/config.toml, used as the key fallback."
  []
  (let [f (java.io.File. (str (System/getProperty "user.home") "/.codex/config.toml"))]
    (when (.isFile f)
      (second (re-find #"(?m)^experimental_bearer_token\s*=\s*\"([^\"]+)\"" (slurp f))))))

(defn api-key [] (or (not-empty (System/getenv "DEEPSEEK_API_KEY")) (codex-api-key)))
(defn base-url [] (or (not-empty (System/getenv "DEEPSEEK_BASE_URL")) "https://api.deepseek.com/"))
(defn model [] (or (not-empty (System/getenv "DEEPSEEK_MODEL")) "deepseek-flash"))

(defn context-window
  "Context window in tokens for the active model; DEEPSEEK_CONTEXT_WINDOW overrides."
  []
  (or (when-let [v (not-empty (System/getenv "DEEPSEEK_CONTEXT_WINDOW"))]
        (try (Long/parseLong v) (catch Exception _ nil)))
      (get {"deepseek-chat" 131072
            "deepseek-reasoner" 131072
            "deepseek-flash" 1048576
            "deepseek-v4-flash" 1048576
            "deepseek-v4-pro" 1048576}
           (model))))

;; ---- HTTP ----

(defn- responses-endpoint [base]
  (str (str/replace base #"/+$" "") "/responses"))

(def ^:private http-client
  ;; daemon executor so a finished CLI exits instead of hanging on idle pools
  (delay (-> (HttpClient/newBuilder)
             (.executor (Executors/newCachedThreadPool
                         (reify ThreadFactory
                           (newThread [_ r] (doto (Thread. r) (.setDaemon true))))))
             (.build))))

(defn- call-model
  "POST one Responses-API request; returns the parsed body map."
  [key base body]
  (let [req (-> (HttpRequest/newBuilder)
                (.uri (URI/create (responses-endpoint base)))
                (.timeout (Duration/ofSeconds 120))
                (.header "Content-Type" "application/json")
                (.header "Authorization" (str "Bearer " key))
                (.POST (HttpRequest$BodyPublishers/ofString (json/write-str body)))
                (.build))
        resp (try
               (.send @http-client req (HttpResponse$BodyHandlers/ofString))
               (catch Exception e
                 (throw (ex-info (str "API request failed: " (.getMessage e)) {} e))))
        code (.statusCode resp)
        text (.body resp)]
    (if (= 200 code)
      (json/read-str text :key-fn keyword)
      (throw (ex-info (str "API HTTP " code ": " text) {:status code :body text})))))

;; ---- tools ----

(defn tools
  "One tool is enough for a task-capable agent: a sandbox-free bash shell."
  []
  [{:type "function"
    :name "bash"
    :description (str "Run a shell command (bash -lc) and return merged stdout+stderr and the exit code. "
                      "Use it to list/read files (ls, cat), run tests, and inspect the project.")
    :parameters {:type "object"
                 :properties {:command {:type "string" :description "Shell command to execute."}}
                 :required ["command"]}}])

(defn run-bash
  "Run a shell command with a 60s timeout; returns {:output ... :exit-code ...}."
  [command]
  (let [p (-> (ProcessBuilder. ["bash" "-lc" command])
              (.redirectErrorStream true)
              (.start))
        out (promise)
        _ (doto (Thread. (fn [] (deliver out (slurp (.getInputStream p)))))
            (.setDaemon true)
            (.start))]
    (if (= ::timeout (deref out 60000 ::timeout))
      (do (.destroyForcibly p)
          {:output "" :exit-code 124})
      {:output (str/trim @out) :exit-code (do (.waitFor p) (.exitValue p))})))

(defn execute-tool
  "Run one tool call; returns the observation string fed back to the model."
  [{:keys [name arguments]}]
  (try
    (case name
      "bash" (let [command (get (json/read-str arguments) "command")]
               (let [{:keys [output exit-code]} (run-bash command)]
                 (format "exit code: %d\n%s" exit-code output)))
      (str "unknown tool: " name))
    (catch Exception e
      (str "tool error: " (.getMessage e)))))

;; ---- token usage ----

(defn- api-usage
  "Map a Responses-API `usage` object to {:input :output :total} token counts."
  [u]
  (when u
    {:input (long (get u :input_tokens 0))
     :output (long (get u :output_tokens 0))
     :total (long (get u :total_tokens 0))}))

;; ---- ReAct loop ----

(def max-tool-rounds
  "Safety bound on tool-call rounds per run."
  10)

(defn function-calls
  "Tool-call items from a Responses-API output (reasoning items are skipped)."
  [items]
  (filter #(= "function_call" (:type %)) items))

(defn answer-text
  "Final answer text from message items (reasoning/encrypted items ignored)."
  [items]
  (->> items
       (filter #(= "message" (:type %)))
       (mapcat :content)
       (filter #(= "output_text" (:type %)))
       (map :text)
       str/join))

(defn assistant-message
  "A cleaned assistant message item, safe to feed back as input next turn."
  [text]
  {:type "message" :role "assistant" :content [{:type "output_text" :text text}]})

(defn react-loop
  "Run the ReAct cycle over `history` (input items). `respond` maps history ->
   output items; `execute` maps a function_call item -> an observation string.
   Returns {:answer ... :rounds ... :history ...}, where :history includes this
   turn's assistant answer so callers can thread it into the next turn."
  [respond execute history]
  (loop [history history
         round 0]
    (let [items (respond history)
          calls (vec (function-calls items))]
      (if (empty? calls)
        (let [text (answer-text items)]
          {:answer text
           :rounds round
           :history (if (seq text) (conj history (assistant-message text)) history)})
        (if (>= round max-tool-rounds)
          {:answer (str "stopped after " max-tool-rounds " tool rounds")
           :rounds round
           :history history}
          (let [call-items (map #(select-keys % [:type :name :arguments :call_id]) calls)
                outputs (map (fn [c] {:type "function_call_output"
                                      :call_id (:call_id c)
                                      :output (execute c)})
                             calls)]
            (recur (into (into history call-items) outputs) (inc round))))))))

(defn run-agent
  "Run one ReAct turn against the live API, optionally continuing `history`
   from earlier turns. Prints tool activity; returns the updated history."
  ([prompt] (run-agent prompt []))
  ([prompt history]
   (let [key (api-key)]
     (when-not key
       (throw (ex-info "No API key: set DEEPSEEK_API_KEY or configure ~/.codex/config.toml" {})))
     (let [usage (atom nil)
           respond (fn [history]
                     (let [body (call-model key (base-url)
                                            {:model (model)
                                             :input history
                                             :tools (tools)
                                             :max_output_tokens 2048})]
                       (reset! usage (api-usage (:usage body)))
                       (:output body)))
           execute (fn [call]
                     (println (format "[tool] %s %s" (:name call) (:arguments call)))
                     (let [out (execute-tool call)]
                       (println "[result]" out)
                       out))]
       (assoc (react-loop respond execute
                          (conj (vec history) {:type "message" :role "user" :content prompt}))
              :usage @usage)))))

;; ---- CLI ----

(defn- print-result [{:keys [answer rounds usage]}]
  (println "\n=== ANSWER ===")
  (println answer)
  (println (str "\n(tool rounds: " rounds ")"))
  (when usage
    (let [win (context-window)
          pct (when win (* 100.0 (/ (:input usage) win)))]
      (println (format "(usage: %d input, %d output, %d total%s)"
                       (:input usage) (:output usage) (:total usage)
                       (if pct (format ", %.1f%% context used" pct) ""))))))

(defn -main
  "With argv: run one task. Without: loop on stdin until EOF or quit/exit/q."
  [& args]
  (if (seq args)
    (print-result (run-agent (str/join " " args)))
    (loop [history []]
      (print "task> ") (flush)
      (when-let [line (read-line)]
        (when-not (contains? #{"quit" "exit" "q"} (str/trim line))
          (let [result (run-agent line history)]
            (print-result result)
            (recur (:history result))))))))
