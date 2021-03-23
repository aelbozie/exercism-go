module SpaceAge (Planet(..), ageOn) where

data Planet = Mercury
            | Venus
            | Earth
            | Mars
            | Jupiter
            | Saturn
            | Uranus
            | Neptune

ageOn :: Planet -> Float -> Float
ageOn planet seconds = seconds / orbitalPeriod planet

orbitalPeriod :: Planet -> Float
orbitalPeriod planet = 31557600 * scale where
    scale = case planet of
        Mercury -> 0.2408467
        Venus   -> 0.61519726
        Earth   -> 1.0
        Mars    -> 1.880815
        Jupiter -> 11.862615
        Saturn  -> 29.447498
        Uranus  -> 84.016846
        Neptune -> 164.79132
