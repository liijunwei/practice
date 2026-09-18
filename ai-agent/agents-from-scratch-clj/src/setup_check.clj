(ns setup-check
  (:require [clojure.java.io :as io]
            [clojure.string :as string]))

(defn- check-clojure []
  (println "✅ Clojure is running (Java" (System/getProperty "java.version") ")")
  true)

(defn- check-deepseek-api []
  (let [key (System/getenv "DEEPSEEK_API_KEY")]
    (if (string/blank? key)
      (do
        (println "❌ DEEPSEEK_API_KEY not set")
        (println "   export DEEPSEEK_API_KEY=...")
        false)
      (do
        (println "✅ DEEPSEEK_API_KEY is set")
        true))))

(defn- check-structure []
  (let [dirs ["src/shared"
              "src/agent"
              "test"]]
    (doseq [dir dirs]
      (if (.isDirectory (io/file dir))
        (println "✅" dir "/ directory exists")
        (println "❌" dir "directory not found")))
    (every? #(.isDirectory (io/file %)) dirs)))

(defn -main [& _]
  (println (apply str (repeat 50 "=")))
  (println "AI Agents from Scratch (Clojure) - Setup Verification")
  (println (apply str (repeat 50 "=")))
  (println)
  (let [checks [["Clojure" check-clojure]
                ["DeepSeek API" check-deepseek-api]
                ["Repository Structure" check-structure]]
        results (mapv (fn [[name check]]
                        (println "\n" name ":")
                        (check))
                      checks)]
    (println "\n" (apply str (repeat 50 "=")))
    (if (every? identity results)
      (println "✅ All checks passed! Run: lein run")
      (println "⚠️  Some checks failed. Fix the issues above."))
    (println (apply str (repeat 50 "=")))
    (shutdown-agents)))
