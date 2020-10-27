def is_pangram(sentence: str) -> bool:
    chars = set(sentence.lower())
    alphabet = set("abcdefghijklmnopqrstuvwxyz")
    return alphabet.issubset(chars)
