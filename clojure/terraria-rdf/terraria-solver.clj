(ns terraria-solver
  (:require [quandary.api :as qa]
            [shacl.core :as shacl])
  (:import (org.apache.jena.rdf.model ModelFactory)
           (org.apache.jena.query QueryFactory DatasetFactory QueryExecutionFactory)))

;; ========== 配置区 ==========
(def max-group-size 3)   ;; 每个小镇最多 3 人
(def max-groups 10)
(def solve-timeout 90.0)
(def score-scale 100000) ;; 把 log 乘数放大成整数，供 CP-SAT 使用

(def ont "https://terraria.example.org/ontology#")
(def rdfs "http://www.w3.org/2000/01/rdf-schema#")

;; ========== 加载 RDF ==========
(def model
  (doto (ModelFactory/createDefaultModel)
    (.read "terraria-ontology.ttl" "TURTLE")
    (.read "terraria-biomes.ttl" "TURTLE")
    (.read "npc-facts.ttl" "TURTLE")))

(def shacl-model
  (doto (ModelFactory/createDefaultModel)
    (.read "terraria-shacl.ttl" "TURTLE")))

(defn run-select [qstr]
  (with-open [qe (QueryExecutionFactory/create
                  (QueryFactory/create qstr)
                  (DatasetFactory/create model))]
    (vec (iterator-seq (.execSelect qe)))))

