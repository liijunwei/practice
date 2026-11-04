(defproject mini-agent1 "0.1.0-SNAPSHOT"
  :description "A minimal ReAct agent on DeepSeek's Responses API"
  :license {:name "EPL-2.0 OR GPL-2.0-or-later WITH Classpath-exception-2.0"
            :url "https://www.eclipse.org/legal/epl-2.0/"}
  :dependencies [[org.clojure/clojure "1.12.2"]
                 [org.clojure/data.json "2.5.1"]]
  :main mini-agent1.core
  :repl-options {:init-ns mini-agent1.core})
