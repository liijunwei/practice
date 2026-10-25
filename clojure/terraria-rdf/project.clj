(defproject terraria-rdf "0.1.0-SNAPSHOT"
  :description "Terraria 晶塔小镇 NPC 分组优化器（RDF + SHACL + CP-SAT）"
  :url "https://example.com/FIXME"
  :license {:name "EPL-2.0 OR GPL-2.0-or-later WITH Classpath-exception-2.0"
            :url "https://www.eclipse.org/legal/epl-2.0/"}
  :dependencies [[org.clojure/clojure "1.12.2"]
                 [org.clojure/tools.reader "1.5.0"]
                 [org.apache.jena/jena-arq "6.2.0"]
                 [io.github.mlimotte/quandary "0.9.1"]
                 [com.dcj/clj-shacl "0.2.0"]
                 [org.slf4j/slf4j-nop "2.0.18"]]
  ;; Jena 6 需要 Java 21+，本机默认 java 是 17，这里固定用 Java 23 运行项目
  :java-cmd "/Users/me/Library/Java/JavaVirtualMachines/temurin-23.0.2/Contents/Home/bin/java"
  :main ^:skip-aot terraria-rdf.core
  :target-path "target/%s"
  :repositories [["central" {:url "https://maven.aliyun.com/repository/central"}]]
  :profiles {:uberjar {:aot :all
                       :jvm-opts ["-Dclojure.compiler.direct-linking=true"]}})
