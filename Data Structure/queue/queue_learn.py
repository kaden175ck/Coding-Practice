# Understanding Queues

# FIFO: Queues operate on the principle of FIFO (First-In, First-Out). Imagine a line for a roller coaster—the first person in line gets to ride first. Similarly, in a queue, the first element added is the first one removed.

# Operations: Here are the fundamental operations on a queue:

# Enqueue: Adds an element to the rear (back) of the queue.
# Dequeue: Removes and returns the element at the front (head) of the queue.
# Peek/Front: Returns the element at the front without removing it.
# Empty: Checks if the queue is empty.
# Size: Returns the number of elements in the queue.

import queue
from collections import deque
# Deques (double-ended queues) from the collections module are versatile. 
# You can use them as queues, with efficient adds and removals from both ends.

# *******************
#  1. Using queue.Queue
# *******************

def queue_queue_example():
    my_queue = queue.Queue()

    # Enqueue elements
    my_queue.put("Task 1")
    my_queue.put(2)  # Queues can handle different data types
    my_queue.put("Process A")

    print("Initial Queue:", my_queue.queue)  # Note: queue.Queue doesn't expose the underlying list
    # Output: initial Queue: deque(['Task 1', 2, 'Process A'])

    # Dequeue elements (FIFO)
    while not my_queue.empty():
        item = my_queue.get()
        print("Processing:", item)

# *******************
#  2. Using collections.deque
# *******************

def queue_deque_example():
    my_queue = deque()
    # collections.deque: Demonstrates using a deque as a queue. 
    # While efficient, keep in mind deques allow operations at both ends, not strictly FIFO.

    # Enqueue (append)
    my_queue.append("Email 1") # Enqueue
    my_queue.append("Email 2")
    my_queue.append("Email 3")

    print("Initial Queue:", my_queue)

    # Dequeue (popleft)
    while my_queue:  # Check if queue is not empty, while true
        email = my_queue.popleft() # Dequeue (from the left end, which is the first one, email 1)
        print("Sending:", email)

# *******************
# Example Scenario: Print Job Queue
# *******************

def print_queue_example():
    print_jobs = queue.Queue()

    print_jobs.put(("Document1.pdf", 10))  # (filename, num_pages)
    print_jobs.put(("Report.xlsx", 2))
    print_jobs.put(("Image.jpg", 1))

    print("*** Printing ***")
    while not print_jobs.empty():
        job, pages = print_jobs.get()
        print("Printing", job, "-", pages, "pages")


# Main Execution:
if __name__ == "__main__":
    queue_queue_example()
    print("\n")
    queue_deque_example()
    print("\n")
    print_queue_example()
