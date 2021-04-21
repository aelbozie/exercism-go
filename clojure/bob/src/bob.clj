(ns bob
  (:require [clojure.string :as string]))

(defn yelling?
  "Checks if shouting if all capitals and has letters."
  [s]
  (and (some #(Character/isLetter %) s)
       (= (string/upper-case s) s)))


(defn response-for [s]
  (let [s (string/trim s)]
    (cond
      (and (yelling? s)
           (string/ends-with? s "?")) "Calm down, I know what I'm doing!"
      (string/ends-with? s "?") "Sure."
      (string/blank? s) "Fine. Be that way!"
      (yelling? s) "Whoa, chill out!"
      :else "Whatever.")))

