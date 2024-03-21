name_file = open("name.txt", "r")
# if you set this to w, then readable will return false


# check if the file is readable
# print(name_file.readable())



# all the info in the file
# print(name_file.read())
# you only want to read the info




# you can also w, which is write
# you can also a, a stands for append
# r+ is read and write

# read indivial line, this will read the first line and move the cursor to next line
# print(name_file.readline()) # start with first line
# print(name_file.readline()) # that is the second line


# take everything and put it in a list
# print(name_file.readlines())

# access to the first element
# print(name_file.readlines()[0]) 


for names in name_file.readlines():
    print(names)


name_file.close()
# always close the file after you are done with it