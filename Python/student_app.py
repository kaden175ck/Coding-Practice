# object is an actual student
# class define what a student data type is


# from a file import a student class
from class_objects import Student



# create an object of a class
# obejct is an instance of a class
student1 = Student("A", "B", 3.1, False)
student2 = Student("AA", "BB", 3.2, True)
print(student1.name)
print(student1.gpa)
print(student2.gpa)

