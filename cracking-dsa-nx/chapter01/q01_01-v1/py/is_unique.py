def is_unique_for_with_set(s):
    """Return True if all characters in string s are unique (Set implementation)."""
    validate_is_unique_input(s)
    seen = set()
    for c in s:
        if c in seen:
            return False
        seen.add(c)
    return True

def is_unique_leader_follower(s):
    """Return True if all characters in string s are unique (Nested Loops implementation -- no additional data structures, optimized for memory footprint)."""
    validate_is_unique_input(s)
    for i in range(len(s)):
        for j in range(i + 1, len(s)):
            if s[i] == s[j]:
                return False
    return True

def is_unique_bit_vector(s):
    """Return True if all characters in string s are unique (Bit vector implementation. Large memory footprint, but highly efficient)."""
    validate_is_unique_input(s)
    seen = 0
    for c in s:
        check_bit = 1 << ord(c)
        if check_bit & seen:
            return False
        seen |= check_bit
    return True

def validate_is_unique_input(s):
    if s is None:
        raise TypeError("Expected a string, got None")
    if not isinstance(s, str):
        raise TypeError("Expected a string, got {type(s).__name__}")
        
