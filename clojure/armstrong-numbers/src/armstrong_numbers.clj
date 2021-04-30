(ns armstrong-numbers)

(defn- digits
  [n]
  (->> n
       (iterate #(quot % 10))
       (take-while pos?)
       (mapv #(mod % 10))
       (rseq)))

(defn- pow
  [base exp]
  (reduce * (repeat exp base)))

(defn armstrong? [n]
  (let [bases (digits n)
        exp (count bases)]
    (= n
       (reduce + (map #(pow % exp) bases)))))

