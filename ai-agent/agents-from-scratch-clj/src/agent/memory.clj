(ns agent.memory
  (:require [clojure.string :as string]))

(defn new-memory []
  (atom []))

(defn add! [memory item]
  (swap! memory (fn [items]
                  (if (and (seq item) (not-any? #(= item %) items))
                    (conj items item)
                    items))))

(defn get-all [memory]
  @memory)

(defn get-recent
  ([memory] (get-recent memory 5))
  ([memory n] (vec (take-last n @memory))))

(defn search [memory query]
  (let [query (string/lower-case query)]
    (vec (filter #(string/includes? (string/lower-case %) query) @memory))))

(defn clear! [memory]
  (reset! memory []))

(defn size [memory]
  (count @memory))
