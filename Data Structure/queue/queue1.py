# * **Implement Stack using Queues 
#([https://leetcode.com/problems/implement-stack-using-queues/]
# Design a data structure that simulates the functionality of a stack 
# using only queue operations (enqueue, dequeue, etc.).

# **What is a Stack?**
# **LIFO:** A stack operates on the principle of LIFO (Last-In, First-Out).  
# **Operations:** Key stack operations include:
# **push(x):** Add element 'x' to the top.
# **pop():** Remove and return the top element.
# **top():** Return the top element without removing it.
# **empty():** Check if the stack is empty. 

from collections import deque
# we use deque becuase we want more flexibility and able to deque from both ends

class MyStack:
    def __init__(self):
        self.queue1 = deque()  # Main queue
        self.queue2 = deque()  # Temporary queue

    # `queue1` 用作主队列，用于存放栈中的元素，而 `queue2` 是临时队列，用于辅助操作。

    # 注意push是push到stack top level
    def push(self, x: int) -> None:
        self.queue2.append(x)
        # append(x): 这是队列的一个方法，用于将元素 x 添加到队列Q2的末尾。

        
        # Move all elements from queue1 to queue2
        while self.queue1:
            self.queue2.append(self.queue1.popleft())
            # popleft(): 这是 deque 的一个方法，它的作用是从队列Q1的左侧（头部）弹出一个元素。这个操作会：
            # 返回队列左侧的第一个元素。
            # 将这个元素从队列Q1中移除并添加到Q2

        # Swap queue names  (queue2 now becomes the main queue)
        self.queue1, self.queue2 = self.queue2, self.queue1

    def pop(self) -> int:
        return self.queue1.popleft()  # Pop from the front of queue1

    def top(self) -> int:
        return self.queue1[0]  # Top is at the front of queue1

    def empty(self) -> bool:
        return not self.queue1


# `push` 方法

# 这样做的目的是确保新添加的元素总是出现在 `queue1` 的前端，从而模拟栈的后进先出（LIFO）特性。
# - 然后，该方法会将 `queue1` 中的所有元素依次移动到 `queue2` 中，保持新元素在 `queue1` 的前端。
# - 最后，`queue1` 和 `queue2` 交换角色，
# 即原来的 `queue2`（现在包含所有元素，新元素在前）成为新的 `queue1`，而原来的 `queue1` 变成空的 `queue2`，
# 准备下一次操作。

# ### `pop` 方法
# - 从 `queue1` 的前端移除一个元素并返回。由于新元素总是被添加到前端，所以这也模拟了栈的后进先出特性。

# ### `top` 方法
# - 返回 `queue1` 前端的元素，但不移除它。这对应于查看栈顶元素。

# ### `empty` 方法
# - 检查 `queue1` 是否为空，从而判断栈是否为空。

# 总的来说，这个类通过两个队列实现了一个栈的所有基本操作，确保了栈的后进先出特性。
# 使用 `deque` 的好处是它允许从两端快速添加和删除元素，这在此实现中尤为重要。