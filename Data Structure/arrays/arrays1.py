# Find Numbers with Even Number of Digits (Easy)**
# Problem: Given an array `nums` of integers, return how many of them contain an even number of digits.
# Example: If the input is `nums = [12, 345, 2, 6, 7896]`, the output should be `2` 
# (because 12 and 7896 have an even number of digits).

def findNumbers(nums):
    count = 0
    for num in nums:
        # Convert number to string to count digits
        if len(str(num)) % 2 == 0:
            count += 1
    return count

# Example usage
nums = [12, 345, 2, 6, 7896]
result = findNumbers(nums)
print(result)  # Output should be 2
