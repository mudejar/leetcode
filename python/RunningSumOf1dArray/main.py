#!/usr/bin/env python3

from typing import List

class Solution:
    def runningSum(self, nums: List[int]) -> List[int]:
        sum = 0
        arr = []
        for i in range(len(nums)):
            sum += nums[i]
            arr.append(sum)
        return arr

sol = Solution()
print(sol.runningSum([1,2,3,4]))
