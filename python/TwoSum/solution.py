#!/usr/bin/env python3
from typing import List

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        diffMap = {}
        length = len(nums)
        for i in range(length):
            diff = target - nums[i]
            val = diffMap.get(nums[i])
            if val != None:
                return [val, i]
            diffMap[diff] = i
        return nums

sol = Solution()
sol.twoSum([1,2,3], 3)
