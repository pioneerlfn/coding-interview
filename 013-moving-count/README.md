## 题目描述

地上有一个m行和n列的方格。一个机器人从坐标0,0的格子开始移动，每一次只能向左，右，上，下四个方向移动一格，但是不能进入行坐标和列坐标的数位之和大于k的格子。

例如，当k为18时，机器人能够进入方格（35,37），因为3+5+3+7 = 18。但是，它不能进入方格（35,38），因为3+5+3+8 = 19。请问该机器人能够达到多少个格子？

---
### 分析


m行n列的方格可以看成是m*n的矩阵。同样在这个矩阵中，除了边界上的格子之外其他格子都有四个相邻的格子。

机器人从坐标（0，0）开始移动。但它准备进入坐标为（i，j）的个姿势，通过检查坐标的数位和来判断机器人是否能够进入（之前是判断字符是否相同）。如果能够进入（i，j），则判断它的四个相邻的格子。递归实现。


