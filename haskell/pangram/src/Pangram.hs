module Pangram (isPangram) where
import qualified Data.Char as C (toLower)
import qualified Data.Set  as S (fromList, isSubsetOf)

isPangram :: String -> Bool
isPangram text =  alphabet `S.isSubsetOf` set
  where set =  S.fromList [C.toLower t | t <- text]
        alphabet = S.fromList ['a' .. 'z']
