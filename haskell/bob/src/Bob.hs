module Bob (responseFor) where

import qualified Data.Text as T
import qualified Data.Char as C
import           Data.Text (Text)


responseFor :: Text -> Text
responseFor text
    | asking
    , yelling
    , hasLetters = T.pack "Calm down, I know what I'm doing!"
    | yelling
    , hasLetters = T.pack "Who, chill out!"
    | asking = T.pack "Sure."
    | empty = T.pack "Fine. Be that way!"
    | otherwise  = T.pack "Whatever."
        where
            strippedText = T.strip text
            hasLetters = any C.isLetter (T.unpack strippedText)
            asking = T.isSuffixOf (T.pack "?") strippedText
            empty = all C.isSpace (T.unpack strippedText)
            yelling = text == T.toUpper strippedText