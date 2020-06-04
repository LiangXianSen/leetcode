# 14. 最长公共前缀

https://leetcode-cn.com/problems/longest-common-prefix/

# Description

编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 `""`。

**示例 1:**

```
输入: ["flower","flow","flight"]
输出: "fl"
```

**示例 2:**

```
输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
```

**说明:**

所有输入只包含小写字母 `a-z` 。



# Conclusion

1. 直接从给定的字符串数组中把第一个赋值给prefix
2. 然后继续遍历后面的字符串，拿着prefix 挨个字符跟它比对，
3. 相同则移动指针到下一个字符，不相同直接从当前位置裁剪prefix，然后比对下一个，依次如此
4. 如果指针移动到比prefix短，比字符串总长度长，裁剪prefix到当前指针
5. 如果都没有相同的prefix 会被裁剪成`0:0` length也是0 ，跳出循环，return结果

