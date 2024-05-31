# Number of Islands (Medium)
# [https://leetcode.com/problems/number-of-islands/]
# Employ a queue for Breadth-First Search (BFS).  
# Start from a land cell ('1'), enqueue it, and within a loop, 
# dequeue and explore neighboring land cells while marking them as visited.  
# Repeat for every unvisited land cell to find all distinct islands.
# 这个问题要求计算一个由 1（陆地）和 0（水）组成的二维网格中岛屿的数量。
# 可以通过使用队列实现广度优先搜索来遍历每个岛屿。

