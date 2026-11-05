(ns mini-agent2.core
  "A minimal ReAct agent with a verification-gated Ralph loop."
  (:require [clojure.data.json :as json]
            [clojure.java.io :as io]
            [clojure.string :as str])
  (:import [java.net URI]
           [java.net.http HttpClient HttpRequest HttpRequest$BodyPublishers
            HttpResponse HttpResponse$BodyHandlers]
           [java.time Duration]
           [java.util.concurrent Executors ThreadFactory]))

;; ---- config: environment variables override Codex's defaults ----

(defn codex-api-key []
  (let [f (io/file (System/getProperty "user.home") ".codex/config.toml")]
    (when (.isFile f)
      (second (re-find #"(?m)^experimental_bearer_token\s*=\s*\"([^\"]+)\"" (slurp f))))))

(defn api-key []
  (or (not-empty (System/getenv "MINI_AGENT2_API_KEY"))
      (not-empty (System/getenv "DEEPSEEK_API_KEY"))
      (codex-api-key)))

(defn base-url []
  (or (not-empty (System/getenv "MINI_AGENT2_BASE_URL"))
      (not-empty (System/getenv "DEEPSEEK_BASE_URL"))
      "https://api.deepseek.com/"))

(defn model []
  (or (not-empty (System/getenv "MINI_AGENT2_MODEL"))
      (not-empty (System/getenv "DEEPSEEK_MODEL"))
      "deepseek-flash"))

(defn context-window []
  (or (when-let [v (not-empty (System/getenv "MINI_AGENT2_CONTEXT_WINDOW"))]
        (try (Long/parseLong v) (catch Exception _ nil)))
      (get {"deepseek-chat" 131072
            "deepseek-reasoner" 131072
            "deepseek-flash" 1048576
            "deepseek-v4-flash" 1048576
            "deepseek-v4-pro" 1048576}
           (model))))

(defonce mem (atom []))

;; ---- HTTP ----

(defn- responses-endpoint [base]
  (str (str/replace base #"/+$" "") "/responses"))

(def ^:private http-client
  (delay (-> (HttpClient/newBuilder)
             (.executor (Executors/newCachedThreadPool
                         (reify ThreadFactory
                           (newThread [_ r] (doto (Thread. r) (.setDaemon true))))))
             (.build))))

(defn- instructions []
  (str "You are mini-agent2, a concise coding agent working in "
       (System/getProperty "user.dir") ". "
       "Use bash to inspect and change files. Run relevant tests before claiming success. "
       "Never expose API keys or other credentials."))

(defn- call-model [key base body]
  (let [^HttpRequest req (-> (HttpRequest/newBuilder)
                             (.uri (URI/create (responses-endpoint base)))
                             (.timeout (Duration/ofSeconds 120))
                             (.header "Content-Type" "application/json")
                             (.header "Authorization" (str "Bearer " key))
                             (.POST (HttpRequest$BodyPublishers/ofString (json/write-str body)))
                             (.build))
        ^HttpResponse resp (try
                             (.send ^HttpClient @http-client req (HttpResponse$BodyHandlers/ofString))
                             (catch Exception e
                               (throw (ex-info (str "API request failed: " (.getMessage e)) {} e))))
        code (.statusCode resp)
        text (.body resp)]
    (if (= 200 code)
      (json/read-str text :key-fn keyword)
      (throw (ex-info (str "API HTTP " code ": " text) {:status code :body text})))))

;; ---- tools ----

(defn tools []
  [{:type "function"
    :name "bash"
    :description (str "Run a shell command (bash -lc) in the current project and return "
                      "merged stdout+stderr plus the exit code. Use it to inspect, edit, "
                      "and test files.")
    :parameters {:type "object"
                 :properties {:command {:type "string" :description "Shell command to execute."}}
                 :required ["command"]}}])

(defn run-bash [command]
  (let [p (-> (ProcessBuilder. (java.util.List/of "bash" "-lc" command))
              (.redirectErrorStream true)
              (.start))
        out (promise)
        _ (doto (Thread. (fn [] (deliver out (slurp (.getInputStream p)))))
            (.setDaemon true)
            (.start))]
    (if (= ::timeout (deref out 60000 ::timeout))
      (do (.destroyForcibly p)
          {:output "command timed out" :exit-code 124})
      {:output (str/trim @out) :exit-code (do (.waitFor p) (.exitValue p))})))

(defn execute-tool [{:keys [name arguments]}]
  (try
    (case name
      "bash" (let [command (get (json/read-str arguments) "command")
                   {:keys [output exit-code]} (run-bash command)]
               (format "exit code: %d\n%s" exit-code output))
      (str "unknown tool: " name))
    (catch Exception e
      (str "tool error: " (.getMessage e)))))

;; ---- ReAct ----

(def max-tool-rounds 20)

(defn function-calls [items]
  (filter #(= "function_call" (:type %)) items))

(defn answer-text [items]
  (->> items
       (filter #(= "message" (:type %)))
       (mapcat :content)
       (filter #(= "output_text" (:type %)))
       (map :text)
       str/join))

(defn assistant-message [text]
  {:type "message" :role "assistant" :content [{:type "output_text" :text text}]})

(defn api-usage [u]
  (when u
    {:input (long (get u :input_tokens 0))
     :output (long (get u :output_tokens 0))
     :total (long (get u :total_tokens 0))}))

(defn react-loop
  ([respond execute history] (react-loop respond execute history max-tool-rounds))
  ([respond execute history round-limit]
   (loop [history history
          round 0]
     (let [items (respond history)
           calls (vec (function-calls items))]
       (if (empty? calls)
         (let [text (answer-text items)]
           {:answer text
            :rounds round
            :history (if (seq text) (conj history (assistant-message text)) history)})
         (if (>= round round-limit)
           {:answer (str "stopped after " round-limit " tool rounds")
            :rounds round
            :history history}
           (let [call-items (map #(select-keys % [:type :name :arguments :call_id]) calls)
                 outputs (map (fn [c]
                                {:type "function_call_output"
                                 :call_id (:call_id c)
                                 :output (execute c)})
                              calls)]
             (recur (into (into history call-items) outputs) (inc round)))))))))

(defn run-agent
  ([prompt] (run-agent prompt []))
  ([prompt history]
   (let [key (api-key)]
     (when-not key
       (throw (ex-info "No API key: set MINI_AGENT2_API_KEY or configure ~/.codex/config.toml" {})))
     (let [usage (atom {:input 0 :output 0 :total 0})
           respond (fn [input]
                     (let [body (call-model key (base-url)
                                            {:model (model)
                                             :instructions (instructions)
                                             :input input
                                             :tools (tools)
                                             :max_output_tokens 4096})
                           turn-usage (api-usage (:usage body))]
                       (when turn-usage
                         (swap! usage #(merge-with + % turn-usage)))
                       (:output body)))
           execute (fn [call]
                     (println (format "[tool] %s %s" (:name call) (:arguments call)))
                     (let [out (execute-tool call)]
                       (println "[result]" out)
                       out))]
       (assoc (react-loop respond execute
                          (conj (vec history)
                                {:type "message" :role "user" :content prompt}))
              :usage @usage)))))

;; ---- Ralph loop ----

(def complete-pattern #"(?m)^\s*<promise>COMPLETE</promise>\s*$")

(defn complete? [answer]
  (boolean (re-find complete-pattern (or answer ""))))

(defn ralph-prompt [task feedback]
  (str "You are running as a Ralph coding-loop iteration. Work directly in the current project.\n\n"
       "Task:\n" task "\n\n"
       "Inspect the existing code, make the smallest complete change, and run the project verification. "
       "Do not stop at a plan. When the task is genuinely complete and verification passes, "
       "end with exactly <promise>COMPLETE</promise>.\n"
       (when feedback
         (str "\nPrevious iteration feedback:\n" feedback "\n"))))

(defn ralph-loop
  "Run fresh-context iterations until the answer has the completion marker and
  `verify` succeeds. `iterate` maps a prompt to a react-loop result."
  [{:keys [task iterate verify max-iterations]
    :or {max-iterations 10}}]
  (loop [iteration 1
         feedback nil
         history []]
    (let [result (iterate (ralph-prompt task feedback))
          verification (when verify (verify (result :answer)))
          verified? (or (nil? verification) (zero? (:exit-code verification)))
          passed? (and (complete? (:answer result)) verified?)]
      (if passed?
        {:status :complete
         :iterations iteration
         :answer (:answer result)
         :verification verification
         :history history}
        (if (>= iteration max-iterations)
          {:status :max-iterations
           :iterations iteration
           :answer (:answer result)
           :verification verification
           :history history}
          (recur (inc iteration)
                 (str "Model answer:\n" (or (:answer result) "")
                      (when verification
                        (str "\n\nVerification exit code: " (:exit-code verification)
                             "\nVerification output:\n" (:output verification))))
                 (conj history {:iteration iteration
                                :answer (:answer result)
                                :verification verification})))))))

(defn ralph-log-file []
  (io/file (System/getProperty "user.dir") ".mini-agent2" "ralph" "progress.md"))

(defn log-ralph-iteration! [iteration result verification]
  (let [f (ralph-log-file)]
    (io/make-parents f)
    (spit f
          (str "\n## Iteration " iteration "\n\n"
               "Model answer:\n\n" (or (:answer result) "") "\n\n"
               (when verification
                 (str "Verification exit code: " (:exit-code verification) "\n\n"
                      "```\n" (:output verification) "\n```\n")))
          :append true)))

(defn run-ralph
  ([task] (run-ralph task {}))
  ([task {:keys [max-iterations verify-command]
          :or {max-iterations 10
               verify-command "lein test"}}]
   (let [iteration (atom 0)
         verify (when (seq (str/trim (or verify-command "")))
                  (fn [_] (run-bash verify-command)))
         result (ralph-loop
                 {:task task
                  :max-iterations max-iterations
                  :verify verify
                  :iterate (fn [prompt]
                             (let [n (swap! iteration inc)]
                               (println (format "\n=== Ralph iteration %d/%d ==="
                                                n max-iterations)))
                             (run-agent prompt))})]
     (doseq [{:keys [iteration answer verification]}
             (conj (:history result)
                   {:iteration (:iterations result)
                    :answer (:answer result)
                    :verification (:verification result)})]
       (log-ralph-iteration! iteration {:answer answer} verification))
     result)))

;; ---- CLI ----

(defn print-result [{:keys [answer rounds usage]}]
  (println "\n=== ANSWER ===")
  (println answer)
  (println (str "\n(tool rounds: " rounds ")"))
  (when usage
    (let [win (context-window)
          pct (when win (* 100.0 (/ (:input usage) win)))]
      (println (format "(usage: %d input, %d output, %d total%s)"
                       (:input usage) (:output usage) (:total usage)
                       (if pct (format ", %.1f%% context used" pct) ""))))))

(defn parse-ralph-args [args]
  (loop [xs (seq args)
         task []
         opts {:max-iterations 10 :verify-command "lein test"}]
    (if-let [x (first xs)]
      (case x
        "--max-iterations"
        (let [value (second xs)
              n (try (Long/parseLong value) (catch Exception _ nil))]
          (when-not (and n (pos? n))
            (throw (ex-info "--max-iterations requires a positive integer" {})))
          (recur (nnext xs) task (assoc opts :max-iterations n)))
        "--verify"
        (let [value (second xs)]
          (when (str/blank? value)
            (throw (ex-info "--verify requires a command" {})))
          (recur (nnext xs) task (assoc opts :verify-command value)))
        (recur (next xs) (conj task x) opts))
      (assoc opts :task (str/join " " task)))))

(defn -main [& args]
  (cond
    (= "ralph" (first args))
    (let [{:keys [task] :as opts} (parse-ralph-args (rest args))]
      (when (str/blank? task)
        (throw (ex-info "Usage: lein run -- ralph [--max-iterations N] [--verify CMD] TASK" {})))
      (let [{:keys [status iterations answer]} (run-ralph task opts)]
        (println (format "\n=== RALPH %s after %d iteration(s) ==="
                         (name status) iterations))
        (println answer)
        (when-not (= :complete status)
          (System/exit 1))))

    (seq args)
    (print-result (run-agent (str/join " " args)))

    :else
    (loop []
      (print "task> ") (flush)
      (when-let [line (read-line)]
        (when-not (contains? #{"quit" "exit" "q"} (str/trim line))
          (let [result (run-agent line @mem)]
            (print-result result)
            (reset! mem (:history result))
            (recur)))))))
