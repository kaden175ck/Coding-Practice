import queue
from collections import deque

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

    # Dequeue elements (FIFO)
    while not my_queue.empty():
        item = my_queue.get()
        print("Processing:", item)

# *******************
#  2. Using collections.deque
# *******************

def queue_deque_example():
    my_queue = deque()

    # Enqueue (append)
    my_queue.append("Email 1")
    my_queue.append("Email 2")
    my_queue.append("Email 3")

    print("Initial Queue:", my_queue)

    # Dequeue (popleft)
    while my_queue:  # Check if queue is not empty
        email = my_queue.popleft()
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
