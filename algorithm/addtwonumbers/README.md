# 2. Add Two Numbers

https://leetcode.com/problems/add-two-numbers/

# Description

You are given two **non-empty** linked lists representing two non-negative integers. The digits are stored in **reverse order** and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

**Example:**

```
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
```

# Conclude

1. 通过修改指针实现的嵌套（golang）
2. 先求和，然后%10求商和余数，最后计算进位
3. 最后的进位记得算在进位