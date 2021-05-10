(ns armstrong-numbers)

(defn- digits
  [n]
  (->> n
       (iterate #(quot % 10))
       (take-while pos?)
       (map #(mod % 10))))


(defn- pow
  [base exp]
  (apply * (repeat exp base)))

(defn armstrong? [n]
  (let [bases (digits n)
        exp (count bases)]
    (->> (map #(pow % exp) bases)
         (apply +)
         (= n))))