;; ========== SHACL 校验 ==========
(defn run-shacl-validate []
  (let [report (shacl/validate (.getGraph model) (.getGraph shacl-model))]
    (println "\n==== SHACL 校验报告 ====")
    (if (:conforms? report)
      (println "✅ 校验通过：无阻断性违规（偏好冲突、类型错误均未发现）")
      (doseq [v (filter #(= :violation (:severity %)) (:results report))]
        (println (str "❌ 违规: " (:message v) " | 节点: " (:focus v)))))
    report))

(def pred-labels
  {(str ont "lovesNeighbor") "喜爱"
   (str ont "likesNeighbor") "喜欢"
   (str ont "dislikesNeighbor") "不喜欢"
   (str ont "hatesNeighbor") "厌恶"})

(defn print-one-way-warnings []
  (let [qstr (str "PREFIX terraria: <" ont "> "
                  "PREFIX rdfs: <" rdfs "> "
                  "SELECT ?npc ?npcLabel ?pred ?other ?otherLabel WHERE { "
                  "?npc ?pred ?other . "
                  "?npc rdfs:label ?npcLabel . "
                  "?other rdfs:label ?otherLabel . "
                  "VALUES ?pred { terraria:lovesNeighbor terraria:likesNeighbor "
                  "terraria:dislikesNeighbor terraria:hatesNeighbor } "
                  "FILTER NOT EXISTS { ?other ?pred ?npc . } }")
        rows (run-select qstr)]
    (when (seq rows)
      (println "\n==== 单向偏好告警（Wiki 原生允许，仅提示） ====")
      (doseq [sol rows]
        (println (str "⚠️ "
                      (.getString (.getLiteral sol "npcLabel"))
                      " 对 "
                      (.getString (.getLiteral sol "otherLabel"))
                      " 存在单向「"
                      (get pred-labels (.getURI (.getResource sol "pred")) "偏好")
                      "」"))))))

;; ========== 数据读取 ==========
(defn get-npcs []
  (let [qstr (str "PREFIX terraria: <" ont "> "
                  "PREFIX rdfs: <" rdfs "> "
                  "SELECT ?npc ?label WHERE { ?npc a terraria:NPC ; rdfs:label ?label . }")]
    (into {} (for [sol (run-select qstr)]
               [(.getURI (.getResource sol "npc"))
                (.getString (.getLiteral sol "label"))]))))

(defn get-biomes []
  (let [qstr (str "PREFIX terraria: <" ont "> "
                  "PREFIX rdfs: <" rdfs "> "
                  "SELECT ?biome ?label WHERE { ?biome a terraria:Biome ; rdfs:label ?label . }")]
    (into {} (for [sol (run-select qstr)]
               [(.getURI (.getResource sol "biome"))
                (.getString (.getLiteral sol "label"))]))))

(def prop-multiplier
  {(str ont "lovesBiome") (Math/log 0.88)
   (str ont "likesBiome") (Math/log 0.94)
   (str ont "dislikesBiome") (Math/log 1.06)
   (str ont "hatesBiome") (Math/log 1.12)
   (str ont "lovesNeighbor") (Math/log 0.88)
   (str ont "likesNeighbor") (Math/log 0.94)
   (str ont "dislikesNeighbor") (Math/log 1.06)
   (str ont "hatesNeighbor") (Math/log 1.12)
   :solitude (Math/log 0.95)
   :crowd (Math/log 1.03)
   :princess (Math/log 0.88)})

(defn scale-score [d]
  (long (Math/round (* d score-scale))))

(defn load-biome-prefs []
  (let [qstr (str "PREFIX terraria: <" ont "> "
                  "SELECT ?npc ?pred ?biome WHERE { "
                  "?npc ?pred ?biome . "
                  "VALUES ?pred { terraria:lovesBiome terraria:likesBiome "
                  "terraria:dislikesBiome terraria:hatesBiome } }")]
    (reduce (fn [m sol]
              (assoc-in m [(.getURI (.getResource sol "npc"))
                           (.getURI (.getResource sol "biome"))]
                        (scale-score (prop-multiplier (.getURI (.getResource sol "pred"))))))
            {} (run-select qstr))))

(defn load-neighbor-prefs []
  (let [qstr (str "PREFIX terraria: <" ont "> "
                  "SELECT ?npc ?pred ?other WHERE { "
                  "?npc ?pred ?other . "
                  "VALUES ?pred { terraria:lovesNeighbor terraria:likesNeighbor "
                  "terraria:dislikesNeighbor terraria:hatesNeighbor } }")]
    (reduce (fn [m sol]
              (let [pred (.getURI (.getResource sol "pred"))]
                (assoc-in m [(.getURI (.getResource sol "npc"))
                             (.getURI (.getResource sol "other"))]
                          {:pred pred
                           :score (scale-score (prop-multiplier pred))})))
            {} (run-select qstr))))

;; ========== 主求解入口 ==========
(defn run-solver []
  (run-shacl-validate)
  (print-one-way-warnings)
  (let [npcs (get-npcs)
        biomes (get-biomes)
        npc-uris (vec (sort (keys npcs)))
        biome-uris (vec (sort (keys biomes)))
        npc-idx (zipmap npc-uris (range))
        biome-idx (zipmap biome-uris (range))
        n (count npc-uris)
        nb (count biome-uris)
        princess-uri (first (filter #(= "公主" (get npcs %)) npc-uris))
        princess-i (get npc-idx princess-uri)
        biome-prefs (load-biome-prefs)
        neighbor-prefs (load-neighbor-prefs)

        a-var (fn [i g] (str "a." i "." g))
        b-var (fn [g k] (str "b." g "." k))
        cnt-var (fn [g] (str "cnt." g))
        is1-var (fn [g] (str "is1." g))
        is3-var (fn [g] (str "is3." g))
        ab-var (fn [i g k] (str "ab." i "." g "." k))
        nn-var (fn [i j g] (str "nn." i "." j "." g))
        pp-var (fn [i g] (str "pp." i "." g))
        cost-var "cost"

        ;; 邻居偏好：[npc-i 目标-j pred-uri 分数]
        neigh-edges (vec (for [[u vmap] neighbor-prefs
                               [v pref] vmap
                               :let [i (get npc-idx u)
                                     j (get npc-idx v)]]
                           [i j (:pred pref) (:score pref)]))
        hate-edges (filter #(= (str ont "hatesNeighbor") (nth % 2)) neigh-edges)
        ;; 群系偏好：[npc-i 群系-k 分数]
        biome-edges (vec (for [[u bmap] biome-prefs
                               [b score] bmap
                               :let [i (get npc-idx u)
                                     k (get biome-idx b)]]
                           [i k score]))

        cost-neg
        (reduce + (filter neg?
                          (concat
                           (for [[_ _ score] biome-edges, _g (range max-groups)] score)
                           (for [[_ _ _ score] neigh-edges, _g (range max-groups)] score)
                           (for [i (range n), :when (not= i princess-i), _g (range max-groups)]
                             (scale-score (:princess prop-multiplier)))
                           (for [_g (range max-groups)]
                             (scale-score (:solitude prop-multiplier))))))
        cost-pos
        (reduce + (filter pos?
                          (concat
                           (for [[_ _ score] biome-edges, _g (range max-groups)] score)
                           (for [[_ _ _ score] neigh-edges, _g (range max-groups)] score)
                           (for [_g (range max-groups)]
                             (scale-score (:crowd prop-multiplier))))))

        domain
        (merge
         (into {} (for [i (range n), g (range max-groups)] [(a-var i g) [:boolean]]))
         (into {} (for [g (range max-groups), k (range nb)] [(b-var g k) [:boolean]]))
         (into {} (for [g (range max-groups)] [(cnt-var g) [:range 0 max-group-size]]))
         (into {} (for [g (range max-groups)] [(is1-var g) [:boolean]]))
         (into {} (for [g (range max-groups)] [(is3-var g) [:boolean]]))
         (into {} (for [[i k _] biome-edges, g (range max-groups)] [(ab-var i g k) [:boolean]]))
         (into {} (for [[i j _ _] neigh-edges, g (range max-groups)] [(nn-var i j g) [:boolean]]))
         (into {} (for [i (range n), :when (not= i princess-i), g (range max-groups)]
                    [(pp-var i g) [:boolean]]))
         {cost-var [:range cost-neg cost-pos]})

        equations
        (vec
         (concat
          ;; 每个 NPC 恰好分到一个小组
          (for [i (range n)]
            (list '= (mapv (fn [g] [(a-var i g)]) (range max-groups)) [[1]]))
          ;; 单组人数上限
          (for [g (range max-groups)]
            (list '<= (mapv (fn [i] [(a-var i g)]) (range n)) [[max-group-size]]))
          ;; cnt.g = 组内人数
          (for [g (range max-groups)]
            (list '= [[(cnt-var g)]] (mapv (fn [i] [(a-var i g)]) (range n))))
          ;; 每个小镇绑定一个群系
          (for [g (range max-groups)]
            (list '= (mapv (fn [k] [(b-var g k)]) (range nb)) [[1]]))
          ;; 群系偏好指示器 ab = a * b
          (for [[i k _] biome-edges, g (range max-groups)]
            (list '= [[(ab-var i g k)]] [[(a-var i g) (b-var g k)]]))
          ;; 邻居偏好指示器 nn = a * a
          (for [[i j _ _] neigh-edges, g (range max-groups)]
            (list '= [[(nn-var i j g)]] [[(a-var i g) (a-var j g)]]))
          ;; 公主 buff 指示器 pp = a * a_princess
          (for [i (range n), :when (not= i princess-i), g (range max-groups)]
            (list '= [[(pp-var i g)]] [[(a-var i g) (a-var princess-i g)]]))
          ;; is1.g <=> cnt.g == 1（独处折扣）
          (mapcat (fn [g] [(list '= [[(cnt-var g)]] [[1]] :only-if (is1-var g))
                           (list '!= [[(cnt-var g)]] [[1]] :only-if (str "!" (is1-var g)))])
                  (range max-groups))
          ;; is3.g <=> cnt.g == 3（拥挤惩罚）
          (mapcat (fn [g] [(list '= [[(cnt-var g)]] [[3]] :only-if (is3-var g))
                           (list '!= [[(cnt-var g)]] [[3]] :only-if (str "!" (is3-var g)))])
                  (range max-groups))
          ;; hate 互斥：互相厌恶的 NPC 不能同组
          (for [[i j _ _] hate-edges, g (range max-groups)]
            (list '<= [[(a-var i g)] [(a-var j g)]] [[1]]))
          ;; 对称破缺：组规模非递增（cnt.0 >= cnt.1 >= ...），大幅缩小等价搜索空间
          (for [g (range (dec max-groups))]
            (list '<= [[(cnt-var (inc g))]] [[(cnt-var g)]]))
          ;; 目标函数：最小化 log 得分总和（等价于最小化总价格乘数）
          [(list '= [[cost-var]]
                  (vec (concat
                        (for [[i k score] biome-edges, g (range max-groups)]
                          [score (ab-var i g k)])
                        (for [[i j _ score] neigh-edges, g (range max-groups)]
                          [score (nn-var i j g)])
                        (for [i (range n), :when (not= i princess-i), g (range max-groups)]
                          [(scale-score (:princess prop-multiplier)) (pp-var i g)])
                        (for [g (range max-groups)]
                          [(scale-score (:solitude prop-multiplier)) (is1-var g)])
                        (for [g (range max-groups)]
                          [(scale-score (:crowd prop-multiplier)) (is3-var g)]))))]
          ))

        problem {:domain domain :equations equations}
        result (qa/solve problem {:objective [:minimize cost-var]
                                  :timeout solve-timeout})
        status (-> result meta :quandary.quandary/status)
        status-name (when status (.name status))]

    (println "\n==== 求解结果 ====")
    (cond
      (= status-name "INFEASIBLE")
      (println "❌ 无解：请检查 hate 约束")

      (and (empty? result) (= status-name "UNKNOWN"))
      (println "⏱️ 超时且未找到可行解")

      (empty? result)
      (println "⚠️ 未返回任何解")

      :else
      (let [best (first (sort-by #(get % cost-var) result))
            total-log (double (/ (get best cost-var) score-scale))]
        (println (str "✅ 求解完成（状态：" status-name "）"))
        (println (str "总 log 得分：" (format "%.4f" total-log)
                      "（约等于总价格乘数 " (format "%.3f" (Math/exp total-log)) "）"))
        (doseq [g (range max-groups)]
          (let [member-is (filter #(= 1 (get best (a-var % g) 0)) (range n))]
            (when (seq member-is)
              (let [k (first (filter #(= 1 (get best (b-var g %) 0)) (range nb)))]
                (println (str "\n===== 小镇 #" g " | 群系：" (biomes (nth biome-uris k)) " ====="))
                (doseq [i member-is]
                  (println (str " · " (npcs (nth npc-uris i)))))))))))))

(defn -main [& _args]
  (run-solver))

;; 直接运行
(run-solver)
