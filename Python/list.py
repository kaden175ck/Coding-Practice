friends = ["Kv","Kr","Jm"]
print(friends)

print(friends[2])
print(friends[0])

print(friends[-1]) 
print(friends[-2]) 
# rememer the last position is always negative 1

print(friends[1:])


friends = ["Kv","Kr","Jm","Kdan","AB","BC"]
print(friends[1:4])
# not including 4, so it grab Kdan and end

friends[1] = "Mike"
print(friends[1])
print("    /| ")
print("   / | ")
print("  /  | ")
print(" /   | ")
print("/____| ")


lucky_numbers = [4, 8, 15, 16, 23, 42]
friends = ["Kevin", "Karen", "Jim", "Oscar", "Toby"]


friends.extend(lucky_numbers)
print(friends)




friends = ["Kevin", "Karen", "Jim", "Oscar", "Toby"]
friends.append("Creed")
print(friends)
friends.insert(1,"Kelly")
print(friends)


print(friends)
friends.remove("Jim")
print(friends)
friends.clear()
print(friends)



friends = ["Kevin", "Karen", "Jim", "Oscar", "Toby"]
print(friends)
friends.pop()#pop the last element in the list
print(friends)


print(friends.index("Kevin"))
#print(friends.index("Toby")) just pop, will give error



# count the number of appeareace
print(friends.count("Kevin"))



# sort this list
friends.sort()
print(friends)
lucky_numbers.sort()
print(lucky_numbers)
lucky_numbers.reverse()
print(lucky_numbers)

print("    /| ")
print("   / | ")
print("  /  | ")
print(" /   | ")
print("/____| ")


# copy list
friends2 = friends.copy()
print(friends2)