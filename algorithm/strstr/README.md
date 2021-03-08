# 28. 实现 strStr()

https://leetcode-cn.com/problems/implement-strstr/

# Description

实现 `strStr()` 函数。

给定一个 `haystack` 字符串和一个 `needle` 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  **-1**。

**示例 1:**

```
输入: haystack = "hello", needle = "ll"
输出: 2
```

**示例 2:**

```
输入: haystack = "aaaaa", needle = "bba"
输出: -1
```

**说明:**

当 `needle` 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 `strstr()` 以及 Java的 `indexOf()` 定义相符。



# Conclusion

1. 设置两个变量，分别记录`haystack` 和 `needle`  每次移动的索引；
2. forloop 分别移动两个字符串下标比较；
3. 如果相同，继续移动下标；
4. 如果不相同，重置`needle` 和 `haystack` 的下标；
5. 如果到达`needle`的最后一个下标则结束比较。


