def is_isogram(string: str) -> bool:
    alphabet = [ch for ch in string.lower() if "a" <= ch <= "z"]
    alphabet_set = set(alphabet)
    return len(alphabet) == len(alphabet_set)
