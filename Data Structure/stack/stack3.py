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
    # Initializes a list called result with the same length as the input. 
    # Each element in result is set to 0. [0,0,0,0,0]


    stack = []  # Each element in the stack is a (temperature, index) pair.

    for i, temp in enumerate(temperatures): 
        while stack and temp > stack[-1][0]:  
    # while stack is not empty AND
    # the current temperature (temp) > the temperature of the last element in the stack (stack[-1][0]).
    # if the current temp is higher than the temp at the Top of the Stack, then we found the warmer day 
    
            


            _, prev_index = stack.pop() # pop the(temperature, index) pair from the top of the stack.
            # _ is choosing to ignore the first element(the temp) in the tuple, which is temp,
            # pop the second part in the tuple, and assigned to prev_index

            result[prev_index] = i - prev_index
            # Calculate the difference in indices between the current day (i) 
            # and the day of the popped temperature. 
            # This difference is the number of days to wait for a warmer temperature.

            # Store this difference in the result list at the popped temperature's index.


        stack.append((temp, i))
        # Regardless of whether we popped anything, 
        # push the current day's (temperature, index) pair onto the stack.
        

    return result
