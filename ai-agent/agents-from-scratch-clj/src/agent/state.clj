(ns agent.state)

(defn new-state []
  (atom {:steps 0
         :done false
         :current-plan nil
         :last-action nil}))

(defn increment-step! [state]
  (swap! state update :steps inc))

(defn mark-done! [state]
  (swap! state assoc :done true))

(defn reset-state! [state]
  (reset! state {:steps 0
                 :done false
                 :current-plan nil
                 :last-action nil}))

(defn state-map [state]
  @state)
