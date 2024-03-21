# you want to add another name to the name.txt
# names = open("name.txt","a")

# names.write("\nOG - HR")
# # this new line char is necessary escpae char

# # however if you run the program twice,
# # you mess up the file
# # it will have this thing twice and on the same line
# # like OG - HROG - HR

# names.write("\nOGA - HRA")


# you can also overwrite a file
# THIS WILL OVERWRITE THE ENITIRE FILE
names = open("name.txt","w")
names.write("\nOGA - HRA")
names.close()


# you can also use w to create a new file
names2 = open("names2.txt","w")
# you can also create index.html if you want
# and add HTML code like
# names2.write("<p>THIS is HTML</p>")
names2.write("\nOGB - HRB")
names2.close()