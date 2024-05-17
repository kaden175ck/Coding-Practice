# **Problem Description**

# You're given an array `temperatures` 
# where `temperatures[i]` is the daily temperature on day `i`. 
# Your task is to return an array 
# where each element represents how many days you have to wait 
# until a warmer temperature occurs. 
# If there's no future day with a warmer temperature, put a 0 in that position.

# **Example:**

# * **Input:** temperatures = [73, 74, 75, 71, 69, 72, 76, 73] 
# * **Output:** [1, 1, 4, 2, 1, 1, 0, 0]

# **Stack-based Solution**

def daily_temperatures(temperatures):
    result = [0] * len(temperatures)
    stack = []  # Stack of (temperature, index) pairs

    for i, temp in enumerate(temperatures): # index, then temperature, (index, temperature)
        while stack and temp > stack[-1][0]:  
            # while stack is true, means stack is not empty
            # -1 get the last element, this case is the top element in the stack
            # because we are storing tuples in stack
            # 0 would give the temperature, 1 would give the indexs
            # so you are comparing the temp


            _, prev_index = stack.pop()
            # choosing to ignore the first element in the tuple, which is temp,
            # pop the index, the index part is assigned to prev_index
            # The _ is used to discard the temperature part of the tuple, as we only need the index.

            result[prev_index] = i - prev_index
        stack.append((temp, i))
        # this tuples are stored in the stack

    return result
