# * **Implement Stack using Queues ([https://leetcode.com/problems/implement-stack-using-queues/](https://leetcode.com/problems/implement-stack-using-queues/))**  This is a classic warm-up to help you understand the interplay between queue (FIFO) behavior and stack (LIFO) behavior.  




# This article dives into how to create a stack using just queues. 
# A stack, you might recall, follows the LIFO (Last In, First Out) principle, 
# similar to a stack of plates. The one you add last is the first one you can remove.



# * Design a data structure that simulates the functionality of a stack using only queue operations (enqueue, dequeue, etc.).

# **What is a Stack?**

# * **LIFO:** A stack operates on the principle of LIFO (Last-In, First-Out).  This is like putting dishes on top of each other – the last dish you add is the first one you take off.
# * **Operations:** Key stack operations include:
#     * **push(x):** Add element 'x' to the top.
#     * **pop():** Remove and return the top element.
#     * **top():** Return the top element without removing it.
#     * **empty():** Check if the stack is empty.

# **The Challenge**

# * **Mismatch:** Queues follow FIFO (First-In, First-Out). The challenge is to reverse this behavior using only queue operations to achieve LIFO. 

# **What the Code Needs to Do**

# Your code must provide the following class, implementing the functionalities as described:

# ```python
# class MyStack:
#     def __init__(self):
#         # Initialize your data structure here.

#     def push(self, x: int) -> None:
#         # Push element x onto stack.

#     def pop(self) -> int:
#         # Removes the element on top of the stack and returns that element.

#     def top(self) -> int:
#         # Get the top element.

#     def empty(self) -> bool:
#         # Returns whether the stack is empty.
# ```