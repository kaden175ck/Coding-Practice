class Student:
    def __init__(self, name, major, gpa):
        self.name = name
        self.major = major
        self.gpa = gpa

# class function can be used as obejct as the class
    def on_honor_roll(self):
        if self.gpa >= 3.5:
            return True
        else:
            return False
