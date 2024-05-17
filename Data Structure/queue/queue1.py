# * **Implement Stack using Queues 
#([https://leetcode.com/problems/implement-stack-using-queues/]
#(https://leetcode.com/problems/implement-stack-using-queues/))
## This is a classic warm-up to help you understand the interplay between queue (FIFO) behavior 
## and stack (LIFO) behavior.  


# how to create a stack using just queues. 
# A stack, you might recall, follows the LIFO (Last In, First Out) principle, 
# similar to a stack of plates. The one you add last is the first one you can remove.



# * Design a data structure that simulates the functionality of a stack 
# using only queue operations (enqueue, dequeue, etc.).

# **What is a Stack?**

# * **LIFO:** A stack operates on the principle of LIFO (Last-In, First-Out).  This is like putting dishes on top of each other – the last dish you add is the first one you take off.
# * **Operations:** Key stack operations include:
#     * **push(x):** Add element 'x' to the top.
#     * **pop():** Remove and return the top element.
#     * **top():** Return the top element without removing it.
#     * **empty():** Check if the stack is empty.

# **The Challenge**

# * **Mismatch:** Queues follow FIFO (First-In, First-Out). 
# The challenge is to reverse this behavior using only queue operations to achieve LIFO. 

# Your code must provide the following class, implementing the functionalities as described:


class MyStack:
    def __init__(self):
        self.queue1 = []  # Main queue
        self.queue2 = []  # Temporary queue

    def push(self, x: int) -> None:
        # Push onto queue2(Temporary) to maintain recent elements at the front/ Top
        self.queue2.append(x)  

        # Move all elements from queue1 to queue2
        while self.queue1:
            self.queue2.append(self.queue1.pop(0))

        # Swap queue names  (queue2 now becomes the main queue)
        self.queue1, self.queue2 = self.queue2, self.queue1 

    def pop(self) -> int:
        return self.queue1.pop(0)  # Pop from the front of queue1

    def top(self) -> int:
        return self.queue1[0]  # Top is at the front of queue1

    def empty(self) -> bool:
        return not self.queue1 