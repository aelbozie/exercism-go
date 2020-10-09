def distance(strand_a: str, strand_b: str) -> int:
    if len(strand_a) != len(strand_b):
        raise ValueError("does not match")
    elif strand_a == strand_b:
        distance = 0
    else:
        both_strands = zip(strand_a, strand_b)
        changes = list(filter(lambda t: t[0] != t[1], both_strands))
        distance = len(changes)
    return distance
