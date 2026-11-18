
(println "SEQUENCE TRANSFORMATIONS")

(println "\n1. VECTOR")

(def numbers [1 2 3 4 5 6 7 8 9 10])

(println "Original Vector:" numbers)

(println "Squares using map:" (map #(* % %) numbers))

(println "Even numbers using filter:" (filter even? numbers))

(println "Sum using reduce:" (reduce + numbers))

(println "First 5 elements using take:" (take 5 numbers))

(println "After dropping first 5 elements:" (drop 5 numbers))

(println "Threading macro result:"
         (->> numbers
              (map #(* % %))
              (filter even?)
              (take 3)))


(println "\n2. LIST")

(def values '(2 4 6 8 10 12))

(println "Original List:" values)

(println "Increment using map:" (map inc values))

(println "Numbers greater than 6 using filter:"
         (filter #(> % 6) values))

(println "Product using reduce:" (reduce * values))

(println "First 3 elements using take:" (take 3 values))

(println "After dropping first 3 elements:" (drop 3 values))

(println "Threading macro result:"
         (->> values
              (map inc)
              (filter #(> % 7))))
