(ns shared.llm
  (:require [clojure.data.json :as json]
            [clojure.java.shell :as shell]
            [clojure.string :as string]))

(def default-stops ["</s>" "\n\n" "User:" "Assistant:"])
(def default-model "deepseek-v4-pro")
(def endpoint "https://api.deepseek.com/chat/completions")

(defn- env [name]
  (System/getenv name))

(defn new-llm
  [model-path & {:keys [temperature max-tokens n-ctx model api-key]
                 :or {temperature 0.2 max-tokens 4096 n-ctx 2048}}]
  (let [path-model (when (and model-path
                              (not (string/blank? model-path))
                              (not (string/includes? model-path ".gguf"))
                              (not (string/includes? model-path "/")))
                     model-path)]
    {:api-key (or api-key (env "DEEPSEEK_API_KEY"))
     :model (or model path-model (env "DEEPSEEK_MODEL") default-model)
     :temperature temperature
     :max-tokens max-tokens
     :n-ctx n-ctx}))

(defn generate
  ([llm prompt] (generate llm prompt nil nil))
  ([llm prompt temperature stop]
   (let [api-key (:api-key llm)
         temperature (or temperature (:temperature llm))
         stop (or stop default-stops)]
     (when (string/blank? api-key)
       (throw (ex-info "DEEPSEEK_API_KEY is not set" {:hint "export DEEPSEEK_API_KEY=..."})))
     (let [body (json/write-str
                 {"model" (:model llm)
                  "messages" [{"role" "user" "content" prompt}]
                  "temperature" temperature
                  "max_tokens" (:max-tokens llm)
                  "stream" false
                  "stop" stop})
           args ["curl" "-sSf"
                 endpoint
                 "-H" (str "Authorization: Bearer " api-key)
                 "-H" "Content-Type: application/json"
                 "-d" body]
       result (apply shell/sh args)]
       (if (zero? (:exit result))
         (some-> result
                 :out
                 json/read-str
                 (get-in ["choices" 0 "message" "content"])
                 string/trim)
         (throw (ex-info (str "DeepSeek request failed: " (:err result))
                         {:exit (:exit result)
                          :err (:err result)})))))))
