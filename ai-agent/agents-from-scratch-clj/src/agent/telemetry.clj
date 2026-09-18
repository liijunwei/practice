(ns agent.telemetry
  (:require [clojure.data.json :as json]
            [clojure.java.io :as io]))

(defn short-id []
  (subs (str (java.util.UUID/randomUUID)) 0 8))

(defn now-iso []
  (str (java.time.Instant/now)))

(defn- span->map [span]
  (into {} (remove (fn [[_ v]] (nil? v)) span)))

(defn new-telemetry
  ([] (new-telemetry "agent_telemetry.jsonl"))
  ([log-file]
   (atom {:log-file log-file
          :trace-id nil
          :metrics {:llm-calls 0
                    :llm-failures 0
                    :llm-retries 0
                    :tool-calls 0
                    :tool-failures 0
                    :memory-operations 0
                    :total-tokens 0
                    :total-latency-ms 0.0}
          :spans []})))

(defn start-trace! [telemetry]
  (let [id (short-id)]
    (swap! telemetry assoc :trace-id id)
    id))

(defn log-span! [telemetry span]
  (swap! telemetry update :spans conj span)
  (when-let [path (:log-file @telemetry)]
    (spit path (str (json/write-str (span->map span)) "\n") :append true)))

(defn- update-metric! [telemetry k f & args]
  (swap! telemetry update-in [:metrics k] #(apply f % args)))

(defn log-llm-call!
  [telemetry prompt-length response-length duration-ms & {:keys [success attempt error]
                                                          :or {success true attempt 1}}]
  (let [span {"span_id" (short-id)
              "trace_id" (or (:trace-id @telemetry) "no-trace")
              "event_type" "llm_call"
              "timestamp" (now-iso)
              "duration_ms" (when duration-ms (/ (Math/round (double (* duration-ms 100))) 100.0))
              "data" {"prompt_length" prompt-length
                      "response_length" response-length
                      "attempt" attempt
                      "success" success}
              "error" error}]
    (log-span! telemetry span)
    (update-metric! telemetry :llm-calls inc)
    (update-metric! telemetry :total-latency-ms + duration-ms)
    (when-not success (update-metric! telemetry :llm-failures inc))
    (when (> attempt 1) (update-metric! telemetry :llm-retries inc))))

(defn log-tool-call!
  [telemetry tool-name arguments & {:keys [result duration-ms error]}]
  (let [span {"span_id" (short-id)
              "trace_id" (or (:trace-id @telemetry) "no-trace")
              "event_type" "tool_call"
              "timestamp" (now-iso)
              "duration_ms" (when duration-ms (/ (Math/round (double (* duration-ms 100))) 100.0))
              "data" {"tool" tool-name
                      "arguments" arguments
                      "result" (when result (subs (str result) 0 (min 200 (count (str result)))))}
              "error" error}]
    (log-span! telemetry span)
    (update-metric! telemetry :tool-calls inc)
    (when error (update-metric! telemetry :tool-failures inc))))

(defn log-memory-operation!
  ([telemetry operation] (log-memory-operation! telemetry operation nil))
  ([telemetry operation data]
   (let [span {"span_id" (short-id)
               "trace_id" (or (:trace-id @telemetry) "no-trace")
               "event_type" "memory"
               "timestamp" (now-iso)
               "data" {"operation" operation
                       "data" (when data (subs data 0 (min 100 (count data))))}}]
     (log-span! telemetry span)
     (update-metric! telemetry :memory-operations inc))))

(defn log-decision! [telemetry choices selected & {:keys [duration-ms]}]
  (let [span {"span_id" (short-id)
              "trace_id" (or (:trace-id @telemetry) "no-trace")
              "event_type" "decision"
              "timestamp" (now-iso)
              "duration_ms" (when duration-ms (/ (Math/round (double (* duration-ms 100))) 100.0))
              "data" {"choices" choices
                      "selected" selected}}]
    (log-span! telemetry span)))

(defn- metrics-map [metrics]
  (let [llm-calls (:llm-calls metrics)
        llm-failures (:llm-failures metrics)
        tool-calls (:tool-calls metrics)
        tool-failures (:tool-failures metrics)
        llm-success-rate (if (pos? llm-calls) (- 1 (/ llm-failures llm-calls)) 0.0)
        tool-success-rate (if (pos? tool-calls) (- 1 (/ tool-failures tool-calls)) 0.0)
        avg-latency (if (pos? llm-calls) (/ (:total-latency-ms metrics) llm-calls) 0.0)]
    {"llm_calls" llm-calls
     "llm_failures" llm-failures
     "llm_retries" (:llm-retries metrics)
     "llm_success_rate" (format "%.2f%%" (* 100.0 llm-success-rate))
     "avg_latency_ms" (/ (Math/round (double (* 10 avg-latency))) 10.0)
     "tool_calls" tool-calls
     "tool_failures" tool-failures
     "tool_success_rate" (format "%.2f%%" (* 100.0 tool-success-rate))
     "memory_operations" (:memory-operations metrics)}))

(defn get-metrics [telemetry]
  (metrics-map (:metrics @telemetry)))

(defn get-recent-spans
  ([telemetry] (get-recent-spans telemetry 10))
  ([telemetry n]
   (mapv span->map (take-last n (:spans @telemetry)))))

(defn get-trace-spans [telemetry trace-id]
  (mapv span->map (filter #(= trace-id (get % "trace_id")) (:spans @telemetry))))

(defn clear! [telemetry]
  (let [log-file (:log-file @telemetry)]
    (swap! telemetry assoc :spans [] :trace-id nil :metrics
           {:llm-calls 0 :llm-failures 0 :llm-retries 0
            :tool-calls 0 :tool-failures 0 :memory-operations 0
            :total-tokens 0 :total-latency-ms 0.0})
    (when (and log-file (.exists (io/file log-file)))
      (.delete (io/file log-file)))))

(defn print-summary [telemetry]
  (let [m (get-metrics telemetry)]
    (println "\n" (apply str (repeat 40 "=")))
    (println "TELEMETRY SUMMARY")
    (println (apply str (repeat 40 "=")))
    (println "LLM Calls:     " (get m "llm_calls"))
    (println "  Success Rate:" (get m "llm_success_rate"))
    (println "  Avg Latency: " (get m "avg_latency_ms") "ms")
    (println "  Retries:     " (get m "llm_retries"))
    (println "Tool Calls:    " (get m "tool_calls"))
    (println "  Success Rate:" (get m "tool_success_rate"))
    (println "Memory Ops:    " (get m "memory_operations"))
    (println (apply str (repeat 40 "=")))))

(defmacro traced [telemetry event-type & body]
  `(let [start# (System/currentTimeMillis)
         result# (try {:value (do ~@body)}
                      (catch Throwable e# {:error e#}))]
     (log-span! ~telemetry
                {"span_id" (short-id)
                 "trace_id" (or (:trace-id @~telemetry) "no-trace")
                 "event_type" ~event-type
                 "timestamp" (now-iso)
                 "duration_ms" (- (System/currentTimeMillis) start#)
                 "data" {"function" (str '~event-type)}})
     (if-let [e# (:error result#)]
       (throw e#)
       (:value result#))))
