---
title: LeetCode Examples
date: 2021-11-30 19:36:20
tags: 
  - Python
  - Leetcode
categories: Python
---

> LeetCode 上的一些习题，挺难做的鸭，目前只是着手了简单题目

<!--MORE-->

## Example1 - 两数之和

给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。

```python
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        l = len(nums)
        for i in range(l):
            for j in range(i+1,l):
                if nums[i] + nums[j] == target:
                    return i,j
```





## Example7 - 整数反转

给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号）。

```python
class Solution:
    def reverse(self, x: int) -> int:
        if x > 0:
            xlist = [int(i) for i in str(x)]
            l = len(xlist)
            r = 0
            for i in range(l):
                r += xlist[l-i-1] * 10 ** (l-i-1)
            return r if -2 ** 31 < r < 2 ** 31 - 1 else 0
        elif x < 0:
            x = x // -1
            xlist = [int(i) for i in str(x)]
            l = len(xlist)
            r = 0
            for i in range(l):
                r += xlist[l-i-1] * 10 ** (l-i-1)
            r *= -1
            return r if -2 ** 31 < r < 2 ** 31 - 1 else 0
        elif x == 0:
            return x
```

## Example9 - 回文数

给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。

```python
class Solution:
    def isPalindrome(self, x: int) -> bool:
        str1 = str(x)
        l = len(str1)
        for i in range(l):
            if str1[i] != str1[l-i-1]:
                return False
        return True
```

## Example217 - 存在重复元素

给定一个整数数组，判断是否存在重复元素。

如果存在一值在数组中出现至少两次，函数返回 `true` 。如果数组中每个元素都不相同，则返回 `false` 。

```python
class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        l = len(nums)
        s = set()
        for i in nums:
            s.add(i)
        if len(s) == len(nums):
            return False
        else:
            return True
#        l = len(nums)
#        for i in range(l):
#            for j in range(i+1,l):
#                if nums[i] == nums[j]:
#                    return True
#        return False
```

