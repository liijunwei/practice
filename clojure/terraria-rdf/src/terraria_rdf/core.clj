(ns terraria-rdf.core
  (:gen-class)
  (:require [terraria-rdf.solver :as solver]))

(defn -main
  "Terraria NPC 晶塔分组求解器入口。"
  [& args]
  (solver/run-solver))
