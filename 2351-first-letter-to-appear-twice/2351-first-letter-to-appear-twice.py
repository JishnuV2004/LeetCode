class Solution:
    def repeatedCharacter(self, s: str) -> str:
        freq = {}
        for i in range(len(s)):
            if s[i] in freq:
                freq[s[i]] += 1
            else:
                freq[s[i]] = 1

            if freq[s[i]] == 2:
                return s[i]

        