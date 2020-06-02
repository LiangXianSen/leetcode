# 1. Two Sum

https://leetcode-cn.com/problems/two-sum/

# Description

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:
```
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```

# Conclusion

无脑遍历，简单粗暴：

$$
时间复杂性： O(n^2)
$$

Hashpmap：

数量级小的情况下看不出优势，而且golang中创建map开销大

$$
时间复杂度：O(n)
$$
