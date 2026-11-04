(defproject mini-agent1 "0.1.0-SNAPSHOT"
  :description "A minimal ReAct agent on DeepSeek's Responses API"
  :license {:name "MIT License"
            :url "https://opensource.org/licenses/MIT"}
  :dependencies [[org.clojure/clojure "1.12.2"]
                 [org.clojure/data.json "2.5.1"]]
  :main mini-agent1.core
  :repl-options {:init-ns mini-agent1.core})
