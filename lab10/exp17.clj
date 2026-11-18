
(println "CLOJURE CORE IMMUTABLE DATA STRUCTURES")

(println "\n1. LIST")

(def my-list '(10 20 30 40))
(println "Original List:" my-list)

(println "First element:" (first my-list))
(println "Rest of List:" (rest my-list))

(def updated-list (conj my-list 5))
(println "After conj:" updated-list)
(println "Original List after operations:" my-list)


(println "\n2. VECTOR")

(def my-vector [10 20 30])
(println "Original Vector:" my-vector)

(def updated-vector (conj my-vector 40))
(println "After conj:" updated-vector)
(println "Element at index 1:" (get my-vector 1))
(println "Count:" (count my-vector))
(println "Original Vector after operations:" my-vector)


(println "\n3. MAP")

(def my-map {:name "Rahul" :age 20 :course "CSE"})
(println "Original Map:" my-map)

(def updated-map (assoc my-map :city "Delhi"))
(println "After assoc:" updated-map)

(def removed-map (dissoc my-map :age))
(println "After dissoc:" removed-map)

(println "Get name:" (get my-map :name))
(println "Original Map after operations:" my-map)


(println "\n4. SET")

(def my-set #{10 20 30})
(println "Original Set:" my-set)

(def updated-set (conj my-set 40))
(println "After conj:" updated-set)

(def removed-set (disj my-set 20))
(println "After disj:" removed-set)

(println "Contains 10:" (contains? my-set 10))
(println "Original Set after operations:" my-set)