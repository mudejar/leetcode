#!/usr/bin/env python3

class Solution:
    def isValid(self, s: str) -> bool:
        letters = []
        bracket_pairs = { "}":"{", "]":"[", ")":"(" }
        closing_brackets = set('}])')
        for letter in s:
            if letter in closing_brackets:
                if len(letters) == 0: return False
                pop = letters.pop()
                if pop != bracket_pairs[letter]:
                    return False
            else:
                letters.append(letter)
        if len(letters) == 0:
            return True
        return False

sol = Solution()
# print(sol.isValid("()[]{}"))
print(sol.isValid("()"))
# print(sol.isValid("(]"))
