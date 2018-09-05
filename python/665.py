'''
665. Non-decreasing Array
Given an array with n integers, your task is to check if it could become
non-decreasing by modifying at most 1 element.
We define an array is non-decreasing if array[i] <= array[i + 1] holds
for every i (1 <= i < n).

Example 1:
Input: [4,2,3]
Output: True
Explanation: You could modify the first 4 to 1 to get a non-decreasing array.
'''

class Solution:
    def checkPossibility(nums):
        # 统计修改的次数
        i, cnt = 1, 0
        while i < len(nums) and cnt <= 1:
            # 发现一个递减的点
            if nums[i-1] > nums[i]:
                # i < 2 或 (i>2 且是小于关系）时修改前面的数字对后面没影响
                if i < 2 or nums[i-2] <= nums[i]:
                    
