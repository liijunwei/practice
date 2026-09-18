(defproject agents-from-scratch-clj "0.1.0-SNAPSHOT"
  :description "AI agents from scratch, translated to Clojure and backed by DeepSeek."
  :dependencies [[org.clojure/clojure "1.12.0"]
                 [org.clojure/data.json "2.5.1"]]
  :main ^:skip-aot complete-example
  :source-paths ["src"]
  :test-paths ["test"]
  :aliases {"setup" ["run" "-m" "setup-check"]
            "example" ["run" "-m" "complete-example"]})
