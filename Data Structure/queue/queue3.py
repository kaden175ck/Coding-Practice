# Rotting Oranges (https://leetcode.com/problems/rotting-oranges/)
# A straightforward application of a queue for Breadth-First Search (BFS). 
# Think of it as the infection spreading step-by-step.
# Approach: Use a queue for multi-source BFS. Treat rotten oranges as initial sources, 
# enqueue them, and iteratively "infect" their fresh neighbors in each time step.

s