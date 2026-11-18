
(println "EXCEPTION HANDLING AND ERROR MANAGEMENT")

(println "\n1. ARITHMETIC EXCEPTION")

(try
  (println "Result:" (/ 10 0))
  (catch ArithmeticException e
    (println "Caught ArithmeticException:" (.getMessage e))
    (println "Recovery: Cannot divide by zero."))
  (finally
    (println "Finally block executed for arithmetic exception.")))


(println "\n2. NULL POINTER EXCEPTION")

(try
  (throw (NullPointerException. "Name cannot be null"))
  (catch NullPointerException e
    (println "Caught NullPointerException:" (.getMessage e))
    (println "Recovery: Using a default name."))
  (finally
    (println "Finally block executed for null pointer exception.")))


(println "\n3. CUSTOM EXCEPTION USING EX-INFO")

(defn check-age [age]
  (if (neg? age)
    (throw (ex-info "Age cannot be negative"
                    {:age age :error :invalid-age}))
    (println "Valid age:" age)))

(try
  (check-age -5)
  (catch clojure.lang.ExceptionInfo e
    (println "Caught Custom Exception:" (ex-message e))
    (println "Exception Data:" (ex-data e))
    (println "Recovery: Please enter a valid age."))
  (finally
    (println "Finally block executed for custom exception.")))


(println "\n4. SUCCESSFUL EXECUTION")

(println "Program continues running after handling exceptions.")
