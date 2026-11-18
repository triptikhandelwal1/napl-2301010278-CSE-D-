
(println "JAVA INTEROPERABILITY IN CLOJURE")

(println "\n1. JAVA STATIC METHODS")

(println "Absolute value:" (Math/abs -25))
(println "Square root:" (Math/sqrt 144))
(println "Current time:" (System/currentTimeMillis))


(println "\n2. CREATING JAVA OBJECTS")

(def java-list (new java.util.ArrayList))
(.add java-list "Apple")
(.add java-list "Banana")
(.add java-list "Mango")

(println "Java ArrayList:" java-list)

(def java-map (new java.util.HashMap))
(.put java-map "Name" "Rahul")
(.put java-map "Age" 20)
(.put java-map "Course" "CSE")

(println "Java HashMap:" java-map)


(println "\n3. INVOKING INSTANCE METHODS")

(println "List size:" (.size java-list))
(println "First element:" (.get java-list 0))
(println "Map value:" (.get java-map "Name"))
(println "Map contains Age:" (.containsKey java-map "Age"))


(println "\n4. CONVERTING COLLECTIONS")

(def clojure-vector [10 20 30 40])

(def converted-list
  (java.util.ArrayList. clojure-vector))

(println "Clojure Vector:" clojure-vector)
(println "Converted to Java ArrayList:" converted-list)

(def converted-vector
  (vec java-list))

(println "Java ArrayList:" java-list)
(println "Converted to Clojure Vector:" converted-vector)
