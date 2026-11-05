(defproject mini-agent2 "0.1.0-SNAPSHOT"
  :description "A minimal ReAct agent with a Ralph loop"
  :license {:name "MIT License"
            :url "https://opensource.org/licenses/MIT"}
  :dependencies [[org.clojure/clojure "1.12.2"]
                 [org.clojure/data.json "2.5.1"]]
  :main mini-agent2.core
  :repl-options {:init-ns mini-agent2.core})
