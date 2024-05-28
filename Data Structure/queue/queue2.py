# Implement Queue using Stacks ([https://leetcode.com/problems/implement-queue-using-stacks/] 
# The counterpart to the previous problem, 
# this reinforces how you can manipulate data structures to achieve different functionalities.
# Approach: Similar to "Implement Stack using Queues", 
# but now implement FIFO behavior of a queue using two stacks. 
# Focus on how you transfer elements between stacks during push and pop operations.



# If input_stack = [1, 2, 3] and output_stack = [], _
# transfer_if_necessary() moves all elements, 
# resulting in input_stack = [] and output_stack = [3, 2, 1].
# Now, pop() removes and returns 3



class MyQueue:
    def __init__(self):
        self.input_stack = []   # initial stack
        self.output_stack = []   # output stack in reverse order

    def push(self, x: int) -> None:
        # Pushes an element onto the input_stack.
        self.input_stack.append(x)

    def pop(self) -> int:
        self._transfer_if_necessary()
        return self.output_stack.pop()
    # removes and returns the (bottom of the input stack x1) 

    def peek(self) -> int:
        self._transfer_if_necessary()
        return self.output_stack[-1] 
    # peek the last element[-1](top stack) on the output_stack, 
    # which is first element on the input stack which simulate the behaviour of first in first out, 
    # you always deal with the first element in the Q, 
    # in this case, that element is top of the output stack
     

    def empty(self) -> bool:
        return not self.input_stack and not self.output_stack

    def _transfer_if_necessary(self) -> None:
        if not self.output_stack: # if output stack is empty
            while self.input_stack:
                self.output_stack.append(self.input_stack.pop())
                # pop the top element in the input stack(which is like last element/top stack element x5 in the Q)
                # store it in the bottom of the output stack, so it will be the last to handle
                # from first to handle x5(top stack) change to last to handle x5.
